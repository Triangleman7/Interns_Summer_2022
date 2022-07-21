package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Triangleman7/Interns_Summer_2022/server"
)

func SetupCloseHandler() {
	channel := make(chan os.Signal)

	// Notify chanel if an interrupt from the OS is received
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)

	go func() {
		// Listen for interrupt from the OS
		<-channel
		log.Print("Detected OS interrupt")

		// Tear down the temporary directory
		server.DirectoryTeardown(server.TEMPDIRECTORY)

		// Exit the program
		log.Print("Terminating program (status code 0)")
		os.Exit(0)
	}()
}

func main() {
	log.Print("Starting program")

	var err error

	// Set up temporary directory
	server.DirectorySetup(server.TEMPDIRECTORY, server.PERMISSIONBITS)

	// Set up output directory
	server.DirectorySetup(server.OUTPUTDIRECTORY, server.PERMISSIONBITS)

	// Create a listener on a new goroutine to handle OS interrupts
	SetupCloseHandler()

	// Register handler functions
	http.HandleFunc("/", server.ProcessRootResponse)

	// Serve necessary directories
	var serveDirectories []string = []string{"client", "temp"}
	for _, directory := range serveDirectories {
		var fs http.Handler = http.FileServer(http.Dir(directory))
		var prefix string = fmt.Sprintf("/%v/", directory)
		http.Handle(prefix, http.StripPrefix(prefix, fs))
		log.Printf("Served %v/ directory (%v)", directory, prefix)
	}

	// Debugging
	log.Printf("Listening on Localhost (Port %v)", server.PORT)

	err = http.ListenAndServe(fmt.Sprintf(":%v", server.PORT), nil)
	if err != nil {
		log.Fatal(err)
	}
}
