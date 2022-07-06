package outputdata

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Triangleman7/Interns_Summer_2022/outputdata/doc"
	"github.com/Triangleman7/Interns_Summer_2022/outputdata/html"
)

func GetTemplate(path string) (error, string) {
	var err error
	var content string
	var extension string = filepath.Ext(path)

	switch extension {
	case ".doc":
		err, content = nil, doc.ReadTemplate(path)
	case ".html":
		err, content = nil, html.ReadTemplate(path)
	default:
		err = errors.New(fmt.Sprintf("Unexpected file extension for Template File: '%v'", extension))
		content = ""
	}
	
	return err, content
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
	err, template = GetTemplate(templatepath)
	if err != nil { panic(err) }

	// Format `values` into format string
	var formatted string = fmt.Sprintf(template, values...)

	// Write formatted string to `targetpath`
	switch targetext {
	// Word Document (.DOC)
	case ".doc":
		err = ioutil.WriteFile(targetpath, []byte(formatted), mode)
		if err != nil { panic(err) }
	// HTML Document (.HTML)
	case ".html":
		err = ioutil.WriteFile(targetpath, []byte(formatted), mode)
		if err != nil { panic(err) }
	default:
		err = errors.New(fmt.Sprintf("Unexpected file extension for Target File: '%v'", targetext))
		panic(err)
	}
}
