//go:build no_runtime_type_checking

// middlewares_traefikcontainous
package middlewares_traefikcontainous

// Building without runtime type checking enabled, so all the below just return nil

func validateMiddleware_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateMiddleware_IsConstructParameters(x interface{}) error {
	return nil
}

func validateMiddleware_ManifestParameters(props *MiddlewareProps) error {
	return nil
}

func validateMiddleware_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewMiddlewareParameters(scope constructs.Construct, id *string, props *MiddlewareProps) error {
	return nil
}

