/*
Package server provides support for hosting, running, and handling a local server
*/
package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	PORT int = 8080 // the localhost port served

	OUTPUTDIRECTORY   string = "out/"             // holds all output files generated at runtime
	TEMPLATEDIRECTORY string = "server/templates" // holds all output template files
	TEMPDIRECTORY     string = "temp/"            // holds all temporary files generated at runtime

	FILEMODE os.FileMode = 0755 // the program-specific defualt permission bits
)

// DirectorySetup creates an empty directory at dirpath using the file mode defined by mode.
// DirectoryTeardown is first called to remove the target directory and its contents, if it already
// exists.
//
// Raises any errors encountered while creating the new directory.
func DirectorySetup(dirpath string, mode os.FileMode) {
	log.Printf("Setting up directory at %s", dirpath)

	// Remove target directory
	DirectoryTeardown(dirpath)

	// Create an empty output directory
	var err = os.Mkdir(dirpath, mode)
	if err != nil {
		panic(err)
	}
	log.Printf("Directory created at %s", dirpath)
}

// DirectoryTeardown removes the directory at dirpath and all its contents, regardless of whether
// the directory exists or no.
//
// Raises any errors encountered while removing the directory.
func DirectoryTeardown(dirpath string) {
	log.Printf("Tearing down directory at %s", dirpath)

	var err = os.RemoveAll(dirpath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Directory removed at %s", dirpath)
}

// SetupCloseHandler cleans up the local file system, called when the program is about to be
// terminated. A goroutine listens for OS interrupt errors and calls DirectoryTeardown if one is
// detected.
func SetupCloseHandler() {
	channel := make(chan os.Signal)

	// Notify chanel if an interrupt from the OS is received
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)

	go func() {
		// Listen for interrupt from the OS
		<-channel
		log.Print("Detected OS interrupt")

		// Tear down the temporary directory
		DirectoryTeardown(TEMPDIRECTORY)

		// Exit the program
		log.Print("Terminating program (status code 0)")
		os.Exit(0)
	}()
}

func requestCheck(w http.ResponseWriter, r *http.Request, path string) {
	if r.URL.Path != path {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	log.Printf("%s - Received request: %s", path, r.Method)
	log.Printf("%s - Request headers: %v", path, r.Header)
}

// ProcessRootRequest processes and responds to all requests made to the root URL ("/"). Only GET
// requests are supported; raises an error if other request types are received.
func ProcessRootRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	var path = "/"
	requestCheck(w, r, path)

	switch r.Method {

	case "GET":
		http.ServeFile(w, r, "client/index.html")

	default:
		err = errors.New("only GET requests supported")
	}

	if err != nil {
		log.Panic(err)
	}
}

// ProcessFormPrimaryRequest processes and responds to all POST requests made to the form
// `form#primary` ("/forms/primary"). Only POST requests are supported; raises an error if other
// request types are received.
//
// Raises any errors encountered while handling POST requests.
func ProcessFormPrimaryRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	var path = "/forms/primary"
	requestCheck(w, r, path)

	switch r.Method {

	case "POST":
		var formPrimary FormPrimary
		err = formPrimary.handle(w, r)

	default:
		err = errors.New("only POST requests supported")
	}

	payload := make(map[string]interface{})
	payload["success"] = (err == nil)

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(payload)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Responding to request: %v", payload)
}

func ProcessFormPrimaryZIPRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	var path = "/forms/primary/zip"
	requestCheck(w, r, path)

	var zippath string

	switch r.Method {

	case "POST":
		var form = Form{"form-primary"}
		zippath, err = form.ZIPOutputDirectory()

	default:
		err = errors.New("only POST requests supported")
	}

	if err != nil {
		log.Panic(err)
	}

	w.Header().Set("Content-Type", "octet/stream")

	var reader []byte
	reader, err = ioutil.ReadFile(zippath)
	if err != nil {
		log.Panic(err)
	}
	w.Write(reader)
}
