package docx

import (
	"fmt"

	"github.com/Triangleman7/Interns_Summer_2022/msword"
)

func Image(key int, document *msword.Docx, src string) {
	var field string = fmt.Sprintf("word/media/image%d.jpg", key)

	var err error = document.ReplaceImage(field, src)
	if err != nil { panic(err) }
}

func Paragraph(key string, document *msword.Docx, content string) {
	var field string = fmt.Sprintf("{%v}", key)

	var err error = document.Replace(field, content, -1)
	if err != nil { panic(err) }
}