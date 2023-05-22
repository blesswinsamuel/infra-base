package externalsecretsio


type ClusterSecretStoreSpecProviderGcpsmAuthWorkloadIdentity struct {
	ClusterLocation *string `field:"required" json:"clusterLocation" yaml:"clusterLocation"`
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// A reference to a ServiceAccount resource.
	ServiceAccountRef *ClusterSecretStoreSpecProviderGcpsmAuthWorkloadIdentityServiceAccountRef `field:"required" json:"serviceAccountRef" yaml:"serviceAccountRef"`
	ClusterProjectId *string `field:"optional" json:"clusterProjectId" yaml:"clusterProjectId"`
}

