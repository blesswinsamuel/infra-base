package generatorsexternalsecretsio


// PasswordSpec controls the behavior of the password generator.
type PasswordSpec struct {
	// set AllowRepeat to true to allow repeating characters.
	AllowRepeat *bool `field:"required" json:"allowRepeat" yaml:"allowRepeat"`
	// Length of the password to be generated.
	//
	// Defaults to 24.
	Length *float64 `field:"required" json:"length" yaml:"length"`
	// Set NoUpper to disable uppercase characters.
	NoUpper *bool `field:"required" json:"noUpper" yaml:"noUpper"`
	// Digits specifies the number of digits in the generated password.
	//
	// If omitted it defaults to 25% of the length of the password.
	Digits *float64 `field:"optional" json:"digits" yaml:"digits"`
	// SymbolCharacters specifies the special characters that should be used in the generated password.
	SymbolCharacters *string `field:"optional" json:"symbolCharacters" yaml:"symbolCharacters"`
	// Symbols specifies the number of symbol characters in the generated password.
	//
	// If omitted it defaults to 25% of the length of the password.
	Symbols *float64 `field:"optional" json:"symbols" yaml:"symbols"`
}

