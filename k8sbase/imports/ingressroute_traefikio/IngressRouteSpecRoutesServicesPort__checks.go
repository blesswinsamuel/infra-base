//go:build !no_runtime_type_checking

package ingressroute_traefikio

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

