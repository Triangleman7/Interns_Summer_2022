package server

import (
	"errors"
	"log"
	"net/http"
	"os"
)

var PORT int = 8080

var OUTPUTDIRECTORY string = "out/"
var TEMPDIRECTORY string = "temp/"
var PERMISSIONBITS os.FileMode = 0755

func DirectorySetup(dirpath string, permissions os.FileMode) {
	log.Printf("Setting up directory at %v", dirpath)

	// Remove target directory
	DirectoryTeardown(dirpath)

	// Create an empty output directory
	var err error = os.Mkdir(dirpath, permissions)
	if err != nil {
		panic(err)
	}
	log.Printf("Directory created at %v", dirpath)
}

func DirectoryTeardown(dirpath string) {
	log.Printf("Tearing down directory at %v", dirpath)

	var err error = os.RemoveAll(dirpath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Directory removed at %v", dirpath)
}

func ProcessRootResponse(w http.ResponseWriter, r *http.Request) {
	var err error

	// Assert URL path directs to the root address
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	log.Printf("Received request: %s", r.Method)
	log.Printf("Request headers: %v", r.Header)
	switch r.Method {

	case "GET":
		// Reply with root HTML document
		http.ServeFile(w, r, "client/index.html")

	case "POST":
		// Handle form submission: <form name="primary">
		err = HandleFormPrimary(w, r)

	default:
		err = errors.New("only GET and POST requests supported")
	}

	if err != nil {
		log.Panic(err)
	}
}
