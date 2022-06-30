package server

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Triangleman7/Interns_Summer_2022/outputdata"
)

var PORT int = 8080

func formatValue(input string, method string) (error, string) {
	var err error

	switch method {
	// 'No Operation'
	case "":
		return nil, input
	// 'Lowercase'
	case "lower":
		return nil, strings.ToLower(input)
	// 'Uppercase'
	case "upper":
		return nil, strings.ToUpper(input)
	// Failsafe, but should not be possible
	default:
		err = errors.New(fmt.Sprintf("Unexpected value received: '%v'", input))
		return err, ""
	}
}

func ProcessRootResponse(w http.ResponseWriter, r *http.Request) {
	var err error

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
		var buf bytes.Buffer
		var res string

		// Catch errors thrown while parsing form submission
		err = r.ParseMultipartForm(0);
		if err != nil { panic(err) }

		// Get values from form submission
		var valueTextField string = r.FormValue("primary-text")
		var valueMenu string = r.FormValue("primary-text-operation")
		file, header, err := r.FormFile("primary-image")
		if err != nil { panic(err) }

		// Process `file`
		io.Copy(&buf, file)

		// Format value of text field
		err, res = formatValue(valueTextField, valueMenu)
		if err != nil { panic(err) }

		// Debugging
		fmt.Fprintf(w, "<form name=\"primary\">: Text Field value = \"%v\"\n", valueTextField)
		fmt.Fprintf(w, "<form name=\"primary\">: Dropdown Menu value = \"%v\"\n", valueMenu)
		fmt.Fprintf(w, "<form name=\"primary\">: Image Field value = \"%v\"\n", header.Filename)
		fmt.Fprint(w, buf)
		fmt.Fprintf(w, "\tFormatted: \"%v\"\n", res)

		// Write formatted value of text field to output files
		outputdata.WriteOutput("primary-text.html", "outputdata/templates/template.html", res)
		outputdata.WriteOutput("primary-text.doc", "outputdata/templates/template.doc", res)

	default:
		fmt.Fprintf(w, "Only GET and POST requests supported")
	}
}