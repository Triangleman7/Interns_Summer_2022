package docx

import (
	"fmt"

	"../../resources/msword"
)

// Image replaces the image in document partially identified by key with the image at path src in
// the local file system.
//
// Raises any errors encountered while replacing the old image in the document with the new image.
func Image(document *msword.Docx, key int, src string) {
	var err error
	var field string

	// Using .jpg file extension
	field = fmt.Sprintf("word/media/image%d.jpg", key)
	err = document.ReplaceImage(field, src)
	if err != nil {
		panic(err)
	}

	// Using .jpeg file extension
	field = fmt.Sprintf("word/media/image%d.jpeg", key)
	err = document.ReplaceImage(field, src)
	if err != nil {
		panic(err)
	}
}

// Paragraph replaces all instances of key found in the body text of document with content.
//
// Raises any errors encountered while replacing the body text in the document.
func Paragraph(document *msword.Docx, key string, content string) {
	var err error
	var field = fmt.Sprintf("{%v}", key)

	err = document.Replace(field, content, -1)
	if err != nil {
		panic(err)
	}
}
