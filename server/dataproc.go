package server

import (
	"errors"
	"fmt"
	"strings"
)

func FormatValue(input string, method string) (error, string) {
	var err error

	switch method {
	// 'No Operation'
	case "":
		return nil, input
	// 'Lowercase'
	case "lower":
		return nil, strings.ToLower(input)
	// 'Uppercase'
	case "upper":
		return nil, strings.ToUpper(input)
	// Failsafe; Theoretically should not be possible
	default:
		err = errors.New(fmt.Sprintf("Unexpected value received: '%v'", input))
		return err, ""
	}
}
