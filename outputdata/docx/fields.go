package docx

import (
	"fmt"

	"github.com/Triangleman7/Interns_Summer_2022/msword"
)

func Paragraph(key string, document *msword.Docx, content string) {
	var field string = fmt.Sprintf("{%v}", key)

	document.ReplaceRaw(field, content, -1)
}