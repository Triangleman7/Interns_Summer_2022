package doc

import (
	"io/ioutil"
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

