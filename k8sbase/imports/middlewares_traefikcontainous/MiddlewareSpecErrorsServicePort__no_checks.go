//go:build no_runtime_type_checking

// middlewares_traefikcontainous
package middlewares_traefikcontainous

// Building without runtime type checking enabled, so all the below just return nil

func validateMiddlewareSpecErrorsServicePort_FromNumberParameters(value *float64) error {
	return nil
}

func validateMiddlewareSpecErrorsServicePort_FromStringParameters(value *string) error {
	return nil
}

