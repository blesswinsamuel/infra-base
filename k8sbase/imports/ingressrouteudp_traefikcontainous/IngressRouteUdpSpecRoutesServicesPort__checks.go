//go:build !no_runtime_type_checking

// ingressrouteudp_traefikcontainous
package ingressrouteudp_traefikcontainous

import (
	"fmt"
)

func validateIngressRouteUdpSpecRoutesServicesPort_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateIngressRouteUdpSpecRoutesServicesPort_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

