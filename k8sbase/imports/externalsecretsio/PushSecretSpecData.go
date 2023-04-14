// external-secretsio
package externalsecretsio


type PushSecretSpecData struct {
	// Match a given Secret Key to be pushed to the provider.
	Match *PushSecretSpecDataMatch `field:"required" json:"match" yaml:"match"`
}

