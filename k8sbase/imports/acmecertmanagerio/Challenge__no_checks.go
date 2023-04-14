//go:build no_runtime_type_checking

// acmecert-managerio
package acmecertmanagerio

// Building without runtime type checking enabled, so all the below just return nil

func validateChallenge_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateChallenge_IsConstructParameters(x interface{}) error {
	return nil
}

func validateChallenge_ManifestParameters(props *ChallengeProps) error {
	return nil
}

func validateChallenge_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewChallengeParameters(scope constructs.Construct, id *string, props *ChallengeProps) error {
	return nil
}

