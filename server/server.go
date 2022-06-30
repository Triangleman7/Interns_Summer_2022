package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Triangleman7/Interns_Summer_2022/outputdata"
)

var PORT int = 8080

func formatValue(input string, method string) string {
	switch method {
	case "":			// 'No Operation'
		return input

	case "lower":		// 'Lowercase'
		return strings.ToLower(input)

	case "upper":		// 'Uppercase'
		return strings.ToUpper(input)

	default:
		return ""
	}
}

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

		var res string

		// Get value of text field
		valueTextField := r.FormValue("primary-text")

		// Get value of dropdown menu
		valueMenu := r.FormValue("primary-text-operation")

		// Format value of text field
		res = formatValue(valueTextField, valueMenu)

		// Debugging
		fmt.Fprintf(w, "<form name=\"primary\">: Text Field value = \"%v\"\n", valueTextField)
		fmt.Fprintf(w, "<form name=\"primary\">: Dropdown Menu value = \"%v\"\n", valueMenu)
		fmt.Fprintf(w, "\tFormatted: \"%v\"\n", res)

		// Write formatted value of text field to HTML document
		outputdata.OutputHTML("primary-text.html", "outputdata/templates/template.html", res)
		outputdata.OutputDOC("primary-text.doc", "outputdata/templates/template.doc", res)

	default:
		fmt.Fprintf(w, "Only GET and POST requests supported")
	}
}