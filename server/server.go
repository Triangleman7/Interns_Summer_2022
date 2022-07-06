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
		// Handle form submission: <form name="primary">
		HandleFormPrimary(w, r)

	default:
		fmt.Fprintf(w, "Only GET and POST requests supported")
	}
}