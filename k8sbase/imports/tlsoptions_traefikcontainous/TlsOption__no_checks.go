//go:build no_runtime_type_checking

// tlsoptions_traefikcontainous
package tlsoptions_traefikcontainous

// Building without runtime type checking enabled, so all the below just return nil

func validateTlsOption_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateTlsOption_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTlsOption_ManifestParameters(props *TlsOptionProps) error {
	return nil
}

func validateTlsOption_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewTlsOptionParameters(scope constructs.Construct, id *string, props *TlsOptionProps) error {
	return nil
}

