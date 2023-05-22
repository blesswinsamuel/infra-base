package externalsecretsio


// Match a given Secret Key to be pushed to the provider.
type PushSecretSpecDataMatch struct {
	// Remote Refs to push to providers.
	RemoteRef *PushSecretSpecDataMatchRemoteRef `field:"required" json:"remoteRef" yaml:"remoteRef"`
	// Secret Key to be pushed.
	SecretKey *string `field:"required" json:"secretKey" yaml:"secretKey"`
}

