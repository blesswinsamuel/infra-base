//go:build no_runtime_type_checking

// middlewarestcp_traefikcontainous
package middlewarestcp_traefikcontainous

// Building without runtime type checking enabled, so all the below just return nil

func validateMiddlewareTcp_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateMiddlewareTcp_IsConstructParameters(x interface{}) error {
	return nil
}

func validateMiddlewareTcp_ManifestParameters(props *MiddlewareTcpProps) error {
	return nil
}

func validateMiddlewareTcp_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewMiddlewareTcpParameters(scope constructs.Construct, id *string, props *MiddlewareTcpProps) error {
	return nil
}

