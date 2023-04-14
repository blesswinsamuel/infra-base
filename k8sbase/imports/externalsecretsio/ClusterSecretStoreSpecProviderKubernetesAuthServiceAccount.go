// external-secretsio
package externalsecretsio


// points to a service account that should be used for authentication.
type ClusterSecretStoreSpecProviderKubernetesAuthServiceAccount struct {
	// A reference to a ServiceAccount resource.
	ServiceAccount *ClusterSecretStoreSpecProviderKubernetesAuthServiceAccountServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
}

