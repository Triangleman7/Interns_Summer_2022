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

		// Tear down the temporary directory
		server.DirectoryTeardown(server.TEMPDIRECTORY)

		// Exit the program
		os.Exit(0)
	}()
}

func main() {
	var err error

	// Set up temporary directory
	server.DirectorySetup(server.TEMPDIRECTORY, server.PERMISSIONBITS)

	// Set up output directory
	server.DirectorySetup(server.OUTPUTDIRECTORY, server.PERMISSIONBITS)

	// Create a listener on a new goroutine to handle OS interrupts
	SetupCloseHandler()

	// Register handler functions
	http.HandleFunc("/", server.ProcessRootResponse)

	var fs http.Handler
	fs = http.FileServer(http.Dir("ui"))
	http.Handle("/ui/", http.StripPrefix("/ui/", fs))
	fs = http.FileServer(http.Dir("temp"))
	http.Handle("/temp/", http.StripPrefix("/temp/", fs))

	// Debugging
	fmt.Printf("Listening on Localhost (Port %v)\n\n", server.PORT)

	err = http.ListenAndServe(fmt.Sprintf(":%v", server.PORT), nil)
	if err != nil { log.Fatal(err) }
}
