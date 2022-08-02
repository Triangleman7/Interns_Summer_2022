/*
Package docx handles writing Word Document (.DOCX) output.
*/
package docx

import (
	"../../resources/msword"
)

// ReadTemplate reads the template Word Document located at path and returns a Word Document
// object (document).
//
// Raises any errors encountered while reading the Word Document.
func ReadTemplate(path string) (reader *msword.ReplaceDocx, document *msword.Docx, err error) {
	reader, err = msword.ReadDocxFile(path)
	if err != nil {
		return
	}

	document = reader.Editable()
	return
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
