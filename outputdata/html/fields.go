package html

import (
	"fmt"
)

func Image(res *string, src string) {
	var tag string = fmt.Sprintf("<img src=\"%v\">", src)
	*res += fmt.Sprintf("%v\n", tag)
}

func Paragraph(res *string, content string) {
	var tag string = fmt.Sprintf("<p>%v</p>", content)
	*res += fmt.Sprintf("%v\n", tag)
}