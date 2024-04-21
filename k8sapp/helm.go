package k8sapp

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	"gopkg.in/yaml.v3"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type ChartInfo struct {
	Repo    *string `json:"repo"`
	Chart   *string `json:"chart"`
	Version *string `json:"version"`
}

type HelmProps struct {
	ChartInfo           ChartInfo
	ChartFileNamePrefix string
	ReleaseName         string
	Namespace           string
	Values              map[string]interface{}
	PatchResource       func(resource *unstructured.Unstructured)
}

func NewHelm(scope kubegogen.Scope, props *HelmProps) {
	globals := GetGlobalContext(scope)
	chartsCacheDir := fmt.Sprintf("%s/%s", globals.CacheDir, "charts")
	if err := os.MkdirAll(chartsCacheDir, os.ModePerm); err != nil {
		log.Fatalln("MkdirAll failed", err)
	}
	if _, err := exec.LookPath("helm"); err != nil {
		log.Fatalln("helm LookPath failed", err)
	}
	if props.ChartInfo.Repo == nil {
		log.Fatalf("props.ChartInfo is nil for %s", props.ReleaseName)
	}
	chartFileName := *props.ChartInfo.Chart + "-" + *props.ChartInfo.Version + ".tgz"
	if props.ChartFileNamePrefix != "" {
		chartFileName = props.ChartFileNamePrefix + *props.ChartInfo.Version + ".tgz"
	}
	chartPath := fmt.Sprintf("%s/%s", chartsCacheDir, chartFileName)
	if _, err := os.Stat(chartPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Printf("Fetching chart '%s' from repo '%s' version '%s' ...", *props.ChartInfo.Chart, *props.ChartInfo.Repo, *props.ChartInfo.Version)
			cmd := exec.Command("helm", "pull", *props.ChartInfo.Chart, "--repo", *props.ChartInfo.Repo, "--destination", chartsCacheDir, "--version", *props.ChartInfo.Version)
			if out, err := cmd.CombinedOutput(); err != nil {
				log.Println("Error occured during helm pull command", string(out))
				log.Fatalln("Error occured during helm pull command", err)
			} else {
				if len(out) > 0 {
					log.Println(string(out))
				}
			}
		} else {
			log.Fatalln("helm Stat failed", err)
		}
	}
	namespace := props.Namespace
	if namespace == "" {
		namespace = scope.Namespace()
	}

	cmd := exec.Command(
		"helm",
		"template",
		props.ReleaseName,
		chartPath,
		"--namespace",
		namespace,
		"--kube-version=v1.27.4",
		"--include-crds",
		"--skip-tests",
		"--no-hooks",
		"--values",
		"-",
	)
	cmd.Stdin = strings.NewReader(infrahelpers.ToJSONString(props.Values))
	out, err := cmd.Output()
	if err != nil {
		msg := fmt.Sprintf("helm template failed: %s", err)
		if ee, ok := err.(*exec.ExitError); ok {
			msg = fmt.Sprintf("helm template failed: %s\n%s", err, string(ee.Stderr))
		}
		panic(msg)
	}

	dec := yaml.NewDecoder(bytes.NewReader(out))
	i := 0
	for {
		i++
		var obj map[string]any
		err := dec.Decode(&obj)
		if err != nil {
			if err == io.EOF {
				break
			}
			// fmt.Println("Error decoding yaml:\n%v", string(out))
			panic(err)
		}
		if len(obj) == 0 {
			continue
		}
		runtimeObj := &unstructured.Unstructured{Object: obj}
		if props.PatchResource != nil {
			props.PatchResource(runtimeObj)
		}
		scope.AddApiObject(runtimeObj)
		// scope.ApiObjectFromMap("api-"+strconv.Itoa(i), kubegogen.ApiObjectProps{
		// 	// TypeMeta: v1.TypeMeta{
		// 	// 	APIVersion: obj["apiVersion"].(string),
		// 	// 	Kind:       obj["kind"].(string),
		// 	// },
		// 	Unstructured: unstructured.Unstructured{Object: obj},
		// })
	}
}
