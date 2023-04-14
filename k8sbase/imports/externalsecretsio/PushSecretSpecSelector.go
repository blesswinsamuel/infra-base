// external-secretsio
package externalsecretsio


// The Secret Selector (k8s source) for the Push Secret.
type PushSecretSpecSelector struct {
	// Select a Secret to Push.
	Secret *PushSecretSpecSelectorSecret `field:"required" json:"secret" yaml:"secret"`
}

