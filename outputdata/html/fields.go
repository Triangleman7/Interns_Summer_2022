package html

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Image replaces all instances of the format field in output identified by key with an <img>
// element that is sourced at src. Since output is a pointer, the replacement is performed
// in-place.
func Image(key string, output *string, src string) {
	var field string = fmt.Sprintf("{%v}", key)

	var tag string = fmt.Sprintf("<img src=\"%v\">", filepath.Join("..", src))

	*output = strings.ReplaceAll(*output, field, tag)
}

// Paragraph replaces all instances of the format filed in output identified by key with a <p>
// element that has content as the innter text. Since output is a pointer, the replacement is
// performed in-place.
func Paragraph(key string, output *string, content string) {
	var field string = fmt.Sprintf("{%v}", key)

	var tag string = fmt.Sprintf("<p>%v</p>", content)

	*output = strings.ReplaceAll(*output, field, tag)
}