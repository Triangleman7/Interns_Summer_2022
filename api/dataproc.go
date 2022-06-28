package api

import (
	"strings"
	"unicode"
)

func CapitalizeInput(input string) string {
	var res string;

	for _, word := range strings.Split(input, ` `) {
		for idx, char := range word {
			if idx == 0 {
				res += string(unicode.ToUpper(char))
			} else {
				res += string(unicode.ToLower(char))
			}
		}

		res += string(` `)
	}

	return res
}