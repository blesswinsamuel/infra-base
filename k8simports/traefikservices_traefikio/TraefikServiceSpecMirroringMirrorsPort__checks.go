//go:build !no_runtime_type_checking

package traefikservices_traefikio

import (
	"fmt"
)

func validateTraefikServiceSpecMirroringMirrorsPort_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateTraefikServiceSpecMirroringMirrorsPort_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

