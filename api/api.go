package api

import (
	"fmt"
	"net/http"
)

func ProcessInput(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "ui/index.html")
	case "POST":
		err := r.ParseForm();
		if err != nil {
			fmt.Println(w, "ParseForm() err: %v", err)
		}

		rawTextInput := r.FormValue("form-text-inp")
		fTextInput := CapitalizeInput(rawTextInput)

		fmt.Printf("form[name='form'] > form[name='form-text-input']:\n")
		fmt.Printf("\tRaw: %v\n", rawTextInput)
		fmt.Printf("\tFormatted: %v\n", fTextInput)
	}
}