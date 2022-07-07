package outputdata

import (
	"errors"
	"fmt"
	"path/filepath"

)

func AssertTemplateMatch(targetpath string, templatepath string) error {
	var err error = nil

	var targetext string = filepath.Ext(targetpath)
	var templateext string = filepath.Ext(templatepath)

	if targetext != templateext {
		err = errors.New(fmt.Sprintf("File extension mismatch (Target File: '%v'; Template File: '%v')", targetext, templateext))
	}

	return err
}
