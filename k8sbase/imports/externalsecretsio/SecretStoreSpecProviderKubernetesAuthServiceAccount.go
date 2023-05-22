package externalsecretsio


// points to a service account that should be used for authentication.
type SecretStoreSpecProviderKubernetesAuthServiceAccount struct {
	// A reference to a ServiceAccount resource.
	ServiceAccount *SecretStoreSpecProviderKubernetesAuthServiceAccountServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
}

