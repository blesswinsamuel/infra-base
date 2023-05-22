//go:build no_runtime_type_checking

package middlewares_traefikio

// Building without runtime type checking enabled, so all the below just return nil

func validateMiddlewareSpecCircuitBreakerRecoveryDuration_FromNumberParameters(value *float64) error {
	return nil
}

func validateMiddlewareSpecCircuitBreakerRecoveryDuration_FromStringParameters(value *string) error {
	return nil
}

