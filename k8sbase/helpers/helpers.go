package helpers

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

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

func NewApp(outputDir string) cdk8s.App {
	app := cdk8s.NewApp(&cdk8s.AppProps{
		Outdir:         jsii.String(outputDir),
		YamlOutputType: cdk8s.YamlOutputType_FILE_PER_CHART,
		// YamlOutputType: cdk8s.YamlOutputType_FOLDER_PER_CHART_FILE_PER_RESOURCE,
		// RecordConstructMetadata: jsii.Bool(true),
	})

	var cacheDir = os.Getenv("CACHE_DIR")
	if cacheDir == "" {
		cacheDir = "./cache"
	}

	k8sapp.SetGlobalContext(app, k8sapp.Globals{
		DefaultSecretStoreName:               "secretstore",
		DefaultSecretStoreKind:               "ClusterSecretStore",
		DefaultExternalSecretRefreshInterval: "10m",
		DefaultCertIssuerName:                "letsencrypt-prod",
		DefaultCertIssuerKind:                "ClusterIssuer",
		CacheDir:                             cacheDir,
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
				map[string]any{
					"kindNamespaceNameMatcher": map[string]any{
						"kind":      "Secret",
						"namespace": "system",
						"name":      "kubernetes-dashboard-csrf",
					},
				},
				map[string]any{
					"kindNamespaceNameMatcher": map[string]any{
						"kind":      "Secret",
						"namespace": "system",
						"name":      "kubernetes-dashboard-key-holder",
					},
				},
			},
		},
	}))
	return chart
}

func Synth(app cdk8s.App) {
	log.Println("Starting synth...")
	startTime := time.Now()
	NewKappConfig(app)

	if err := os.RemoveAll(*app.Outdir()); err != nil {
		log.Fatalf("RemoveAll: %v", err)
	}

	app.Synth()

	log.Printf("Synth done in %s.", time.Since(startTime))
}

func TemplateOutputFiles(app cdk8s.App, vars any) {
	files, err := filepath.Glob(filepath.Join(*app.Outdir(), "*.yaml"))
	if err != nil {
		log.Fatalf("Glob: %v", err)
	}
	for _, file := range files {
		bytes, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("ReadFile: %v", err)
		}
		tpl := template.New("tpl")
		tpl.Delims("[{@", "@}]")
		out, err := tpl.Parse(string(bytes))
		if err != nil {
			log.Fatalf("Parse: %v", err)
		}
		f, err := os.OpenFile(file, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatalf("OpenFile: %v", err)
		}
		out.Execute(f, vars)
		f.Close()
	}
}

type IngressAnnotationsProps struct {
	Group       string
	Name        string
	Icon        string
	Description string
}

func HomepageIngressAnnotations(props IngressAnnotationsProps) map[string]string {
	return map[string]string{
		"gethomepage.dev/description": props.Description,
		"gethomepage.dev/enabled":     "true",
		"gethomepage.dev/group":       props.Group,
		"gethomepage.dev/icon":        props.Icon,
		"gethomepage.dev/name":        props.Name,
	}
}
