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


func ImgMargin(align string) (margin string) {
	switch align {
	// Left alignment
	case "left":
		return "0 auto 0 0"
	// Right alignment
	case "right":
		return "0 0 0 auto"
	// Center alignment
	case "center":
		return "0 auto"
	// Default alignment
	default:
		return "inherit"
	}
}

func PFontStyle(italic bool) (fontStyle string) {
	if italic {
		return "italic"
	} else {
		return "normal"
	}
}

func PFontWeight(bold bool) (fontWeight string) {
	if bold {
		return "bold"
	} else {
		return "normal"
	}
}

func PTextDecoration(strikethrough bool, underline bool) (textDecoration string) {
	if strikethrough && underline {
		return "line-through underline"
	} else if strikethrough {
		return "line-through"
	} else if underline {
		return "underline"
	} else {
		return "normal"
	}
}

