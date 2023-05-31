package certmanagerio


// Use the DigitalOcean DNS API to manage DNS01 challenge records.
type ClusterIssuerSpecAcmeSolversDns01Digitalocean struct {
	// A reference to a specific 'key' within a Secret resource.
	//
	// In some instances, `key` is a required field.
	TokenSecretRef *ClusterIssuerSpecAcmeSolversDns01DigitaloceanTokenSecretRef `field:"required" json:"tokenSecretRef" yaml:"tokenSecretRef"`
}
