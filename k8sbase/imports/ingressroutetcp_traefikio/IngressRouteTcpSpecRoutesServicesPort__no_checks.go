//go:build no_runtime_type_checking

package ingressroutetcp_traefikio

// Building without runtime type checking enabled, so all the below just return nil

func validateIngressRouteTcpSpecRoutesServicesPort_FromNumberParameters(value *float64) error {
	return nil
}

func validateIngressRouteTcpSpecRoutesServicesPort_FromStringParameters(value *string) error {
	return nil
}

