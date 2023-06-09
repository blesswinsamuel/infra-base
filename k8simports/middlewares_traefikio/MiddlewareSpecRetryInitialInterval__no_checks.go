//go:build no_runtime_type_checking

package middlewares_traefikio

// Building without runtime type checking enabled, so all the below just return nil

func validateMiddlewareSpecRetryInitialInterval_FromNumberParameters(value *float64) error {
	return nil
}

func validateMiddlewareSpecRetryInitialInterval_FromStringParameters(value *string) error {
	return nil
}

