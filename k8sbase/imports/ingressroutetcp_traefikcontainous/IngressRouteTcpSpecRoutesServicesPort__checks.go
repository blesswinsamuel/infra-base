//go:build !no_runtime_type_checking

// ingressroutetcp_traefikcontainous
package ingressroutetcp_traefikcontainous

import (
	"fmt"
)

func validateIngressRouteTcpSpecRoutesServicesPort_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateIngressRouteTcpSpecRoutesServicesPort_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

