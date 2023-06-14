package k8sapp

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/blesswinsamuel/infra-base/packager"
)

func NewApp(outputDir string) packager.App {
	app := packager.NewApp(packager.AppProps{
		Outdir:       outputDir,
		DeleteOutDir: true,
	})

	var cacheDir = os.Getenv("CACHE_DIR")
	if cacheDir == "" {
		cacheDir = "./cache"
	}

	SetGlobalContext(app, Globals{
		DefaultSecretStoreName:               "secretstore",
		DefaultSecretStoreKind:               "ClusterSecretStore",
		DefaultExternalSecretRefreshInterval: "10m",
		DefaultCertIssuerName:                "letsencrypt-prod",
		DefaultCertIssuerKind:                "ClusterIssuer",
		CacheDir:                             cacheDir,
	})
	return app
}

func NewKappConfig(scope packager.Construct) packager.Construct {
	chart := scope.Chart("kapp-config", packager.ChartProps{})
	chart.ApiObjectFromMap("config", map[string]interface{}{
		"apiVersion": "kapp.k14s.io/v1alpha1",
		"kind":       "Config",
		"rebaseRules": []any{
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
		},
	})
	return chart
}

func Synth(app packager.App) {
	log.Println("Starting synth...")
	startTime := time.Now()
	NewKappConfig(app)

	app.Synth()

	log.Printf("Synth done in %s.", time.Since(startTime))
}

func TemplateOutputFiles(app packager.App, vars any) {
	files, err := filepath.Glob(filepath.Join(app.OutDir(), "*.yaml"))
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
