//go:build no_runtime_type_checking

// external-secretsio
package externalsecretsio

// Building without runtime type checking enabled, so all the below just return nil

func validateClusterSecretStoreV1Beta1_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateClusterSecretStoreV1Beta1_IsConstructParameters(x interface{}) error {
	return nil
}

func validateClusterSecretStoreV1Beta1_ManifestParameters(props *ClusterSecretStoreV1Beta1Props) error {
	return nil
}

func validateClusterSecretStoreV1Beta1_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewClusterSecretStoreV1Beta1Parameters(scope constructs.Construct, id *string, props *ClusterSecretStoreV1Beta1Props) error {
	return nil
}

