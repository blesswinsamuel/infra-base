//go:build !no_runtime_type_checking

// middlewares_traefikcontainous
package middlewares_traefikcontainous

import (
	"fmt"
)

func validateMiddlewareSpecRateLimitPeriod_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateMiddlewareSpecRateLimitPeriod_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

