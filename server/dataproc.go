package server

import (
	"fmt"
	"log"
	"strings"
)

// FormatValue applies to input the string operation corresponding to method. The modified string
// is returned as output.
//
// Supported values for method:
// - "": Performs no operation to input
// - "lower": Applies lowercase casing to input
// - "upper": Applies uppercase casing to input
// An error is raised if an unexpected value for method is passed.
func FormatValue(input string, method string) (output string, err error) {
	switch method {
	// 'No Operation'
	case "":
		output = input
	// 'Lowercase'
	case "lower":
		output = strings.ToLower(input)
	// 'Uppercase'
	case "upper":
		output = strings.ToUpper(input)
	// Invalid value passed
	default:
		err = fmt.Errorf("unexpected value for method passed: '%v'", input)
		return
	}

	log.Printf("Formatted \"%s\" (%s): \"%s\"", input, method, output)
	return
}
