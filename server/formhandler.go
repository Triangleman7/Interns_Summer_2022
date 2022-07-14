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

func HandleFormPrimary(w http.ResponseWriter, r *http.Request) (err error) {
	log.Print("Handling form submission to form#primary")

	// Parse form submission
	err = r.ParseMultipartForm(0)
	if err != nil {
		return
	}
	log.Print("Parsed form submission")

	// Process <input name="primary-text"> form field
	var vTextField, vMenu, fvTextField string
	vTextField = r.FormValue("primary-text")
	vMenu = r.FormValue("primary-text-operation")
	fvTextField, err = FormatValue(vTextField, vMenu)
	if err != nil {
		return
	}
	log.Print("Processed <input name=\"primary-text\"> field")

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
	log.Print("Processed <input name=\"primary-image\"> field")

	// Output DOCX
	var docxpath string = filepath.Join(OUTPUTDIRECTORY, "form-primary.docx")
	var reader *msword.ReplaceDocx
	reader, err = msword.ReadDocxFile("outputdata/template/template.docx")
	if err != nil {
		return
	}
	defer reader.Close()
	var outDOCX *msword.Docx = reader.Editable()
	docx.Paragraph("primary-text", outDOCX, fvTextField)
	docx.Image(1, outDOCX, uploadpath)
	docx.WriteDOCX(docxpath, outDOCX)
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
	html.WriteHTML(htmlpath, PERMISSIONBITS, outHTML)
	log.Printf("HTML output written to %s", htmlpath)

	return
}
