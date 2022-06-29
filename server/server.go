package server

import (
	"fmt"
	"net/http"
	"strings"
)

var PORT int = 8080

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
		fmt.Fprintf(w, "<form name=\"primary\">: Text Field value = \"%v\"\n", valueTextField)
		fmt.Fprintf(w, "\tFormatted: \"%v\"\n", fvalueTextField)

		// Write formatted value of text field to HTML document
		var filename string = "primary-form-text.html"
		WriteHTML(filename, fvalueTextField)

	default:
		fmt.Fprintf(w, "Only GET and POST requests supported")
	}
}