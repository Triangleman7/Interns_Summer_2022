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
		return
	}

	var path string = filepath.Join(OutputDirectory, filename)

	os.RemoveAll(OutputDirectory)
	err := os.Mkdir(OutputDirectory, PermissionBits)
	if err != nil { log.Fatal(err) }

	err = ioutil.WriteFile(path, []byte(value), PermissionBits)
	if err != nil { log.Fatal(err) }
}