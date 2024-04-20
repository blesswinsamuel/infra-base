package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Config struct {
	KubeconfigPath         string
	FolderAnnotation       string
	Label                  string
	LabelValue             string
	TargetFolder           string
	Resources              []string
	RequestMethod          string
	RequestURL             string
	RequestPayload         string
	Namespaces             []string
	UniqueFilenames        bool
	Enable5xx              bool
	IgnoreAlreadyProcessed bool
	Method                 string
}

// https://github.com/kiwigrid/k8s-sidecar/blob/master/src/sidecar.py
func main() {
	config, err := parseConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse config")
	}

	k8sSidecar, err := newK8sSidecar(config)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create k8s sidecar")
	}

	ctx := context.Background()

	switch config.Method {
	case "LIST":
		for _, res := range config.Resources {
			for _, ns := range config.Namespaces {
				if err := k8sSidecar.listResources(ctx, res, ns); err != nil {
					log.Fatal().Err(err).Msg("Failed to list resources")
				}
			}
		}
	default:
		k8sSidecar.watchForChanges(ctx)
	}
}

func parseConfig() (*Config, error) {
	kubeconfigPath := os.Getenv("KUBECONFIG")
	folderAnnotation := os.Getenv("FOLDER_ANNOTATION")
	if folderAnnotation == "" {
		folderAnnotation = "k8s-sidecar-target-directory"
	}
	label := os.Getenv("LABEL")
	if label == "" {
		return nil, fmt.Errorf("LABEL is empty")
	}
	labelValue := os.Getenv("LABEL_VALUE")
	targetFolder := os.Getenv("FOLDER")
	if targetFolder == "" {
		return nil, fmt.Errorf("FOLDER is empty")
	}
	resourcesStr := os.Getenv("RESOURCES")
	if resourcesStr == "" {
		resourcesStr = "configmap"
	} else if resourcesStr == "both" {
		resourcesStr = "configmap,secret"
	}
	resources := strings.Split(resourcesStr, ",")

	requestMethod := os.Getenv("REQ_METHOD")
	requestURL := os.Getenv("REQ_URL")
	requestPayload := os.Getenv("REQ_PAYLOAD")

	namespace := os.Getenv("NAMESPACE")
	if namespace == "" {
		ns, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
		if err != nil && !os.IsNotExist(err) {
			return nil, fmt.Errorf("failed to get namespace from service account")
		}
		if ns == nil {
			return nil, fmt.Errorf("namespace is empty")
		}
		namespace = string(ns)
	}
	if namespace == "ALL" {
		namespace = metav1.NamespaceAll
	}
	namespaces := strings.Split(namespace, ",")

	uniqueFilenames := os.Getenv("UNIQUE_FILENAMES") == "true"
	enable5xx := os.Getenv("ENABLE_5XX") == "true"
	ignoreAlreadyProcessed := os.Getenv("IGNORE_ALREADY_PROCESSED") == "true"
	method := os.Getenv("METHOD")
	return &Config{
		KubeconfigPath:         kubeconfigPath,
		FolderAnnotation:       folderAnnotation,
		Label:                  label,
		LabelValue:             labelValue,
		TargetFolder:           targetFolder,
		Resources:              resources,
		RequestMethod:          requestMethod,
		RequestURL:             requestURL,
		RequestPayload:         requestPayload,
		Namespaces:             namespaces,
		UniqueFilenames:        uniqueFilenames,
		Enable5xx:              enable5xx,
		IgnoreAlreadyProcessed: ignoreAlreadyProcessed,
		Method:                 method,
	}, nil
}

