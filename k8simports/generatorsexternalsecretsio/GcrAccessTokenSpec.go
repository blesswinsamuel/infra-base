package generatorsexternalsecretsio


type GcrAccessTokenSpec struct {
	// Auth defines the means for authenticating with GCP.
	Auth *GcrAccessTokenSpecAuth `field:"required" json:"auth" yaml:"auth"`
	// ProjectID defines which project to use to authenticate with.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

