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
	// Assert that the target file is an HTML document
	if filepath.Ext(filename) != ".html" {
		fmt.Printf("Invalid HTML file name: '%v'", filename)
		return
	}

	// Join the output directory path to the target file name
	var path string = filepath.Join(OutputDirectory, filename)

	// Remove the target directory and its children
	os.RemoveAll(OutputDirectory)		// Returns `nil` if directory exists

	// Create an empty output directory
	err := os.Mkdir(OutputDirectory, PermissionBits)
	if err != nil { log.Fatal(err) }

	// Write `value` argument to determined file
	err = ioutil.WriteFile(path, []byte(value), PermissionBits)
	if err != nil { log.Fatal(err) }
}