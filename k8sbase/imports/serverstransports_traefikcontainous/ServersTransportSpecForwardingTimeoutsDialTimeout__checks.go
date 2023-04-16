//go:build !no_runtime_type_checking

// serverstransports_traefikcontainous
package serverstransports_traefikcontainous

import (
	"fmt"
)

func validateServersTransportSpecForwardingTimeoutsDialTimeout_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateServersTransportSpecForwardingTimeoutsDialTimeout_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

