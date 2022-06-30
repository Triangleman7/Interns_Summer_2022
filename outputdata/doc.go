package outputdata

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func ReadDOCTemplate(path string) string {
	reader, err := ioutil.ReadFile(path)
	if err != nil { panic(err) }

	var content string = string(reader)
	return content
}

func OutputDOC(targetpath string, templatepath string, values ...interface{}) {
	var err error

	// Assert that the target file is a DOC document
	if filepath.Ext(targetpath) != ".doc" {
		err = errors.New(fmt.Sprintf("Expected file extension '.doc' ('%v')", targetpath))
		panic(err)
	}

	// Construct relative path to the target file
	var path string = filepath.Join(OutputDirectory, targetpath)

	// Format `values` into format string read from `templatepath`
	var template string = ReadDOCTemplate(templatepath)
	var formatted string = fmt.Sprintf(template, values...)

	// Create empty DOC file at `targetpath`
	err = ioutil.WriteFile(path, []byte(formatted), PermissionBits)
	if err != nil { panic(err) }
}