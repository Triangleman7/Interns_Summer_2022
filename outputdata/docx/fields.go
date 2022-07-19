package docx

import (
	"fmt"

	"github.com/Triangleman7/Interns_Summer_2022/msword"
)

// Image replaces the image in document partially identified by key with the image at path src in
// the local file system.
//
// Raises any errors encountered while replacing the old image in the document with the new image.
func Image(document *msword.Docx, key int, src string) (err error) {
	var field string = fmt.Sprintf("word/media/image%d.jpg", key)

	err = document.ReplaceImage(field, src)

	return
}

// Paragraph replaces all instances of key found in the body text of document with content.
//
// Raises any errors encountered while replacing the body text in the document.
func Paragraph(document *msword.Docx, key string, content string) (err error) {
	var field string = fmt.Sprintf("{%v}", key)

	err = document.Replace(field, content, -1)

	return
}
