package scss

import (
	"fmt"
	"strings"
)

// Rule
func Rule(output *string, selector string, content map[string]string) {
	var styles []string = make([]string, len(content))
	for key, value := range content {
		styles = append(styles, fmt.Sprintf("%s: %s;", key, value))
	}

	var rule string = fmt.Sprintf("%s {%s\n}", selector, strings.Join(styles, "\n\t"))

	*output = strings.Join([]string{*output, rule}, "\n\n")
}
