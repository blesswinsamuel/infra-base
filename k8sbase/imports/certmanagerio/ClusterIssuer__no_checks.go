//go:build no_runtime_type_checking

// cert-managerio
package certmanagerio

// Building without runtime type checking enabled, so all the below just return nil

func validateClusterIssuer_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateClusterIssuer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateClusterIssuer_ManifestParameters(props *ClusterIssuerProps) error {
	return nil
}

func validateClusterIssuer_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewClusterIssuerParameters(scope constructs.Construct, id *string, props *ClusterIssuerProps) error {
	return nil
}

