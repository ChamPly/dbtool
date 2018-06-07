package conversion

// SQLConversionRegister SqlConversion register method
var SQLConversionRegister map[string]IConversion

func init() {
	SQLConversionRegister = map[string]IConversion{}

	SQLConversionRegister["mysql"] = NewMySQLConversion()
}
