package middlewares_traefikio


// PassTLSClientCert holds the pass TLS client cert middleware configuration.
//
// This middleware adds the selected data from the passed client TLS certificate to a header. More info: https://doc.traefik.io/traefik/v2.10/middlewares/http/passtlsclientcert/
type MiddlewareSpecPassTlsClientCert struct {
	// Info selects the specific client certificate details you want to add to the X-Forwarded-Tls-Client-Cert-Info header.
	Info *MiddlewareSpecPassTlsClientCertInfo `field:"optional" json:"info" yaml:"info"`
	// PEM sets the X-Forwarded-Tls-Client-Cert header with the certificate.
	Pem *bool `field:"optional" json:"pem" yaml:"pem"`
}

