//go:build no_runtime_type_checking

package serverstransports_traefikio

// Building without runtime type checking enabled, so all the below just return nil

func validateServersTransport_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateServersTransport_IsConstructParameters(x interface{}) error {
	return nil
}

func validateServersTransport_ManifestParameters(props *ServersTransportProps) error {
	return nil
}

func validateServersTransport_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewServersTransportParameters(scope constructs.Construct, id *string, props *ServersTransportProps) error {
	return nil
}

