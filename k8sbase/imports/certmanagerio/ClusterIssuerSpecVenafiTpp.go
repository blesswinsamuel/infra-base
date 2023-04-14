// cert-managerio
package certmanagerio


// TPP specifies Trust Protection Platform configuration settings.
//
// Only one of TPP or Cloud may be specified.
type ClusterIssuerSpecVenafiTpp struct {
	// CredentialsRef is a reference to a Secret containing the username and password for the TPP server.
	//
	// The secret must contain two keys, 'username' and 'password'.
	CredentialsRef *ClusterIssuerSpecVenafiTppCredentialsRef `field:"required" json:"credentialsRef" yaml:"credentialsRef"`
	// URL is the base URL for the vedsdk endpoint of the Venafi TPP instance, for example: "https://tpp.example.com/vedsdk".
	Url *string `field:"required" json:"url" yaml:"url"`
	// Base64-encoded bundle of PEM CAs which will be used to validate the certificate chain presented by the TPP server.
	//
	// Only used if using HTTPS; ignored for HTTP. If undefined, the certificate bundle in the cert-manager controller container is used to validate the chain.
	CaBundle *string `field:"optional" json:"caBundle" yaml:"caBundle"`
}

