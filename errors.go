package kopeechka

import "fmt"

var (
	// MissingRequiredParameterError is returned when on or more required parameters are missing.
	MissingRequiredParameterError = fmt.Errorf("missing required parameters")
)
