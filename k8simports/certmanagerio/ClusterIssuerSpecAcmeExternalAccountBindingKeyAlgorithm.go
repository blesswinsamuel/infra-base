package certmanagerio


// Deprecated: keyAlgorithm field exists for historical compatibility reasons and should not be used.
//
// The algorithm is now hardcoded to HS256 in golang/x/crypto/acme.
type ClusterIssuerSpecAcmeExternalAccountBindingKeyAlgorithm string

const (
	// HS256.
	ClusterIssuerSpecAcmeExternalAccountBindingKeyAlgorithm_HS256 ClusterIssuerSpecAcmeExternalAccountBindingKeyAlgorithm = "HS256"
	// HS384.
	ClusterIssuerSpecAcmeExternalAccountBindingKeyAlgorithm_HS384 ClusterIssuerSpecAcmeExternalAccountBindingKeyAlgorithm = "HS384"
	// HS512.
	ClusterIssuerSpecAcmeExternalAccountBindingKeyAlgorithm_HS512 ClusterIssuerSpecAcmeExternalAccountBindingKeyAlgorithm = "HS512"
)

