/*
Package server provides support for hosting, running, and handling a local server
*/
package server

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	PORT int = 8080		// the localhost port served

	OUTPUTDIRECTORY string = "out/"							// holds all output files generated at runtime
	TEMPLATEDIRECTORY string = "outputdata/templates"		// holds all output template files
	TEMPDIRECTORY string = "temp/"							// holds all temporary files generated at runtime

	FILEMODE os.FileMode = 0755		// the program-specific defualt permission bits
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
	var err error = os.Mkdir(dirpath, mode)
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

	var err error = os.RemoveAll(dirpath)
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

// ProcessRootResponse processes and responds to all requests made to the root url ("/"). Only GET
// and POST requests are accepted and raises an error if other requests are received.
//
// Raises any errors encountered while handling POST requests.
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
		// Serve the root HTML document
		http.ServeFile(w, r, "client/index.html")

	case "POST":
		// Handle form submission to element form#primary
		err = HandleFormPrimary(w, r)

	default:
		err = errors.New("only GET and POST requests supported")
	}

	if err != nil {
		log.Panic(err)
	}
}
