package packager

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
)

type App interface {
	Construct
	OutDir() string
	Synth()
}

type AppProps struct {
	Outdir       string
	CacheDir     string
	DeleteOutDir bool
	PatchNdots   bool
}

type app struct {
	*construct
	props AppProps
}

func NewApp(props AppProps) App {
	app := &app{construct: &construct{}, props: props}
	// app.construct.node = newNode("root", Construct(app))
	app.construct.node = (*node[Construct])(nil).AddChildNode("$$root", app)
	return app
}

func patchObject(apiObject ApiObject) {
	switch apiObject.GetKind() {
	case "Deployment":
		deploymentUnstructured := apiObject.GetObject().(*unstructured.Unstructured)
		var deployment appsv1.Deployment
		err := runtime.DefaultUnstructuredConverter.FromUnstructured(deploymentUnstructured.UnstructuredContent(), &deployment)
		if err != nil {
			log.Fatalf("FromUnstructured: %v", err)
		}
		deployment.Spec.Template.Spec.DNSConfig = &corev1.PodDNSConfig{
			Options: []corev1.PodDNSConfigOption{
				{
					Name:  "ndots",
					Value: infrahelpers.Ptr("1"),
				},
			},
		}
		unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&deployment)
		if err != nil {
			log.Fatalf("ToUnstructured: %v", err)
		}
		apiObject.SetObject(unstructured.Unstructured{Object: unstructuredObj})
	case "StatefulSet":
		statefulsetUnstructured := apiObject.GetObject().(*unstructured.Unstructured)
		var statefulset appsv1.StatefulSet
		err := runtime.DefaultUnstructuredConverter.FromUnstructured(statefulsetUnstructured.UnstructuredContent(), &statefulset)
		if err != nil {
			log.Fatalf("FromUnstructured: %v", err)
		}
		statefulset.Spec.Template.Spec.DNSConfig = &corev1.PodDNSConfig{
			Options: []corev1.PodDNSConfigOption{
				{
					Name:  "ndots",
					Value: infrahelpers.Ptr("1"),
				},
			},
		}
		unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&statefulset)
		if err != nil {
			log.Fatalf("ToUnstructured: %v", err)
		}
		apiObject.SetObject(unstructured.Unstructured{Object: unstructuredObj})
	case "CronJob":
		cronjobUnstructured := apiObject.GetObject().(*unstructured.Unstructured)
		var cronjob batchv1.CronJob
		err := runtime.DefaultUnstructuredConverter.FromUnstructured(cronjobUnstructured.UnstructuredContent(), &cronjob)
		if err != nil {
			log.Fatalf("FromUnstructured: %v", err)
		}
		cronjob.Spec.JobTemplate.Spec.Template.Spec.DNSConfig = &corev1.PodDNSConfig{
			Options: []corev1.PodDNSConfigOption{
				{
					Name:  "ndots",
					Value: infrahelpers.Ptr("1"),
				},
			},
		}
		unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&cronjob)
		if err != nil {
			log.Fatalf("ToUnstructured: %v", err)
		}
		apiObject.SetObject(unstructured.Unstructured{Object: unstructuredObj})
	}
}

func (a *app) Synth() {
	fileNo := 0
	files := map[string][]ApiObject{}
	var synth func(node *node[Construct], currentChartID []string, level int)
	synth = func(node *node[Construct], currentChartID []string, level int) {
		if node == nil {
			return
		}
		apiObjects := []ApiObject{}
		chartCount := 0
		for _, childNode := range node.children {
			// for i := 0; i < level; i++ {
			// 	fmt.Print("  ")
			// }
			// fmt.Println(node.id)
			thisChartID := currentChartID
			if _, ok := childNode.value.(Chart); ok {
				chartCount++
				thisChartID = append(thisChartID, fmt.Sprintf("%02d", chartCount), childNode.ID())
			}
			// fmt.Println(strings.Join(currentChartID, "-"), reflect.TypeOf(childNode.value), thisChartID)
			synth(childNode, thisChartID, level+1)
			if apiObject, ok := childNode.value.(ApiObject); ok {
				if a.props.PatchNdots {
					patchObject(apiObject)
				}
				apiObjects = append(apiObjects, apiObject)
			}
		}
		if len(apiObjects) > 0 {
			currentChartID := strings.Join(currentChartID, "-")
			if _, ok := files[currentChartID]; !ok {
				fileNo++
			}
			files[currentChartID] = append(files[currentChartID], apiObjects...)
		}
	}
	synth(a.node, []string{}, 0)
	fileContents := map[string][]byte{}
	for _, currentChartID := range infrahelpers.MapKeys(files) {
		apiObjects := files[currentChartID]
		filePath := path.Join(a.props.Outdir, fmt.Sprintf("%s.yaml", currentChartID))
		// fmt.Println(filePath, len(apiObjects))
		for i, apiObject := range apiObjects {
			// fmt.Printf("  - %s/%s/%s\n", apiObject.GetAPIVersion(), apiObject.GetNamespace(), apiObject.GetName())
			if i != 0 {
				fileContents[filePath] = append(fileContents[filePath], []byte("---\n")...)
			}
			fileContents[filePath] = append(fileContents[filePath], apiObject.ToYAML()...)
		}
	}
	if a.props.DeleteOutDir {
		if err := os.RemoveAll(a.props.Outdir); err != nil {
			log.Fatalf("RemoveAll: %v", err)
		}
	}
	if err := os.MkdirAll(a.props.Outdir, 0755); err != nil {
		log.Fatalf("MkdirAll: %v", err)
	}
	for filePath, fileContent := range fileContents {
		if err := os.WriteFile(filePath, fileContent, 0644); err != nil {
			log.Fatalf("WriteFile: %v", err)
		}
	}
}

func (a *app) OutDir() string {
	return a.props.Outdir
}
