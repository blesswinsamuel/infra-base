// external-secretsio
package externalsecretsio


// Select a Secret to Push.
type PushSecretSpecSelectorSecret struct {
	// Name of the Secret.
	//
	// The Secret must exist in the same namespace as the PushSecret manifest.
	Name *string `field:"required" json:"name" yaml:"name"`
}

