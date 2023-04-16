//go:build no_runtime_type_checking

// tlsstores_traefikcontainous
package tlsstores_traefikcontainous

// Building without runtime type checking enabled, so all the below just return nil

func validateTlsStore_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateTlsStore_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTlsStore_ManifestParameters(props *TlsStoreProps) error {
	return nil
}

func validateTlsStore_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewTlsStoreParameters(scope constructs.Construct, id *string, props *TlsStoreProps) error {
	return nil
}