func newKubeClient(kubeconfig string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

type k8sSidecar struct {
	cfg                    *Config
	kubeClient             *kubernetes.Clientset
	resourcesVersionMap    map[string]map[string]string
	resourcesObjectMap     map[string]map[string]runtime.Object
	resourcesDestFolderMap map[string]map[string]string
}

func newK8sSidecar(cfg *Config) (*k8sSidecar, error) {
	kubeClient, err := newKubeClient(cfg.KubeconfigPath)
	if err != nil {
		return nil, err
	}
	return &k8sSidecar{
		cfg:                    cfg,
		kubeClient:             kubeClient,
		resourcesVersionMap:    map[string]map[string]string{"secret": {}, "configmap": {}},
		resourcesObjectMap:     map[string]map[string]runtime.Object{"secret": {}, "configmap": {}},
		resourcesDestFolderMap: map[string]map[string]string{"secret": {}, "configmap": {}},
	}, nil
}

func (ks *k8sSidecar) listResources(ctx context.Context, resource string, namespace string) error {
	listOpts := metav1.ListOptions{}
	if ks.cfg.LabelValue != "" {
		listOpts.LabelSelector = fmt.Sprintf("%s=%s", ks.cfg.Label, ks.cfg.LabelValue)
	} else {
		listOpts.LabelSelector = ks.cfg.Label
	}
	switch resource {
	case "configmap":
		configmaps, err := ks.kubeClient.CoreV1().ConfigMaps(namespace).List(ctx, listOpts)
		if err != nil {
			return err
		}
		existKeys := make(map[string]bool)
		for _, item := range configmaps.Items {
			metadata := item.ObjectMeta
			existKeys[metadata.Namespace+"/"+metadata.Name] = true
			if ks.cfg.IgnoreAlreadyProcessed {
				if ks.resourcesVersionMap[resource][metadata.Namespace+"/"+metadata.Name] == metadata.ResourceVersion {
					continue
				}
				ks.resourcesVersionMap[resource][metadata.Namespace+"/"+metadata.Name] = metadata.ResourceVersion
			}
			log.Debug().Str("namespace", metadata.Namespace).Str("name", metadata.Name).Msg("Processing configmap")

			ks.processConfigMap(item, false)
		}
	case "secret":
		// secrets, err := ks.kubeClient.CoreV1().Secrets(namespace).List(ctx, listOpts)
		// if err != nil {
		// 	return err
		// }
		// TODO
	}
	return nil
}

func (ks *k8sSidecar) processConfigMap(cm corev1.ConfigMap, isRemoved bool) (bool, error) {
	filesChanged := false
	resource := "configmap"
	metadata := cm.ObjectMeta
	destFolder := ks.getDestinationFolder(metadata)
	oldConfigMap, _ := ks.resourcesObjectMap[resource][metadata.Namespace+"/"+metadata.Name].(*corev1.ConfigMap)
	oldDestFolder := ks.resourcesDestFolderMap[resource][metadata.Namespace+"/"+metadata.Name]
	if isRemoved {
		delete(ks.resourcesObjectMap[resource], metadata.Namespace+"/"+metadata.Name)
	} else {
		ks.resourcesObjectMap[resource][metadata.Namespace+"/"+metadata.Name] = cm.DeepCopy()
		ks.resourcesDestFolderMap[resource][metadata.Namespace+"/"+metadata.Name] = destFolder
	}

	if cm.Data == nil && cm.BinaryData == nil {
		log.Warn().Str("namespace", metadata.Namespace).Str("name", metadata.Name).Msg("No data/binaryData field in configmap")
		return false, nil
	}

	if cm.Data != nil {
		log.Debug().Msg("Found data on configmap")
		changed, err := ks.iterateData(cm.Data, destFolder, cm.ObjectMeta, resource, ContentTypeText, isRemoved)
		if err != nil {
			return false, err
		}
		filesChanged = filesChanged || changed
	}
	if oldConfigMap.Data != nil && !isRemoved {
		if oldDestFolder == destFolder {
			for _, key := range oldConfigMap.Data {
				if _, ok := cm.Data[key]; ok {
					delete(oldConfigMap.Data, key)
				}
			}
		}
		changed, err := ks.iterateData(oldConfigMap.Data, oldDestFolder, oldConfigMap.ObjectMeta, resource, ContentTypeText, true)
		if err != nil {
			return false, err
		}
		filesChanged = filesChanged || changed
	}
	return filesChanged, nil
}

func (ks *k8sSidecar) processSecret() {

}

func (ks *k8sSidecar) iterateData(data map[string]string, destFolder string, metadata metav1.ObjectMeta, resource string, contentType string, removeFiles bool) (bool, error) {
	filesChanged := false
	for dataKey, dataContent := range data {
		fc, err := ks.updateFile(dataKey, dataContent, destFolder, metadata, resource, contentType, removeFiles)
		if err != nil {
			return false, err
		}
		filesChanged = filesChanged || fc
	}
	return filesChanged, nil
}

func (ks *k8sSidecar) updateFile(dataKey string, dataContent string, destFolder string, metadata metav1.ObjectMeta, resource string, contentType string, removeFile bool) (bool, error) {
	fileName, fileData, err := ks.getFileDataAndName(dataKey, dataContent, contentType)
	if err != nil {
		return false, err
	}
	if ks.cfg.UniqueFilenames {
		fileName = uniqueFilename(fileName, metadata.Namespace, resource, metadata.Name)
	}
	if removeFile {
		return ks.removeFile(destFolder, fileName)
	}
	return ks.writeDataToFile(destFolder, fileName, fileData, contentType)
}

func uniqueFilename(filename string, namespace string, resource string, resourceName string) string {
	return "namespace_" + namespace + "." + resource + "_" + resourceName + "." + filename
}

func (ks *k8sSidecar) removeFile(destFolder string, fileName string) (bool, error) {
	fullFilename := filepath.Join(destFolder, fileName)
	if _, err := os.Stat(fullFilename); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	if err := os.Remove(fullFilename); err != nil {
		return false, err
	}
	return true, nil
}

func (ks *k8sSidecar) writeDataToFile(destFolder string, fileName string, fileData []byte, contentType string) (bool, error) {
	fullFilename := filepath.Join(destFolder, fileName)
	if _, err := os.Stat(destFolder); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(destFolder, 0755); err != nil {
				return false, err
			}
		} else {
			return false, err
		}
	}
	if _, err := os.Stat(fullFilename); err != nil {
		if os.IsNotExist(err) {
			if err := os.WriteFile(fullFilename, fileData, 0644); err != nil {
				return false, err
			}
			return true, nil
		}
		return false, err
	}
	if err := os.WriteFile(fullFilename, fileData, 0644); err != nil {
		return false, err
	}
	return true, nil
}

