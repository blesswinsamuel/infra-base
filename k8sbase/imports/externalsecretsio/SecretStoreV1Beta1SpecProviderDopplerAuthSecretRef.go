// external-secretsio
package externalsecretsio


type SecretStoreV1Beta1SpecProviderDopplerAuthSecretRef struct {
	// The DopplerToken is used for authentication.
	//
	// See https://docs.doppler.com/reference/api#authentication for auth token types. The Key attribute defaults to dopplerToken if not specified.
	DopplerToken *SecretStoreV1Beta1SpecProviderDopplerAuthSecretRefDopplerToken `field:"required" json:"dopplerToken" yaml:"dopplerToken"`
}

