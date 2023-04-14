//go:build no_runtime_type_checking

// ingressroute_traefikcontainous
package ingressroute_traefikcontainous

// Building without runtime type checking enabled, so all the below just return nil

func validateIngressRoute_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateIngressRoute_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIngressRoute_ManifestParameters(props *IngressRouteProps) error {
	return nil
}

func validateIngressRoute_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewIngressRouteParameters(scope constructs.Construct, id *string, props *IngressRouteProps) error {
	return nil
}

