package certmanagerio


// CredentialsRef is a reference to a Secret containing the username and password for the TPP server.
//
// The secret must contain two keys, 'username' and 'password'.
type ClusterIssuerSpecVenafiTppCredentialsRef struct {
	// Name of the resource being referred to.
	//
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `field:"required" json:"name" yaml:"name"`
}

