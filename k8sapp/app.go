package k8sapp

import (
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/blesswinsamuel/infra-base/kubegogen"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.TimeOnly})
}

func NewApp(props kubegogen.AppProps) kubegogen.App {
	app := kubegogen.NewApp(props)

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

func NewKappConfig(scope kubegogen.Scope) kubegogen.Scope {
	chart := scope.CreateScope("kapp-config", kubegogen.ScopeProps{})
	pvResourceMatchers := []any{
		map[string]any{
			"apiVersionKindMatcher": map[string]any{
				"apiVersion": "v1",
				"kind":       "PersistentVolume",
			},
		},
	}
	chart.AddApiObjectFromMap(map[string]interface{}{
		"apiVersion":             "kapp.k14s.io/v1alpha1",
		"kind":                   "Config",
		"minimumRequiredVersion": "0.23.0",
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
			// https://github.com/carvel-dev/kapp/issues/49
			// https://gist.github.com/cppforlife/149872f132d6afdc6f0240d70f598a16
			map[string]any{
				"paths": [][]string{
					{"spec", "claimRef"},
					{"spec", "claimRef", "resourceVersion"},
					{"spec", "claimRef", "uid"},
					{"spec", "claimRef", "apiVersion"},
					{"spec", "claimRef", "kind"},
				},
				"type":             "copy",
				"sources":          []string{"new", "existing"},
				"resourceMatchers": pvResourceMatchers,
			},
			// map[string]any{
			// 	"path":             []string{"spec", "persistentVolumeReclaimPolicy"},
			// 	"type":             "copy",
			// 	"sources":          []any{"new", "existing"},
			// 	"resourceMatchers": pvResourceMatchers,
			// },
			// map[string]any{
			// 	"path":             []string{"spec", "volumeMode"},
			// 	"type":             "copy",
			// 	"sources":          []any{"new", "existing"},
			// 	"resourceMatchers": pvResourceMatchers,
			// },
			// map[string]any{
			// 	"path":             []string{"metadata", "annotations", "pv.kubernetes.io/bound-by-controller"},
			// 	"type":             "copy",
			// 	"sources":          []any{"new", "existing"},
			// 	"resourceMatchers": pvResourceMatchers,
			// },
		},
	})
	return chart
}

func Synth(app kubegogen.App) {
	log.Info().Msg("Starting synth...")
	startTime := time.Now()
	NewKappConfig(app)

	app.Synth()

	log.Info().Msgf("Synth done in %s.", time.Since(startTime))
}

func TemplateOutputFiles(app kubegogen.App, vars any) {
	files, err := filepath.Glob(filepath.Join(app.OutDir(), "*.yaml"))
	if err != nil {
		log.Panic().Err(err).Msg("Glob")
	}
	for _, file := range files {
		bytes, err := os.ReadFile(file)
		if err != nil {
			log.Panic().Err(err).Msg("ReadFile")
		}
		tpl := template.New("tpl")
		tpl.Delims("[{@", "@}]")
		out, err := tpl.Parse(string(bytes))
		if err != nil {
			log.Panic().Err(err).Msg("Parse")
		}
		f, err := os.OpenFile(file, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Panic().Err(err).Msg("OpenFile")
		}
		out.Execute(f, vars)
		f.Close()
	}
}
