package outputdata

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var OutputDirectory string = "out/"
var PermissionBits os.FileMode = 0755

func DirectorySetup() {
	// Remove the target directory and its children
	os.RemoveAll(OutputDirectory)		// Returns `nil` if directory exists

	// Create an empty output directory
	err := os.Mkdir(OutputDirectory, PermissionBits)
	if err != nil { log.Fatal(err) }
}

func assertFileExt(filename string, extension string) error {
	if filepath.Ext(filename) == extension { return nil }
	return errors.New(fmt.Sprintf("Expected file extension '%v' ('%v')", extension, filename))
}

func formatHTML(value string, formatter string) string {
	content, err := ioutil.ReadFile(formatter)
	if err != nil { log.Fatal(err) }

	return fmt.Sprintf(string(content), value)
}

func OutputHTML(filename string, value string) {
	// Assert that the target file is an HTML document
	err := assertFileExt(filename, ".html")
	if err != nil { log.Fatal(err) }

	// Construct the path to the target file name
	var path string = filepath.Join(OutputDirectory, filename)

	// Write `value` argument to determined file
	var valueFormatted string = formatHTML(value, "outputdata/output-template.html")
	err = ioutil.WriteFile(path, []byte(valueFormatted), PermissionBits)
	if err != nil { log.Fatal(err) }
}