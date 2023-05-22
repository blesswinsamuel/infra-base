package externalsecretsio


type SecretStoreSpecProviderGcpsmAuthWorkloadIdentity struct {
	ClusterLocation *string `field:"required" json:"clusterLocation" yaml:"clusterLocation"`
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// A reference to a ServiceAccount resource.
	ServiceAccountRef *SecretStoreSpecProviderGcpsmAuthWorkloadIdentityServiceAccountRef `field:"required" json:"serviceAccountRef" yaml:"serviceAccountRef"`
	ClusterProjectId *string `field:"optional" json:"clusterProjectId" yaml:"clusterProjectId"`
}

