// cert-managerio
package certmanagerio


// Cloud specifies the Venafi cloud configuration settings.
//
// Only one of TPP or Cloud may be specified.
type ClusterIssuerSpecVenafiCloud struct {
	// APITokenSecretRef is a secret key selector for the Venafi Cloud API token.
	ApiTokenSecretRef *ClusterIssuerSpecVenafiCloudApiTokenSecretRef `field:"required" json:"apiTokenSecretRef" yaml:"apiTokenSecretRef"`
	// URL is the base URL for Venafi Cloud.
	//
	// Defaults to "https://api.venafi.cloud/v1".
	Url *string `field:"optional" json:"url" yaml:"url"`
}

