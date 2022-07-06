package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Triangleman7/Interns_Summer_2022/outputdata"
)

var PORT int = 8080

var OUTPUTDIRECTORY string = "out/"
var TEMPDIRECTORY string = "temp/"
var PERMISSIONBITS os.FileMode = 0755

func DirectorySetup(dirpath string, permissions os.FileMode) {
	DirectoryTeardown(dirpath)

	// Create an empty output directory
	var err error = os.Mkdir(dirpath, permissions)
	if err != nil { panic(err) }
}

func DirectoryTeardown(dirpath string) {
	var err error = os.RemoveAll(dirpath)
	if err != nil { panic(err) }
}

func ProcessRootResponse(w http.ResponseWriter, r *http.Request) {
	var err error

	// Assert URL path directs to the root address
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {

	case "GET":
		// Reply with root HTML document
		http.ServeFile(w, r, "ui/index.html")

	case "POST":
		var res string

		// Catch errors thrown while parsing form submission
		err = r.ParseMultipartForm(0);
		if err != nil { panic(err) }

		// Get values from form submission
		var valueTextField string = r.FormValue("primary-text")
		var valueMenu string = r.FormValue("primary-text-operation")
		file, header, err := r.FormFile("primary-image")
		if err != nil { panic(err) }

		// Format value of text field
		err, res = formatValue(valueTextField, valueMenu)
		if err != nil { panic(err) }

		// Upload `file`
		var uploadpath string = UploadFile(file, header)

		// Debugging
		fmt.Fprintf(w, "<form name=\"primary\">: Text Field value = \"%v\"\n", valueTextField)
		fmt.Fprintf(w, "<form name=\"primary\">: Dropdown Menu value = \"%v\"\n", valueMenu)
		fmt.Fprintf(w, "<form name=\"primary\">: Image Field value = \"%v\"\n", header.Filename)
		fmt.Fprintf(w, "\tUploaded: \"%v\"\n", uploadpath)
		fmt.Fprintf(w, "\tFormatted: \"%v\"\n", res)

		// Write formatted value of text field to output files
		var htmlpath string = filepath.Join(OUTPUTDIRECTORY, "primary-text.html")
		var docpath string = filepath.Join(OUTPUTDIRECTORY, "primary-text.doc")
		outputdata.WriteOutput(htmlpath, "outputdata/templates/template.html", PERMISSIONBITS, res)
		outputdata.WriteOutput(docpath, "outputdata/templates/template.doc", PERMISSIONBITS, res)

		file.Close()

	default:
		fmt.Fprintf(w, "Only GET and POST requests supported")
	}
}