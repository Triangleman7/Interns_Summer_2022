package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Triangleman7/Interns_Summer_2022/server"
	"github.com/Triangleman7/Interns_Summer_2022/outputdata"
)

func main() {
	// Register handler functions
	http.HandleFunc("/", server.ProcessRootResponse)
	fs := http.FileServer(http.Dir("ui"))
	http.Handle("/ui/", http.StripPrefix("/ui/", fs))

	// Initialize output directory
	outputdata.DirectorySetup()

	// Debugging
	fmt.Printf("Listening on Localhost (Port %v)\n\n", server.PORT)

	err := http.ListenAndServe(fmt.Sprintf(":%v", server.PORT), nil)
	if err != nil { log.Fatal(err) }
}
