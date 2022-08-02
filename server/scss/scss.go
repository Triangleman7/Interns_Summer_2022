/*
Package scss handles writing SCSS Document (.SCSS) output.
*/
package scss

import (
	"io/ioutil"
	"os"
)

// ReadTemplate reads the template SCSS Stylesheet located at path and returns the content of the
// SCSS Stylesheet (content).
//
// Raises any errors encountered while reading the SCSS Stylesheet.
func ReadTemplate(path string) (content string, err error) {
	var reader []byte

	reader, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}

	return string(reader), nil
}

// WriteSCSS writes content, the content of the output SCSS Stylesheet, to a SCSS Stylesheet in the
// local file system at path.
//
// Raises any errors encountered while writing the SCSS Stylesheet contents to the target
// file.
func WriteSCSS(path string, content string) (err error) {
	var file *os.File

	file, err = os.Create(path)
	if err != nil {
		return
	}
	defer func() {
		var e = file.Close()
		if e != nil {
			panic(e)
		}
	}()

	_, err = file.Write([]byte(content))
	if err != nil {
		return
	}

	return
}
