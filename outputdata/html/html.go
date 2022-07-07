package html

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Triangleman7/Interns_Summer_2022/outputdata"
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

func WriteHTML (targetpath string, templatepath string, mode os.FileMode, content string) {
	var err error

	err = outputdata.AssertTemplateMatch(targetpath, templatepath)
	if err != nil { panic(err) }

	var template string
	err, template = ReadTemplate(templatepath)
	if err != nil { panic(err) }

	err = ioutil.WriteFile(targetpath, []byte(fmt.Sprintf(template, content)), mode)
	if err != nil { panic(err) }
}