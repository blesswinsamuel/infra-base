//go:build !no_runtime_type_checking

package traefikservices_traefikio

import (
	"fmt"
)

func validateTraefikServiceSpecWeightedServicesPort_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateTraefikServiceSpecWeightedServicesPort_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

