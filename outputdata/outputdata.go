package outputdata

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func readTemplate(path string) (error, string) {
	var err error
	var reader []byte
	var content string

	// Read format string from `path`
	var extension string = filepath.Ext(path)
	switch extension {
	// Word Document (.DOC), HTML Document (.HTML)
	case ".doc", ".html":
		reader, err = ioutil.ReadFile(path)
		if err != nil { panic(err) }

		content = string(reader)

	default:
		err = errors.New(fmt.Sprintf("Unexpected file extension for Template File: '%v'", extension))
		return err, ""
	}

	return nil, content
}

func WriteOutput(targetpath string, templatepath string, mode os.FileMode, values ...interface{}) {
	var err error

	// Assert that the target file and template file have the same file extension
	var targetext string = filepath.Ext(targetpath)
	var templateext string = filepath.Ext(templatepath)
	if targetext != templateext {
		err = errors.New(fmt.Sprintf("File extension mismatch (Target File: '%v'; Template File: '%v')", targetext, templateext))
		panic(err)
	}

	// Read format string from `templatepath`
	var template string
	err, template = readTemplate(templatepath)
	if err != nil { panic(err) }

	// Format `values` into format string
	var formatted string = fmt.Sprintf(template, values...)

	// Write formatted string to `targetpath`
	switch targetext {
	// Word Document (.DOC), HTML Document(.HTML)
	case ".doc", ".html":
		err = ioutil.WriteFile(targetpath, []byte(formatted), mode)
		if err != nil { panic(err) }

	default:
		err = errors.New(fmt.Sprintf("Unexpected file extension for Target File: '%v'", targetext))
		panic(err)
	}
}
