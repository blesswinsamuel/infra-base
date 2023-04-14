//go:build no_runtime_type_checking

// cert-managerio
package certmanagerio

// Building without runtime type checking enabled, so all the below just return nil

func validateCertificateRequest_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateCertificateRequest_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCertificateRequest_ManifestParameters(props *CertificateRequestProps) error {
	return nil
}

func validateCertificateRequest_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewCertificateRequestParameters(scope constructs.Construct, id *string, props *CertificateRequestProps) error {
	return nil
}

