package server

import (
	"fmt"
	"log"
	"strings"
)

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
	// Failsafe; Theoretically should not be possible
	default:
		err = fmt.Errorf("unexpected value received: '%v'", input)
		return
	}

	log.Printf("Formatted \"%s\" (%s): \"%s\"", input, method, output)
	return
}
