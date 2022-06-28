package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Triangleman7/Interns_Summer_2022/api"
)

var PORT int = 8080

func main() {
	// Register handler functions
	http.HandleFunc("/", api.ProcessRootResponse)

	fmt.Printf("Listening on Localhost (Port %v)\n\n", PORT)	// Debuggin
	err := http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil)
	if err != nil { log.Fatal(err) }
}
