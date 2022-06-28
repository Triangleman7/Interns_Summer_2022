package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Triangleman7/Interns_Summer_2022/server"
)

func main() {
	// Register handler functions
	http.HandleFunc("/", server.ProcessRootResponse)

	fmt.Printf("Listening on Localhost (Port %v)\n\n", server.PORT)		// Debugging
	err := http.ListenAndServe(fmt.Sprintf(":%v", server.PORT), nil)
	if err != nil { log.Fatal(err) }
}