const (
	ContentTypeText         = "text"
	ContentTypeBase64Binary = "base64binary"
)

func (ks *k8sSidecar) getFileDataAndName(fullFilename string, content string, contentType string) (fileName string, fileData []byte, err error) {
	if contentType == ContentTypeBase64Binary {
		fileData, err = base64.StdEncoding.DecodeString(content)
		if err != nil {
			return "", nil, err
		}
	} else {
		fileData = []byte(content)
	}
	fileName = fullFilename
	if strings.HasSuffix(fullFilename, ".url") {
		fileName = fullFilename[:len(fullFilename)-4]
		if contentType == ContentTypeBase64Binary {
			fileURL := string(fileData)
			resp, err := http.Get(fileURL)
			if err != nil {
				return "", nil, err
			}
			defer resp.Body.Close()
			fileData, err = io.ReadAll(resp.Body)
			if err != nil {
				return "", nil, err
			}
		} else {
			resp, err := http.Get(string(fileData))
			if err != nil {
				return "", nil, err
			}
			defer resp.Body.Close()
			fileData, err = io.ReadAll(resp.Body)
			if err != nil {
				return "", nil, err
			}
		}
	}
	return fileName, fileData, nil
}

func (ks *k8sSidecar) getDestinationFolder(metadata metav1.ObjectMeta) string {
	if metadata.Annotations != nil {
		if val, ok := metadata.Annotations[ks.cfg.FolderAnnotation]; ok {
			var destFolder string
			if filepath.IsAbs(val) {
				destFolder = val
			} else {
				destFolder = filepath.Join(ks.cfg.TargetFolder, val)
			}
			log.Info().Msgf("Found a folder override annotation, placing the %s in: %s", metadata.Name, destFolder)
			return destFolder
		}
	}
	return ks.cfg.TargetFolder
}

func (ks *k8sSidecar) watchForChanges(ctx context.Context) {

}
