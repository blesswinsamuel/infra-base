package packager

import (
	"fmt"
	"log"
	"os"
	"path"

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
				thisChartID = node.FullID()
			}
			synth(node, thisChartID, level+1)
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
