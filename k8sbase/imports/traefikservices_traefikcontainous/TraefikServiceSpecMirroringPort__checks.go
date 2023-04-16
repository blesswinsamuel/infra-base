//go:build !no_runtime_type_checking

// traefikservices_traefikcontainous
package traefikservices_traefikcontainous

import (
	"fmt"
)

func validateTraefikServiceSpecMirroringPort_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateTraefikServiceSpecMirroringPort_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}
