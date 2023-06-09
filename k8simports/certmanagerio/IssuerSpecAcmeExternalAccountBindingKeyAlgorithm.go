package certmanagerio


// Deprecated: keyAlgorithm field exists for historical compatibility reasons and should not be used.
//
// The algorithm is now hardcoded to HS256 in golang/x/crypto/acme.
type IssuerSpecAcmeExternalAccountBindingKeyAlgorithm string

const (
	// HS256.
	IssuerSpecAcmeExternalAccountBindingKeyAlgorithm_HS256 IssuerSpecAcmeExternalAccountBindingKeyAlgorithm = "HS256"
	// HS384.
	IssuerSpecAcmeExternalAccountBindingKeyAlgorithm_HS384 IssuerSpecAcmeExternalAccountBindingKeyAlgorithm = "HS384"
	// HS512.
	IssuerSpecAcmeExternalAccountBindingKeyAlgorithm_HS512 IssuerSpecAcmeExternalAccountBindingKeyAlgorithm = "HS512"
)

