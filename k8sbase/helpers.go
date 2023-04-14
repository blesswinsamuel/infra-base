package k8sbase

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/k8s"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"gopkg.in/yaml.v3"
)

var cacheDir = os.Getenv("CACHE_DIR")

func init() {
	if cacheDir == "" {
		cacheDir = "./cache"
	}
}

type NamespaceProps struct {
	Name string
}

func NewNamespace(scope constructs.Construct, namespaceName string) constructs.Construct {
	scope.Node().SetContext(jsii.String("namespace"), namespaceName)
	chart := cdk8s.NewChart(scope, jsii.String("namespace"), &cdk8s.ChartProps{
		DisableResourceNameHashes: jsii.Bool(true),
	})
	return k8s.NewKubeNamespace(chart, jsii.String("namespace"), &k8s.KubeNamespaceProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String(namespaceName),
			Labels: &map[string]*string{
				"name": jsii.String(namespaceName),
			},
		},
	})
}

func GetNamespace(scope constructs.Construct) *string {
	return jsii.String(scope.Node().TryGetContext(jsii.String("namespace")).(string))
}

func MergeAnnotations(annotations ...map[string]string) map[string]string {
	merged := make(map[string]string)
	for _, annotation := range annotations {
		for k, v := range annotation {
			merged[k] = v
		}
	}
	return merged
}

func MergeMaps[K comparable, V any](annotations ...map[K]V) map[K]V {
	merged := make(map[K]V)
	for _, annotation := range annotations {
		for k, v := range annotation {
			merged[k] = v
		}
	}
	return merged
}

func MergeLists[T any](annotations ...[]T) []T {
	merged := make([]T, 0)
	for _, annotation := range annotations {
		merged = append(merged, annotation...)
	}
	return merged
}

func Ternary[V any](cond bool, trueVal V, falseVal V) V {
	if cond {
		return trueVal
	}
	return falseVal
}

func Ptr[T any](v T) *T {
	return &v
}

func ToYamlString(v any) *string {
	buf := bytes.NewBuffer(nil)
	enc := yaml.NewEncoder(buf)
	enc.SetIndent(2)
	err := enc.Encode(v)
	if err != nil {
		panic(err)
	}
	return jsii.String(buf.String())
}

func FromYamlString[T any](s string) T {
	var v T
	err := yaml.Unmarshal([]byte(s), &v)
	if err != nil {
		panic(err)
	}
	return v
}

func GoTemplate(s string, data interface{}) *string {
	tmpl, err := template.New("tmpl").Parse(s)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		panic(err)
	}
	return jsii.String(buf.String())
}

func NewApp(outputDir string) cdk8s.App {
	app := cdk8s.NewApp(&cdk8s.AppProps{
		Outdir:         jsii.String(outputDir),
		YamlOutputType: cdk8s.YamlOutputType_FILE_PER_CHART,
		// YamlOutputType: cdk8s.YamlOutputType_FOLDER_PER_CHART_FILE_PER_RESOURCE,
		// RecordConstructMetadata: jsii.Bool(true),
	})
	return app
}

// func patchObjects(node constructs.Node) {
// 	if node == nil {
// 		return
// 	}
// 	for _, child := range *node.Children() {
// 		patchObjects(child.Node())
// 		if !*cdk8s.ApiObject_IsApiObject(child) {
// 			continue
// 		}
// 		obj := cdk8s.ApiObject_Of(child)
// 		// fmt.Println(*obj.Kind(), *obj.Metadata().Name())
// 		if *obj.Kind() == "Deployment" || *obj.Kind() == "DaemonSet" || *obj.Kind() == "StatefulSet" {
// 			obj.Metadata().AddAnnotation(jsii.String("kapp.k14s.io/update-strategy"), jsii.String("fallback-on-replace"))
// 		}
// 	}
// }

func NewKappConfig(scope constructs.Construct) constructs.Construct {
	chart := cdk8s.NewChart(scope, jsii.String("kapp-config"), &cdk8s.ChartProps{})
	obj := cdk8s.NewApiObject(chart, jsii.String("config"), &cdk8s.ApiObjectProps{
		ApiVersion: jsii.String("kapp.k14s.io/v1alpha1"),
		Kind:       jsii.String("Config"),
	})
	obj.AddJsonPatch(cdk8s.JsonPatch_Add(jsii.String("/rebaseRules"), []any{
		map[string]any{
			"path":    []string{"data"},
			"type":    "copy",
			"sources": []any{"new", "existing"},
			"resourceMatchers": []any{
				map[string]any{
					"kindNamespaceNameMatcher": map[string]any{
						"kind":      "Secret",
						"namespace": "secrets",
						"name":      "external-secrets-webhook",
					},
				},
			},
		},
	}))
	return chart
}

func Synth(app cdk8s.App) {
	fmt.Println("Patching...")
	// for _, chart := range *app.Charts() {
	// 	patchObjects(chart.Node())
	// }
	NewKappConfig(app)

	if err := os.RemoveAll(*app.Outdir()); err != nil {
		log.Fatalf("RemoveAll: %v", err)
	}

	app.Synth()

	// // read all files and escape {{ }} to avoid templating by helm
	// files, err := filepath.Glob(filepath.Join(*app.Outdir(), "*.yaml"))
	// if err != nil {
	// 	log.Fatalf("Glob: %v", err)
	// }
	// for _, file := range files {
	// 	bytes, err := os.ReadFile(file)
	// 	if err != nil {
	// 		log.Fatalf("ReadFile: %v", err)
	// 	}
	// 	replacer := strings.NewReplacer("{{", "{{\"{{\"}}", "}}", "{{\"}}\"}}")
	// 	bytes = []byte(replacer.Replace(string(bytes)))
	// 	// write back to file
	// 	if err := os.WriteFile(file, bytes, 0644); err != nil {
	// 		log.Fatalf("WriteFile: %v", err)
	// 	}
	// }
	log.Println("Done.")
}
