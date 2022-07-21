package html

import (
	"io/ioutil"
	"os"
)

func ReadTemplate(path string) (content string, err error) {
	var reader []byte

	reader, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}

	content = string(reader)

	return
}

func WriteHTML(targetpath string, mode os.FileMode, content string) (err error) {
	err = ioutil.WriteFile(targetpath, []byte(content), mode)

	return
}
