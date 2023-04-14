//go:build no_runtime_type_checking

// cert-managerio
package certmanagerio

// Building without runtime type checking enabled, so all the below just return nil

func validateIssuer_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateIssuer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIssuer_ManifestParameters(props *IssuerProps) error {
	return nil
}

func validateIssuer_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewIssuerParameters(scope constructs.Construct, id *string, props *IssuerProps) error {
	return nil
}

