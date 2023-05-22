//go:build no_runtime_type_checking

package ingressrouteudp_traefikio

// Building without runtime type checking enabled, so all the below just return nil

func validateIngressRouteUdpSpecRoutesServicesPort_FromNumberParameters(value *float64) error {
	return nil
}

func validateIngressRouteUdpSpecRoutesServicesPort_FromStringParameters(value *string) error {
	return nil
}

