/*
Package outputdata/docx handles writing Word Document (.DOCX) output.
*/
package docx

import (
	"github.com/Triangleman7/Interns_Summer_2022/msword"
)

// ReadTemplate reads the template Word Document located at path and returns the a Word Document
// object (document).
//
// Raises any errors encountered while reading the Word Document.
func ReadTemplate(path string) (document *msword.Docx, err error) {
	var reader *msword.ReplaceDocx

	reader, err = msword.ReadDocxFile(path)
	if err != nil {
		return
	}
	defer reader.Close()

	return reader.Editable(), nil
}

// WriteDOCX writes the Word Document object document to a Word Document in the local file system
// at path.
//
// Raises any errors encountered while writing the Word Document object contents to the target
// file.
func WriteDOCX(path string, document *msword.Docx) (err error) {
	err = document.WriteToFile(path)

	return
}
