// external-secretsio
package externalsecretsio


// ClusterExternalSecretSpec defines the desired state of ClusterExternalSecret.
type ClusterExternalSecretSpec struct {
	// The spec for the ExternalSecrets to be created.
	ExternalSecretSpec *ClusterExternalSecretSpecExternalSecretSpec `field:"required" json:"externalSecretSpec" yaml:"externalSecretSpec"`
	// The labels to select by to find the Namespaces to create the ExternalSecrets in.
	NamespaceSelector *ClusterExternalSecretSpecNamespaceSelector `field:"required" json:"namespaceSelector" yaml:"namespaceSelector"`
	// The name of the external secrets to be created defaults to the name of the ClusterExternalSecret.
	ExternalSecretName *string `field:"optional" json:"externalSecretName" yaml:"externalSecretName"`
	// The time in which the controller should reconcile it's objects and recheck namespaces for labels.
	RefreshTime *string `field:"optional" json:"refreshTime" yaml:"refreshTime"`
}

