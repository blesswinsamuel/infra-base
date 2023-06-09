package certmanagerio


// KeyUsage specifies valid usage contexts for keys.
//
// See: https://tools.ietf.org/html/rfc5280#section-4.2.1.3 https://tools.ietf.org/html/rfc5280#section-4.2.1.12
// Valid KeyUsage values are as follows: "signing", "digital signature", "content commitment", "key encipherment", "key agreement", "data encipherment", "cert sign", "crl sign", "encipher only", "decipher only", "any", "server auth", "client auth", "code signing", "email protection", "s/mime", "ipsec end system", "ipsec tunnel", "ipsec user", "timestamping", "ocsp signing", "microsoft sgc", "netscape sgc".
type CertificateRequestSpecUsages string

const (
	// signing.
	CertificateRequestSpecUsages_SIGNING CertificateRequestSpecUsages = "SIGNING"
	// digital signature.
	CertificateRequestSpecUsages_DIGITAL_SIGNATURE CertificateRequestSpecUsages = "DIGITAL_SIGNATURE"
	// content commitment.
	CertificateRequestSpecUsages_CONTENT_COMMITMENT CertificateRequestSpecUsages = "CONTENT_COMMITMENT"
	// key encipherment.
	CertificateRequestSpecUsages_KEY_ENCIPHERMENT CertificateRequestSpecUsages = "KEY_ENCIPHERMENT"
	// key agreement.
	CertificateRequestSpecUsages_KEY_AGREEMENT CertificateRequestSpecUsages = "KEY_AGREEMENT"
	// data encipherment.
	CertificateRequestSpecUsages_DATA_ENCIPHERMENT CertificateRequestSpecUsages = "DATA_ENCIPHERMENT"
	// cert sign.
	CertificateRequestSpecUsages_CERT_SIGN CertificateRequestSpecUsages = "CERT_SIGN"
	// crl sign.
	CertificateRequestSpecUsages_CRL_SIGN CertificateRequestSpecUsages = "CRL_SIGN"
	// encipher only.
	CertificateRequestSpecUsages_ENCIPHER_ONLY CertificateRequestSpecUsages = "ENCIPHER_ONLY"
	// decipher only.
	CertificateRequestSpecUsages_DECIPHER_ONLY CertificateRequestSpecUsages = "DECIPHER_ONLY"
	// any.
	CertificateRequestSpecUsages_ANY CertificateRequestSpecUsages = "ANY"
	// server auth.
	CertificateRequestSpecUsages_SERVER_AUTH CertificateRequestSpecUsages = "SERVER_AUTH"
	// client auth.
	CertificateRequestSpecUsages_CLIENT_AUTH CertificateRequestSpecUsages = "CLIENT_AUTH"
	// code signing.
	CertificateRequestSpecUsages_CODE_SIGNING CertificateRequestSpecUsages = "CODE_SIGNING"
	// email protection.
	CertificateRequestSpecUsages_EMAIL_PROTECTION CertificateRequestSpecUsages = "EMAIL_PROTECTION"
	// s/mime.
	CertificateRequestSpecUsages_S_MIME CertificateRequestSpecUsages = "S_MIME"
	// ipsec end system.
	CertificateRequestSpecUsages_IPSEC_END_SYSTEM CertificateRequestSpecUsages = "IPSEC_END_SYSTEM"
	// ipsec tunnel.
	CertificateRequestSpecUsages_IPSEC_TUNNEL CertificateRequestSpecUsages = "IPSEC_TUNNEL"
	// ipsec user.
	CertificateRequestSpecUsages_IPSEC_USER CertificateRequestSpecUsages = "IPSEC_USER"
	// timestamping.
	CertificateRequestSpecUsages_TIMESTAMPING CertificateRequestSpecUsages = "TIMESTAMPING"
	// ocsp signing.
	CertificateRequestSpecUsages_OCSP_SIGNING CertificateRequestSpecUsages = "OCSP_SIGNING"
	// microsoft sgc.
	CertificateRequestSpecUsages_MICROSOFT_SGC CertificateRequestSpecUsages = "MICROSOFT_SGC"
	// netscape sgc.
	CertificateRequestSpecUsages_NETSCAPE_SGC CertificateRequestSpecUsages = "NETSCAPE_SGC"
)

