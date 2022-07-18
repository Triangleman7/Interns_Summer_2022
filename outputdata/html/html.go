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

	return string(reader), nil
}

func WriteHTML(path string, content string) (err error) {
	var file *os.File

	file, err = os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	file.Write([]byte(content))

	return
}
