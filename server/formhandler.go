package server

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/Triangleman7/Interns_Summer_2022/msword"
	"github.com/Triangleman7/Interns_Summer_2022/outputdata/docx"
	"github.com/Triangleman7/Interns_Summer_2022/outputdata/html"
)

func HandleFormPrimary(w http.ResponseWriter, r *http.Request) (err error) {
	// Parse form submission
	err = r.ParseMultipartForm(0)
	if err != nil {
		return
	}

	// Process <input name="primary-text"> form field
	var vTextField, vMenu, fvTextField string
	vTextField = r.FormValue("primary-text")
	vMenu = r.FormValue("primary-text-operation")
	fvTextField, err = FormatValue(vTextField, vMenu)
	if err != nil {
		return
	}

	// Process <input name="primary-image"> form field
	var file multipart.File
	var header *multipart.FileHeader
	var uploadpath string
	file, header, err = r.FormFile("primary-image")
	if err != nil {
		return
	}
	defer file.Close()
	uploadpath, err = UploadFile(file, header)
	if err != nil {
		return
	}

	// Debugging
	fmt.Fprintf(w, "<form name=\"primary\">: Text Field value = '%v', Dropdown Menu value = '%v', Image Field value = '%v'\n", vTextField, vMenu, header.Filename)
	fmt.Fprintf(w, "\tUploaded: \"%v\"\n", uploadpath)
	fmt.Fprintf(w, "\tFormatted: \"%v\"\n", fvTextField)

	// Construct paths to output files
	var htmlpath string = filepath.Join(OUTPUTDIRECTORY, "form-primary.html")
	var docxpath string = filepath.Join(OUTPUTDIRECTORY, "form-primary.docx")

	// Construct DOCX output
	var reader *msword.ReplaceDocx
	reader, err = msword.ReadDocxFile("outputdata/templates/template.docx")
	if err != nil {
		return
	}
	var outDOCX *msword.Docx = reader.Editable()
	docx.Paragraph("primary-text", outDOCX, fvTextField)
	docx.Image(1, outDOCX, uploadpath)

	// Construct HTML output
	var outHTML string
	outHTML, err = html.ReadTemplate("outputdata/templates/template.html")
	if err != nil {
		return
	}
	html.Paragraph("primary-text", &outHTML, fvTextField)
	html.Image("primary-image", &outHTML, uploadpath)

	// Write output to corresponding files
	docx.WriteDOCX(docxpath, outDOCX)
	html.WriteHTML(htmlpath, PERMISSIONBITS, outHTML)

	return
}
