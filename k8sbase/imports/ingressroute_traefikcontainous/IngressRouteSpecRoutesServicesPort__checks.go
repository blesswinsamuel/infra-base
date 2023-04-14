//go:build !no_runtime_type_checking

// ingressroute_traefikcontainous
package ingressroute_traefikcontainous

import (
	"fmt"
)

func validateIngressRouteSpecRoutesServicesPort_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateIngressRouteSpecRoutesServicesPort_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

