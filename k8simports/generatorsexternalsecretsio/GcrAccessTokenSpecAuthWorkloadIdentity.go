package generatorsexternalsecretsio


type GcrAccessTokenSpecAuthWorkloadIdentity struct {
	ClusterLocation *string `field:"required" json:"clusterLocation" yaml:"clusterLocation"`
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// A reference to a ServiceAccount resource.
	ServiceAccountRef *GcrAccessTokenSpecAuthWorkloadIdentityServiceAccountRef `field:"required" json:"serviceAccountRef" yaml:"serviceAccountRef"`
	ClusterProjectId *string `field:"optional" json:"clusterProjectId" yaml:"clusterProjectId"`
}

