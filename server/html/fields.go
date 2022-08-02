package html

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Image replaces all instances of the format field in output identified by key with an <img>
// element that is sourced at src. Since output is a pointer, the replacement is performed
// in-place.
func Image(output *string, key string, src string) {
	var field = fmt.Sprintf("{%s}", key)

	var tag = fmt.Sprintf(
		"<img class=\"%s\" src=\"%s\">",
		key,
		filepath.Join(RELROOT, src),
	)

	*output = strings.ReplaceAll(*output, field, tag)
}

// Paragraph replaces all instances of the format filed in output identified by key with a <p>
// element that has content as the innter text. Since output is a pointer, the replacement is
// performed in-place.
func Paragraph(output *string, key string, content string) {
	var field = fmt.Sprintf("{%s}", key)

	var tag = fmt.Sprintf(
		"<p class=\"%s\">%s</p>",
		key,
		content,
	)

	*output = strings.ReplaceAll(*output, field, tag)
}
