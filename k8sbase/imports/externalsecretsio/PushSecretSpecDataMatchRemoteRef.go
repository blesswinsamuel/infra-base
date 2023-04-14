// external-secretsio
package externalsecretsio


// Remote Refs to push to providers.
type PushSecretSpecDataMatchRemoteRef struct {
	// Name of the resulting provider secret.
	RemoteKey *string `field:"required" json:"remoteKey" yaml:"remoteKey"`
}

