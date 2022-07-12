package html

import (
	"fmt"
	"path/filepath"
	"strings"
)

func Image(key string, output *string, src string) {
	var field string = fmt.Sprintf("{%v}", key)

	var tag string = fmt.Sprintf("<img src=\"%v\">", filepath.Join("..", src))

	*output = strings.ReplaceAll(*output, field, tag)
}

func Paragraph(key string, output *string, content string) {
	var field string = fmt.Sprintf("{%v}", key)

	var tag string = fmt.Sprintf("<p>%v</p>", content)

	*output = strings.ReplaceAll(*output, field, tag)
}