package packager

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
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

func (a *app) Synth() {
	fileNo := 0
	files := map[string][]ApiObject{}
	fileNos := map[string]int{}
	var synth func(node *node[Construct], currentChartID string, level int)
	synth = func(node *node[Construct], currentChartID string, level int) {
		if node == nil {
			return
		}
		apiObjects := []ApiObject{}
		for _, node := range node.children {
			for i := 0; i < level; i++ {
				fmt.Print("  ")
			}
			fmt.Println(node.id)
			thisChartID := currentChartID
			if _, ok := node.value.(Chart); ok {
				// chartID := node.id
				// if currentChartID != "" {
				// 	thisChartID = fmt.Sprintf("%s-%s", currentChartID, chartID)
				// } else {
				// 	thisChartID = chartID
				// }
				thisChartID = node.FullID()
			}
			synth(node, thisChartID, level+1)
			// if chart, ok := node.value.(Chart); ok {
			// 	chart.Synth()
			// }
			if apiObject, ok := node.value.(ApiObject); ok {
				apiObjects = append(apiObjects, apiObject)
			}
		}
		if len(apiObjects) > 0 {
			if _, ok := files[currentChartID]; !ok {
				fileNos[currentChartID] = fileNo
				fileNo++
			}
			files[currentChartID] = append(files[currentChartID], apiObjects...)
		}
	}
	synth(a.node, "", 0)
	fileContents := map[string][]byte{}
	for _, currentChartID := range infrahelpers.MapKeys(files) {
		apiObjects := files[currentChartID]
		filePath := path.Join(a.props.Outdir, fmt.Sprintf("%04d-%s.yaml", fileNos[currentChartID], currentChartID))
		fmt.Println(filePath, len(apiObjects))
		for i, apiObject := range apiObjects {
			fmt.Printf("  - %s/%s/%s\n", apiObject.GetAPIVersion(), apiObject.GetNamespace(), apiObject.GetName())
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

type cdk8sapp struct {
	cdk8sConstruct
	app cdk8s.App

	props AppProps
}

func NewCdk8sApp(props AppProps) App {
	appProps := &cdk8s.AppProps{
		Outdir:         jsii.String(props.Outdir),
		YamlOutputType: cdk8s.YamlOutputType_FILE_PER_CHART,
		// YamlOutputType: cdk8s.YamlOutputType_FOLDER_PER_CHART_FILE_PER_RESOURCE,
		// RecordConstructMetadata: jsii.Bool(true),
	}
	app := cdk8s.NewApp(appProps)
	return &cdk8sapp{
		cdk8sConstruct: cdk8sConstruct{app},
		app:            app,
		props:          props,
	}
}

func (a *cdk8sapp) Synth() {
	if a.props.DeleteOutDir {
		if err := os.RemoveAll(a.props.Outdir); err != nil {
			log.Fatalf("RemoveAll: %v", err)
		}
	}

	a.app.Synth()
}

func (a *cdk8sapp) OutDir() string {
	return *a.app.Outdir()
}
