package server

import (
	"log"
	"regexp"
	"strings"
	"unicode"
)

func lowercase(str string) (s string) {
	return strings.ToLower(str)
}

func uppercase(str string) (s string) {
	return strings.ToUpper(str)
}

func camelcase(str string) (s string) {
	for i, w := range strings.Split(startcase(str), " ") {
		if i == 0 {
			s += strings.ToLower(w)
		} else {
			s += w
		}
	}
	return
}

func dotcase(str string) (s string) {
	return strings.Join(strings.Split(strings.ToLower(str), " "), ".")
}

func kebabcase(str string) (s string) {
	return strings.Join(strings.Split(strings.ToLower(str), " "), "-")
}

func oppositecase(str string) (s string) {
	for _, c := range str {
		if unicode.IsLower(c) {
			s += string(unicode.ToUpper(c))
		} else if unicode.IsUpper(c) {
			s += string(unicode.ToLower(c))
		} else {
			s += string(c)
		}
	}
	return
}

func pascalcase(str string) (s string) {
	return strings.Join(strings.Split(startcase(str), " "), "")
}

func sarcasticcase(str string) (s string) {
	for i, c := range str {
		if i%2 == 0 {
			s += string(unicode.ToLower(c))
		} else {
			s += string(unicode.ToUpper(c))
		}
	}
	return
}

func snakecase(str string) (s string) {
	return strings.Join(strings.Split(strings.ToLower(str), " "), "_")
}

func startcase(str string) (s string) {
	var regex *regexp.Regexp
	regex, _ = regexp.Compile(`\b\w+\s*`)
	for _, w := range regex.FindAllString(str, -1) {
		for i, c := range w {
			if i == 0 {
				s += string(unicode.ToUpper(c))
			} else {
				s += string(unicode.ToLower(c))
			}
		}
	}
	return s
}

func traincase(str string) (s string) {
	return strings.Join(strings.Split(startcase(str), " "), "-")
}

// FormatValue applies to input the string operation corresponding to method. The modified string
// is returned as output.
func FormatValue(input string, method string) (output string) {
	switch method {
	// 'Lowercase'
	case "lower":
		output = lowercase(input)
	// 'Uppercase'
	case "upper":
		output = uppercase(input)
	// 'Camel Case'
	case "camel":
		output = camelcase(input)
	// 'Dot Case'
	case "dot":
		output = dotcase(input)
	// 'Kebab Case'
	case "kebab":
		output = kebabcase(input)
	// 'Opposite Case'
	case "opposite":
		output = oppositecase(input)
	// 'Pascal Case'
	case "pascal":
		output = pascalcase(input)
	// 'Sarcastic Case'
	case "sarcastic":
		output = sarcasticcase(input)
	// 'Snake Case'
	case "snake":
		output = snakecase(input)
	// 'Start Case'
	case "start":
		output = startcase(input)
	// 'Train Case'
	case "train":
		output = traincase(input)
	// 'No operation', or equivalent
	default:
		output = input
	}

	log.Printf("Formatted \"%s\" (%s): \"%s\"", input, method, output)
	return
}
