package server

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/Triangleman7/Interns_Summer_2022/outputdata"
	"github.com/Triangleman7/Interns_Summer_2022/outputdata/html"
)

func HandleFormPrimary(w http.ResponseWriter, r *http.Request) {
	var err error
	
	// var outDOC string = ""
	var outHTML string = ""

	// Parse form submission
	err = r.ParseMultipartForm(0);
	if err != nil { panic(err) }

	// 
	var vTextField, vMenu, fvTextField string
	vTextField = r.FormValue("primary-text")
	vMenu = r.FormValue("primary-text-operation")
	err, fvTextField = FormatValue(vTextField, vMenu)
	
	//
	var file multipart.File
	var header *multipart.FileHeader
	var uploadpath string
	file, header, err = r.FormFile("primary-image")
	if err != nil { panic(err) }
	uploadpath = filepath.Join("..", UploadFile(file, header))
	err = file.Close()
	if err != nil { panic(err) }

	// Debugging
	fmt.Fprintf(w, "<form name=\"primary\">: Text Field value = \"%v\"\n", vTextField)
	fmt.Fprintf(w, "<form name=\"primary\">: Dropdown Menu value = \"%v\"\n", vMenu)
	fmt.Fprintf(w, "<form name=\"primary\">: Image Field value = \"%v\"\n", header.Filename)
	fmt.Fprintf(w, "\tUploaded: \"%v\"\n", uploadpath)
	fmt.Fprintf(w, "\tFormatted: \"%v\"\n", fvTextField)

	// Construct paths to output files
	var htmlpath string = filepath.Join(OUTPUTDIRECTORY, "primary-text.html")
	// var docpath string = filepath.Join(OUTPUTDIRECTORY, "primary-text.doc")

	// Construct DOC output

	// Construct HTML output
	html.Paragraph(&outHTML, fvTextField)
	html.Image(&outHTML, uploadpath)

	// Write output to corresponding files
	// outputdata.WriteOutput(docpath, "outputdata/templates/template.doc", PERMISSIONBITS, outDOC)
	outputdata.WriteOutput(htmlpath, "outputdata/templates/template.html", PERMISSIONBITS, outHTML)
}