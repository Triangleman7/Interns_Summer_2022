package outputdata

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/nguyenthenguyen/docx"
)

func ReadDOCXTemplate(path string) string {
	reader, err := docx.ReadDocxFile(path)
	if err != nil { panic(err) }

	var content string = reader.GetContent()
	reader.Close()
	return content
}

func OutputDOCX(targetpath string, templatepath string, values ...interface{}) {
	var err error

	// Assert that the target file is a DOCX document
	if filepath.Ext(targetpath) != ".docx" {
		err = errors.New(fmt.Sprintf("Expected file extension '.docx' ('%v')", targetpath))
		panic(err)
	}

	// Construct relative path to the target file
	var path string = filepath.Join(OutputDirectory, targetpath)

	// Format `values` into format string read from `templatepath`
	var template string = ReadDOCXTemplate(templatepath)
	var formatted string = fmt.Sprintf(template, values...)

	// Create empty DOCX file at `targetpath`
	err = ioutil.WriteFile(path, []byte(""), PermissionBis)
	if err != nil { panic(err) }

	// Write formatted string to `targetpath`
	reader, err := docx.ReadDocxFile(path)
	if err != nil { panic(err) }
	writer = reader.Editable()
	writer.SetContent(formatted)
	reader.Close()
}