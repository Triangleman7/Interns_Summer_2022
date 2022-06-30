package outputdata

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func ReadHTMLTemplate(path string) string {
	reader, err := ioutil.ReadFile(path)
	if err != nil { panic(err) }

	var content string = string(reader)
	return content
}

func OutputHTML(targetpath string, templatepath string, values ...interface{}) {
	var err error

	// Assert that the target file is an HTML document
	if filepath.Ext(targetpath) != ".html" {
		err = errors.New(fmt.Sprintf("Expecte file extension '.html' (%v)", targetpath))
		panic(err)
	}

	// Construct relative path to the target file
	var path string = filepath.Join(OutputDirectory, targetpath)

	// Format `values` into format string read from `templatepath`
	var template string = ReadHTMLTemplate(templatepath)
	var formatted string = fmt.Sprintf(template, values...)

	// Write formatted string to `targetpath`
	err = ioutil.WriteFile(path, []byte(formatted), PermissionBits)
	if err != nil { panic(err) }
}
