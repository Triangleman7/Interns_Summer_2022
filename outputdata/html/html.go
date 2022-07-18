// Package outputdata/html handles writing HTML Document (.HTML) output.
package html

import (
	"io/ioutil"
	"os"
)

// ReadTemplate reads the template HTML Document located at path and returns the content of the
// HTML Document (content).
//
// Raises any errors encountered while reading the HTML Document.
func ReadTemplate(path string) (content string, err error) {
	var reader []byte

	reader, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}

	return string(reader), nil
}

// WriteHTML writes content, the content of the output HTML Document, an HTML Document in the local
// file system at path.
//
// Raises any errors encountered while writing the Word Document object contents to the target
// file.
func WriteHTML(path string, content string) (err error) {
	var file *os.File

	file, err = os.Create(path)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = file.Write([]byte(content))
	if err != nil {
		return
	}

	return
}
