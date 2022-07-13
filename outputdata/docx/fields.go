package docx

import (
	"fmt"

	"github.com/Triangleman7/Interns_Summer_2022/msword"
)

func Image(key int, document *msword.Docx, src string) (err error) {
	var field string = fmt.Sprintf("word/media/image%d.jpg", key)

	err = document.ReplaceImage(field, src)

	return
}

func Paragraph(key string, document *msword.Docx, content string) (err error) {
	var field string = fmt.Sprintf("{%v}", key)

	err = document.Replace(field, content, -1)

	return
}
