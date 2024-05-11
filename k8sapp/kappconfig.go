package k8sapp

import "github.com/blesswinsamuel/infra-base/kubegogen"

func NewKappConfig(scope kubegogen.Scope) kubegogen.Scope {
	// https://carvel.dev/kapp/docs/v0.45.0/config/#rebaserules
	scope = scope.CreateScope("kapp-config", kubegogen.ScopeProps{})
	pvResourceMatchers := []map[string]any{
		{"apiVersionKindMatcher": map[string]any{"apiVersion": "v1", "kind": "PersistentVolume"}},
	}
	// https://github.com/carvel-dev/kapp/blob/2d0b7edbcd49a58263a37c48c5a614704c0d091f/pkg/kapp/config/resource_matchers.go#L12
	// https://carvel.dev/kapp/docs/v0.62.x/rebase-pvc/
	clusterOwnedFields := func(paths [][]string, resourceMatchers []map[string]any) map[string]any {
		return map[string]any{
			"paths":            paths,
			"type":             "copy",
			"sources":          []any{"new", "existing"},
			"resourceMatchers": resourceMatchers,
		}
	}
	scope.AddApiObjectFromMap(map[string]interface{}{
		"apiVersion":             "kapp.k14s.io/v1alpha1",
		"kind":                   "Config",
		"minimumRequiredVersion": "0.23.0",
		"rebaseRules": []any{
			// https://github.com/carvel-dev/kapp/issues/49
			// https://gist.github.com/cppforlife/149872f132d6afdc6f0240d70f598a16
			clusterOwnedFields([][]string{{"data"}}, []map[string]any{
				{"kindNamespaceNameMatcher": map[string]any{"kind": "Secret", "namespace": "secrets", "name": "external-secrets-webhook"}},
				{"kindNamespaceNameMatcher": map[string]any{"kind": "Secret", "namespace": "system", "name": "kubernetes-dashboard-csrf"}},
				{"kindNamespaceNameMatcher": map[string]any{"kind": "Secret", "namespace": "system", "name": "kubernetes-dashboard-key-holder"}},
			}),
			clusterOwnedFields([][]string{{"spec", "template", "metadata", "annotations", "reloader.stakater.com/last-reloaded-from"}}, []map[string]any{
				{"apiVersionKindMatcher": map[string]any{"apiVersion": "apps/v1", "kind": "Deployment"}},
				{"apiVersionKindMatcher": map[string]any{"apiVersion": "apps/v1", "kind": "StatefulSet"}},
				{"apiVersionKindMatcher": map[string]any{"apiVersion": "apps/v1", "kind": "Daemonset"}},
			}),
			clusterOwnedFields([][]string{{"metadata", "annotations", "force-resync"}}, []map[string]any{
				{"apiVersionKindMatcher": map[string]any{"apiVersion": "external-secrets.io/v1beta1", "kind": "ExternalSecret"}},
			}),
			clusterOwnedFields([][]string{
				{"spec", "claimRef"},
				{"spec", "claimRef", "resourceVersion"},
				{"spec", "claimRef", "uid"},
				{"spec", "claimRef", "apiVersion"},
				{"spec", "claimRef", "kind"},
			}, pvResourceMatchers),
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
