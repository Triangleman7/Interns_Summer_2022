package main

import (
	"log"
	"net/http"

	"github.com/Triangleman7/Interns_Summer_2022/api"
)

func main() {
	http.HandleFunc("/", api.ProcessInput)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
