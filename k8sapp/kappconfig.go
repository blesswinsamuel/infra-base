package k8sapp

import "github.com/blesswinsamuel/kgen"

// https://github.com/carvel-dev/kapp/blob/develop/pkg/kapp/config/resource_matchers.go
type ResourceMatchers []ResourceMatcher

type ResourceMatcher struct {
	AllMatcher               *AllMatcher               `json:"allMatcher,omitempty"`
	AnyMatcher               *AnyMatcher               `json:"anyMatcher,omitempty"`
	NotMatcher               *NotMatcher               `json:"notMatcher,omitempty"`
	AndMatcher               *AndMatcher               `json:"andMatcher,omitempty"`
	APIGroupKindMatcher      *APIGroupKindMatcher      `json:"apiGroupKindMatcher,omitempty"`
	APIVersionKindMatcher    *APIVersionKindMatcher    `json:"apiVersionKindMatcher,omitempty"`
	KindNamespaceNameMatcher *KindNamespaceNameMatcher `json:"kindNamespaceNameMatcher,omitempty"`
	HasAnnotationMatcher     *HasAnnotationMatcher     `json:"hasAnnotationMatcher,omitempty"`
	HasNamespaceMatcher      *HasNamespaceMatcher      `json:"hasNamespaceMatcher,omitempty"`
	CustomResourceMatcher    *CustomResourceMatcher    `json:"customResourceMatcher,omitempty"`
	EmptyFieldMatcher        *EmptyFieldMatcher        `json:"emptyFieldMatcher,omitempty"`
}

type AllMatcher struct{}

type AnyMatcher struct {
	Matchers []ResourceMatcher `json:"matchers,omitempty"`
}

type NotMatcher struct {
	Matcher ResourceMatcher `json:"matcher,omitempty"`
}

type AndMatcher struct {
	Matchers []ResourceMatcher `json:"matchers,omitempty"`
}

type APIGroupKindMatcher struct {
	APIGroup string `json:"apiGroup,omitempty"`
	Kind     string `json:"kind,omitempty"`
}

type APIVersionKindMatcher struct {
	APIVersion string `json:"apiVersion,omitempty"`
	Kind       string `json:"kind,omitempty"`
}

type KindNamespaceNameMatcher struct {
	Kind      string `json:"kind,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name,omitempty"`
}

type HasAnnotationMatcher struct {
	Keys []string `json:"keys,omitempty"`
}

type HasNamespaceMatcher struct {
	Names []string `json:"names,omitempty"`
}

type CustomResourceMatcher struct{}

type EmptyFieldMatcher struct {
	Path string `json:"path,omitempty"`
	// Path ctlres.Path
}

func NewKappConfig(scope kgen.Scope) kgen.Scope {
	// https://carvel.dev/kapp/docs/v0.45.0/config/#rebaserules
	scope = scope.CreateScope("kapp-config", kgen.ScopeProps{})

	// https://github.com/carvel-dev/kapp/blob/2d0b7edbcd49a58263a37c48c5a614704c0d091f/pkg/kapp/config/resource_matchers.go#L12
	// https://carvel.dev/kapp/docs/v0.62.x/rebase-pvc/
	clusterOwnedFields := func(paths [][]string, resourceMatchers any) map[string]any {
		return map[string]any{
			"paths":            paths,
			"type":             "copy",
			"sources":          []any{"existing", "new"},
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
			clusterOwnedFields([][]string{{"data"}}, []ResourceMatcher{
				{KindNamespaceNameMatcher: &KindNamespaceNameMatcher{Kind: "Secret", Namespace: "ingress", Name: "cilium-ca"}},
				{KindNamespaceNameMatcher: &KindNamespaceNameMatcher{Kind: "Secret", Namespace: "ingress", Name: "hubble-server-certs"}},

				{KindNamespaceNameMatcher: &KindNamespaceNameMatcher{Kind: "Secret", Namespace: "external-secrets", Name: "external-secrets-webhook"}},
				{KindNamespaceNameMatcher: &KindNamespaceNameMatcher{Kind: "Secret", Namespace: "secrets", Name: "external-secrets-webhook"}},

				{KindNamespaceNameMatcher: &KindNamespaceNameMatcher{Kind: "Secret", Namespace: "metallb", Name: "metallb-webhook-cert"}},

				{KindNamespaceNameMatcher: &KindNamespaceNameMatcher{Kind: "Secret", Namespace: "system", Name: "kubernetes-dashboard-csrf"}},
				{KindNamespaceNameMatcher: &KindNamespaceNameMatcher{Kind: "Secret", Namespace: "system", Name: "kubernetes-dashboard-key-holder"}},
			}),
			clusterOwnedFields([][]string{{"spec", "template", "metadata", "annotations", "reloader.stakater.com/last-reloaded-from"}}, []ResourceMatcher{
				{APIVersionKindMatcher: &APIVersionKindMatcher{APIVersion: "apps/v1", Kind: "Deployment"}},
				{APIVersionKindMatcher: &APIVersionKindMatcher{APIVersion: "apps/v1", Kind: "StatefulSet"}},
				{APIVersionKindMatcher: &APIVersionKindMatcher{APIVersion: "apps/v1", Kind: "Daemonset"}},
			}),
			clusterOwnedFields([][]string{{"metadata", "annotations", "force-resync"}}, []map[string]any{
				{"apiVersionKindMatcher": map[string]any{"apiVersion": "external-secrets.io/v1beta1", "kind": "ExternalSecret"}},
			}),
			// clusterOwnedFields([][]string{{"spec", "conversion", "webhook", "clientConfig", "caBundle"}}, []map[string]any{
			// 	{"kindNamespaceNameMatcher": map[string]any{"kind": "CustomResourceDefinition", "namespace": "", "name": "bgppeers.metallb.io"}},
			// }),
			clusterOwnedFields([][]string{{"spec", "conversion", "webhook", "clientConfig", "caBundle"}}, []map[string]any{
				{"apiVersionKindMatcher": map[string]any{"apiVersion": "apiextensions.k8s.io/v1", "kind": "CustomResourceDefinition"}},
			}),
			clusterOwnedFields([][]string{
				// {"spec", "claimRef"},
				{"spec", "claimRef", "resourceVersion"},
				{"spec", "claimRef", "uid"},
				{"spec", "claimRef", "apiVersion"},
				{"spec", "claimRef", "kind"},
			}, []map[string]any{
				{"apiVersionKindMatcher": map[string]any{"apiVersion": "v1", "kind": "PersistentVolume"}},
			}),
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
