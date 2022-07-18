// Package outputdata/docx handles writing Word Document (.DOCX) output.
package docx

import (
	"github.com/Triangleman7/Interns_Summer_2022/msword"
)

// WriteDOCX writes the Word Document object document to a Word Document in the local file system
// at path.
//
// Raises any errors encountered while writing the Word Document object contents to the target
// file.
func WriteDOCX(path string, document *msword.Docx) (err error) {
	err = document.WriteToFile(path)

	return
}
