package doc

import (
	"io/ioutil"
	"os"

	"github.com/Triangleman7/Interns_Summer_2022/outputdata"
)

func WriteDOC(targetpath string, templatepath string, mode os.FileMode, content []byte) {
	var err error

	err = outputdata.AssertTemplateMatch(targetpath, templatepath)
	if err != nil { panic(err) }

	err = ioutil.WriteFile(targetpath, content, mode)
	if err != nil { panic(err) }
}
