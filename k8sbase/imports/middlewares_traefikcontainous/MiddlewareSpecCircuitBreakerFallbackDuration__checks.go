//go:build !no_runtime_type_checking

// middlewares_traefikcontainous
package middlewares_traefikcontainous

import (
	"fmt"
)

func validateMiddlewareSpecCircuitBreakerFallbackDuration_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateMiddlewareSpecCircuitBreakerFallbackDuration_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

