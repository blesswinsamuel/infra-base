//go:build no_runtime_type_checking

package ingressroute_traefikio

// Building without runtime type checking enabled, so all the below just return nil

func validateIngressRouteSpecRoutesServicesPort_FromNumberParameters(value *float64) error {
	return nil
}

func validateIngressRouteSpecRoutesServicesPort_FromStringParameters(value *string) error {
	return nil
}

