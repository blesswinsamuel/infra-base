// external-secretsio
package externalsecretsio


type SecretStoreV1Beta1SpecProviderGcpsmAuthWorkloadIdentity struct {
	ClusterLocation *string `field:"required" json:"clusterLocation" yaml:"clusterLocation"`
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// A reference to a ServiceAccount resource.
	ServiceAccountRef *SecretStoreV1Beta1SpecProviderGcpsmAuthWorkloadIdentityServiceAccountRef `field:"required" json:"serviceAccountRef" yaml:"serviceAccountRef"`
	ClusterProjectId *string `field:"optional" json:"clusterProjectId" yaml:"clusterProjectId"`
}

