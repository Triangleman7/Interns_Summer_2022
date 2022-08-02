package scss

import (
	"fmt"
	"strings"
)

// Rules appends to the end of the SCSS-valid string a new rule for selector, with the CSS
// property-value determined by content.
func Rule(output *string, selector string, content map[string]string) {
	var styles = make([]string, len(content))
	for property, value := range content {
		styles = append(styles, fmt.Sprintf("%s: %s;", property, value))
	}

	var rule = fmt.Sprintf("%s {%s}", selector, strings.Join(styles, ""))

	*output = strings.Join([]string{*output, rule}, "\n\n")
}

// ImgMargin returns the value for a <img> 'margin' CSS property from the alignment determined by
// align.
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

// PFontStyle returns the value for a <p> 'font-style' CSS property from whether the text should be
// italicized, as determined by italic.
func PFontStyle(italic bool) (fontStyle string) {
	if italic {
		return "italic"
	} else {
		return "normal"
	}
}

// PFontWeight returns the value for a <p> 'font-weight' CSS property from whether the text should
// be bolded, as determined by bold.
func PFontWeight(bold bool) (fontWeight string) {
	if bold {
		return "bold"
	} else {
		return "normal"
	}
}

// PTextDecoration returns the value for a <p> 'text-decoration' CSS property from whether the text
// should be struck through and/or underlined, as determined by strikethrough and underline,
// respectively.
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
