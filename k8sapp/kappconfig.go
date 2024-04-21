package k8sapp

import "github.com/blesswinsamuel/infra-base/kubegogen"

func NewKappConfig(scope kubegogen.Scope) kubegogen.Scope {
	scope = scope.CreateScope("kapp-config", kubegogen.ScopeProps{})
	pvResourceMatchers := []any{
		map[string]any{
			"apiVersionKindMatcher": map[string]any{
				"apiVersion": "v1",
				"kind":       "PersistentVolume",
			},
		},
	}
	scope.AddApiObjectFromMap(map[string]interface{}{
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
	return scope
}
