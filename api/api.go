package api

import (
	"fmt"
	"net/http"
	"strings"
)

func ProcessRootResponse(w http.ResponseWriter, r *http.Request) {
	// Assert URL path directs to the root address
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		// Reply with root HTML document
		http.ServeFile(w, r, "ui/index.html")

	case "POST":
		// Catch errors thrown while parsing form submission
		err := r.ParseForm();
		if err != nil { fmt.Fprintf(w, "ParseForm() err: %v", err) }

		// Get value of text field
		valueTextField := r.FormValue("primary-form-text")

		// Format value of text field
		fvalueTextField := strings.ToUpper(valueTextField)

		// Debugging
		fmt.Printf("<form name=\"primary\">: Text Field value = \"%v\"\n", valueTextField)
		fmt.Printf("\tFormatted: \"%v\"\n", fvalueTextField)

	default:
		fmt.Fprintf(w, "Only GET and POST requests supported")
	}
}