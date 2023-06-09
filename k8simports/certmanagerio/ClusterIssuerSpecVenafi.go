package certmanagerio


// Venafi configures this issuer to sign certificates using a Venafi TPP or Venafi Cloud policy zone.
type ClusterIssuerSpecVenafi struct {
	// Zone is the Venafi Policy Zone to use for this issuer.
	//
	// All requests made to the Venafi platform will be restricted by the named zone policy. This field is required.
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// Cloud specifies the Venafi cloud configuration settings.
	//
	// Only one of TPP or Cloud may be specified.
	Cloud *ClusterIssuerSpecVenafiCloud `field:"optional" json:"cloud" yaml:"cloud"`
	// TPP specifies Trust Protection Platform configuration settings.
	//
	// Only one of TPP or Cloud may be specified.
	Tpp *ClusterIssuerSpecVenafiTpp `field:"optional" json:"tpp" yaml:"tpp"`
}

