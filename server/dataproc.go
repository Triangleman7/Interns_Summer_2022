package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var OutputDirectory string = "out/"
var PermissionBits os.FileMode = 0755

func WriteHTML(filename string, value string) {
	if filepath.Ext(filename) != ".html" {
		fmt.Printf("Invalid HTML file name: '%v'", filename)
	}

	var path string = filepath.Join(OutputDirectory, filename)

	err := os.Mkdir(OutputDirectory, PermissionBits)
	if err != nil && os.IsExist(err) { os.RemoveAll(OutputDirectory) }

	err = ioutil.WriteFile(path, []byte(value), PermissionBits)
	if err != nil { log.Fatal(err) }
}