// tlsoptions_traefikcontainous
package tlsoptions_traefikcontainous


// ClientAuthType defines the client authentication type to apply.
type TlsOptionSpecClientAuthClientAuthType string

const (
	// NoClientCert.
	TlsOptionSpecClientAuthClientAuthType_NO_CLIENT_CERT TlsOptionSpecClientAuthClientAuthType = "NO_CLIENT_CERT"
	// RequestClientCert.
	TlsOptionSpecClientAuthClientAuthType_REQUEST_CLIENT_CERT TlsOptionSpecClientAuthClientAuthType = "REQUEST_CLIENT_CERT"
	// RequireAnyClientCert.
	TlsOptionSpecClientAuthClientAuthType_REQUIRE_ANY_CLIENT_CERT TlsOptionSpecClientAuthClientAuthType = "REQUIRE_ANY_CLIENT_CERT"
	// VerifyClientCertIfGiven.
	TlsOptionSpecClientAuthClientAuthType_VERIFY_CLIENT_CERT_IF_GIVEN TlsOptionSpecClientAuthClientAuthType = "VERIFY_CLIENT_CERT_IF_GIVEN"
	// RequireAndVerifyClientCert.
	TlsOptionSpecClientAuthClientAuthType_REQUIRE_AND_VERIFY_CLIENT_CERT TlsOptionSpecClientAuthClientAuthType = "REQUIRE_AND_VERIFY_CLIENT_CERT"
)

