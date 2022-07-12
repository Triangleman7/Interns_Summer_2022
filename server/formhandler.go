package server

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/Triangleman7/Interns_Summer_2022/outputdata/doc"
	"github.com/Triangleman7/Interns_Summer_2022/outputdata/html"
)

func HandleFormPrimary(w http.ResponseWriter, r *http.Request) {
	var err error

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
	uploadpath = UploadFile(file, header)
	err = file.Close()
	if err != nil { panic(err) }

	// Debugging\
	fmt.Fprintf(w, "<form name=\"primary\">: Text Field value = '%v', Dropdown Menu value = '%v', Image Field value = '%v'\n", vTextField, vMenu, header.Filename)
	fmt.Fprintf(w, "\tUploaded: \"%v\"\n", uploadpath)
	fmt.Fprintf(w, "\tFormatted: \"%v\"\n", fvTextField)

	// Construct paths to output files
	var htmlpath string = filepath.Join(OUTPUTDIRECTORY, "form-primary.html")
	var docpath string = filepath.Join(OUTPUTDIRECTORY, "form-primary.doc")

	// Construct DOC output
	var outDOC []byte
	doc.Paragraph(&outDOC, fvTextField)
	doc.Image(&outDOC, uploadpath)

	// Construct HTML output
	var outHTML string
	err, outHTML = html.ReadTemplate("outputdata/templates/template.html")
	if err != nil { panic(err) }
	html.Paragraph("primary-text", &outHTML, fvTextField)
	html.Image("primary-image", &outHTML, uploadpath)

	// Write output to corresponding files
	doc.WriteDOC(docpath, "outputdata/templates/template.doc", PERMISSIONBITS, outDOC)
	html.WriteHTML(htmlpath, PERMISSIONBITS, outHTML)
}