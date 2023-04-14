// cert-managerio
package certmanagerio


// ExternalAccountBinding is a reference to a CA external account of the ACME server.
//
// If set, upon registration cert-manager will attempt to associate the given external account credentials with the registered ACME account.
type IssuerSpecAcmeExternalAccountBinding struct {
	// keyID is the ID of the CA key that the External Account is bound to.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
	// keySecretRef is a Secret Key Selector referencing a data item in a Kubernetes Secret which holds the symmetric MAC key of the External Account Binding.
	//
	// The `key` is the index string that is paired with the key data in the Secret and should not be confused with the key data itself, or indeed with the External Account Binding keyID above. The secret key stored in the Secret **must** be un-padded, base64 URL encoded data.
	KeySecretRef *IssuerSpecAcmeExternalAccountBindingKeySecretRef `field:"required" json:"keySecretRef" yaml:"keySecretRef"`
	// Deprecated: keyAlgorithm field exists for historical compatibility reasons and should not be used.
	//
	// The algorithm is now hardcoded to HS256 in golang/x/crypto/acme.
	KeyAlgorithm IssuerSpecAcmeExternalAccountBindingKeyAlgorithm `field:"optional" json:"keyAlgorithm" yaml:"keyAlgorithm"`
}

