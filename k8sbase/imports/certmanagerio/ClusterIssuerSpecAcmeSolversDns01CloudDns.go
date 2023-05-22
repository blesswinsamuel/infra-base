package certmanagerio


// Use the Google Cloud DNS API to manage DNS01 challenge records.
type ClusterIssuerSpecAcmeSolversDns01CloudDns struct {
	Project *string `field:"required" json:"project" yaml:"project"`
	// HostedZoneName is an optional field that tells cert-manager in which Cloud DNS zone the challenge record has to be created.
	//
	// If left empty cert-manager will automatically choose a zone.
	HostedZoneName *string `field:"optional" json:"hostedZoneName" yaml:"hostedZoneName"`
	// A reference to a specific 'key' within a Secret resource.
	//
	// In some instances, `key` is a required field.
	ServiceAccountSecretRef *ClusterIssuerSpecAcmeSolversDns01CloudDnsServiceAccountSecretRef `field:"optional" json:"serviceAccountSecretRef" yaml:"serviceAccountSecretRef"`
}

