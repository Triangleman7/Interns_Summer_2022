package server

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/Triangleman7/Interns_Summer_2022/msword"
	"github.com/Triangleman7/Interns_Summer_2022/server/docx"
	"github.com/Triangleman7/Interns_Summer_2022/server/html"
)

type FormOutput interface {
	handle(http.ResponseWriter, *http.Request) error // Handles submissions to the form element

	outputDOCX() error // Outputs to a Word Document (.docx)
	outputHTML() error // Outputs to an HTML Document (.html)
}

type Form struct {
	Name string // Field element 'name' attribute
}

// TemplateDOCX returns the path (relative to the root directory) to the template DOCX file for the
// form f.
func (f Form) TemplateDOCX() (path string) {
	var filename string = fmt.Sprintf("%s.docx", f.Name)
	return filepath.Join(TEMPLATEDIRECTORY, filename)
}

// TemplateHTML returns the path (relative to the root directory) to the template HTML file for the
// form f.
func (f Form) TemplateHTML() (path string) {
	var filename string = fmt.Sprintf("%s.html", f.Name)
	return filepath.Join(TEMPLATEDIRECTORY, filename)
}

// OutDOCX returns the path (relative to the root directory) to the output DOCX file for the form
// f.
func (f Form) OutDOCX() (path string) {
	var filename string = fmt.Sprintf("%s.docx", f.Name)
	return filepath.Join(OUTPUTDIRECTORY, filename)
}

// OutHTML returns the path (relative to the root directory) to the output HTML file for the form
// f.
func (f Form) OutHTML() (path string) {
	var filename string = fmt.Sprintf("%s.html", f.Name)
	return filepath.Join(OUTPUTDIRECTORY, filename)
}

type FormPrimary struct {
	form Form

	primaryImage string // {primary-image}
	primaryText  string // {primary-text}
}

func (f FormPrimary) handle(w http.ResponseWriter, r *http.Request) (err error) {
	f.form = Form{"form-primary"}
	log.Print("Handling form submission to form#primary")

	// Parse form submission
	err = r.ParseMultipartForm(0)
	if err != nil {
		return
	}
	log.Print("Parsed form submission")

	// Process element input[name="primary-text"]
	var primaryText string = r.FormValue("primary-text")
	var primaryTextOperation string = r.FormValue("primary-text-operation")
	f.primaryText, err = FormatValue(primaryText, primaryTextOperation)
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
	f.primaryImage, err = UploadFile(file, header)
	if err != nil {
		return
	}
	log.Print("Processed <input name=\"primary-image\"> field")

	// Write output
	err = f.outputDOCX()
	if err != nil {
		return
	}
	err = f.outputHTML()
	if err != nil {
		return
	}
	log.Printf("Successfully wrote all output: %s", OUTPUTDIRECTORY)

	return
}

func (f FormPrimary) outputDOCX() (err error) {
	var templatepath = f.form.TemplateDOCX()
	var outpath = f.form.OutDOCX()

	var reader *msword.ReplaceDocx
	var outDOCX *msword.Docx
	reader, outDOCX, err = docx.ReadTemplate(templatepath)
	if err != nil {
		return
	}
	defer reader.Close()

	docx.Image(outDOCX, 1, f.primaryImage)
	docx.Paragraph(outDOCX, "primary-text", f.primaryText)

	err = docx.WriteDOCX(outpath, outDOCX)
	if err != nil {
		return
	}

	log.Printf("DOCX output written to %s", outpath)

	return
}

func (f FormPrimary) outputHTML() (err error) {
	var templatepath = f.form.TemplateHTML()
	var outpath = f.form.OutHTML()

	var outHTML string
	outHTML, err = html.ReadTemplate(templatepath)
	if err != nil {
		return
	}

	html.Image(&outHTML, "primary-image", f.primaryImage)
	html.Paragraph(&outHTML, "primary-text", f.primaryText)

	err = html.WriteHTML(outpath, outHTML)
	if err != nil {
		return
	}

	log.Printf("HTML output written to %s", outpath)

	return
}
