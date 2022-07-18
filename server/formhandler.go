package server

import (
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/Triangleman7/Interns_Summer_2022/msword"
	"github.com/Triangleman7/Interns_Summer_2022/outputdata/docx"
	"github.com/Triangleman7/Interns_Summer_2022/outputdata/html"
)

// HandleFormPrimary handles form submission to form#primary.
//
// Raises any errors encountered while handling the form or procesing form input.
func HandleFormPrimary(w http.ResponseWriter, r *http.Request) (err error) {
	log.Print("Handling form submission to form#primary")

	// Parse form submission
	err = r.ParseMultipartForm(0)
	if err != nil {
		return
	}
	log.Print("Parsed form submission")

	// Process element input[name="primary-text"]
	var vTextField, vMenu, fvTextField string
	vTextField = r.FormValue("primary-text")
	vMenu = r.FormValue("primary-text-operation")
	fvTextField, err = FormatValue(vTextField, vMenu)
	if err != nil {
		return
	}
	log.Print("Processed <input name=\"primary-text\"> field")

	// Process element input[name="primary-image"]
	var file multipart.File
	var header *multipart.FileHeader
	file, header, err = r.FormFile("primary-image")
	if err != nil {
		return
	}
	defer file.Close()
	var uploadpath string
	uploadpath, err = UploadFile(file, header)
	if err != nil {
		return
	}
	log.Print("Processed <input name=\"primary-image\"> field")

	// Output DOCX
	var docxpath string = filepath.Join(OUTPUTDIRECTORY, "form-primary.docx")
	var reader *msword.ReplaceDocx
	reader, err = msword.ReadDocxFile("outputdata/templates/template.docx")
	if err != nil {
		return
	}
	defer reader.Close()
	var outDOCX *msword.Docx = reader.Editable()
	docx.Paragraph("primary-text", outDOCX, fvTextField)
	docx.Image(1, outDOCX, uploadpath)
	err = docx.WriteDOCX(docxpath, outDOCX)
	if err != nil {
		return
	}
	log.Printf("DOCX output written to %s", docxpath)

	// Output HTML
	var htmlpath string = filepath.Join(OUTPUTDIRECTORY, "form-primary.html")
	var outHTML string
	outHTML, err = html.ReadTemplate("outputdata/templates/template.html")
	if err != nil {
		return
	}
	html.Paragraph("primary-text", &outHTML, fvTextField)
	html.Image("primary-image", &outHTML, uploadpath)
	err = html.WriteHTML(htmlpath, outHTML)
	if err != nil {
		return
	}
	log.Printf("HTML output written to %s", htmlpath)

	log.Print("Successfully wrote all output to out/")
	return
}
