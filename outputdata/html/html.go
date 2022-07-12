package html

import (
	"io/ioutil"
	"os"
)

func ReadTemplate(path string) (error, string) {
	var err error
	var reader []byte
	var content string

	reader, err = ioutil.ReadFile(path)
	if err != nil { panic(err) }

	content = string(reader)

	return nil, content
}

func WriteHTML(targetpath string, mode os.FileMode, content string) {
	var err error

	err = ioutil.WriteFile(targetpath, []byte(content), mode)
	if err != nil { panic(err) }
}