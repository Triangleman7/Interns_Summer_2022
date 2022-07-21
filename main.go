/*

 */

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Triangleman7/Interns_Summer_2022/server"
)

func main() {
	log.Print("Starting program")

	var err error

	// Set up temporary directory
	server.DirectorySetup(server.TEMPDIRECTORY, server.FILEMODE)

	// Set up output directory
	server.DirectorySetup(server.OUTPUTDIRECTORY, server.FILEMODE)

	// Create a listener on a new goroutine to handle OS interrupts
	server.SetupCloseHandler()

	// Register handler functions
	http.HandleFunc("/", server.ProcessRootResponse)

	// Serve necessary directories
	var serveDirectories []string = []string{"client", "temp"}
	for _, directory := range serveDirectories {
		var fs http.Handler = http.FileServer(http.Dir(directory))
		var prefix string = fmt.Sprintf("/%s/", directory)
		http.Handle(prefix, http.StripPrefix(prefix, fs))
		log.Printf("Served %s/ directory (%s)", directory, prefix)
	}

	// Debugging
	log.Printf("Listening on Localhost (Port %d)", server.PORT)

	err = http.ListenAndServe(fmt.Sprintf(":%d", server.PORT), nil)
	if err != nil {
		log.Fatal(err)
	}
}
