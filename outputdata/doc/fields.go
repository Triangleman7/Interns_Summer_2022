package doc

import (
	"io/ioutil"
)

func Image(res *[]byte, src string) {
	var err error
	var element []byte

	element, err = ioutil.ReadFile(src)
	if err != nil { panic(err) }

	element = append(element[:], []byte("\n")...)
	*res = append((*res)[:], element...)
}

func Paragraph(res *[]byte, content string) {
	var element []byte = []byte(content)

	element = append(element[:], []byte("\n")...)
	*res = append((*res)[:], element...)
}