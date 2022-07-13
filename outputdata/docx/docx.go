package docx

import (
	"github.com/Triangleman7/Interns_Summer_2022/msword"
)

func WriteDOCX(targetpath string, document *msword.Docx) (err error) {
	err = document.WriteToFile(targetpath)

	return
}
