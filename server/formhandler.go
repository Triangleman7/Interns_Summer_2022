package server

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/Triangleman7/Interns_Summer_2022/msword"
	"github.com/Triangleman7/Interns_Summer_2022/outputdata/docx"
	"github.com/Triangleman7/Interns_Summer_2022/outputdata/html"
)

type FormOutput struct {
	Name string // Field element 'name' attribute
}

// TemplateDOCX returns the path (relative to the root directory) to the template DOCX file for the
// form f.
func (f FormOutput) TemplateDOCX() (path string) {
	var filename string = fmt.Sprintf("%s.docx", f.Name)
	return filepath.Join(TEMPLATEDIRECTORY, filename)
}

// TemplateHTML returns the path (relative to the root directory) to the template HTML file for the
// form f.
func (f FormOutput) TemplateHTML() (path string) {
	var filename string = fmt.Sprintf("%s.html", f.Name)
	return filepath.Join(TEMPLATEDIRECTORY, filename)
}

// OutDOCX returns the path (relative to the root directory) to the output DOCX file for the form
// f.
func (f FormOutput) OutDOCX() (path string) {
	var filename string = fmt.Sprintf("%s.docx", f.Name)
	return filepath.Join(OUTPUTDIRECTORY, filename)
}

// OutHTML returns the path (relative to the root directory) to the output HTML file for the form
// f.
func (f FormOutput) OutHTML() (path string) {
	var filename string = fmt.Sprintf("%s.html", f.Name)
	return filepath.Join(OUTPUTDIRECTORY, filename)
}

type FormPrimary struct {
	Output FormOutput // Output specifications

	PrimaryImage string // {primary-image}
	PrimaryText  string // {primary-text}
}

// HandleFormPrimary handles form submission to form#primary.
//
// Raises any errors encountered while handling the form or procesing form input.
func HandleFormPrimary(w http.ResponseWriter, r *http.Request) (err error) {
	var form FormPrimary
	form.Output = FormOutput{"form-primary"}
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
	form.PrimaryText, err = FormatValue(primaryText, primaryTextOperation)
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
	form.PrimaryImage, err = UploadFile(file, header)
	if err != nil {
		return
	}
	log.Print("Processed <input name=\"primary-image\"> field")

	// Write output
	err = FormPrimaryDOCX(form)
	if err != nil {
		return
	}
	err = FormPrimaryHTML(form)
	if err != nil {
		return
	}
	log.Printf("Successfully wrote all output: %s", OUTPUTDIRECTORY)

	return
}

func FormPrimaryDOCX(form FormPrimary) (err error) {
	var templatepath = form.Output.TemplateDOCX()
	var outpath = form.Output.OutDOCX()

	var reader *msword.ReplaceDocx
	reader, err = msword.ReadDocxFile(templatepath)
	if err != nil {
		return
	}
	defer reader.Close()

	var outDOCX *msword.Docx = reader.Editable()

	docx.Image(outDOCX, 1, form.PrimaryImage)
	docx.Paragraph(outDOCX, "primary-text", form.PrimaryText)

	err = docx.WriteDOCX(outpath, outDOCX)
	if err != nil {
		return
	}

	log.Printf("DOCX output written to %s", outpath)

	return
}

func FormPrimaryHTML(form FormPrimary) (err error) {
	var templatepath = form.Output.TemplateHTML()
	var outpath = form.Output.OutHTML()

	var outHTML string
	outHTML, err = html.ReadTemplate(templatepath)
	if err != nil {
		return
	}

	html.Image(&outHTML, "primary-image", form.PrimaryImage)
	html.Paragraph(&outHTML, "primary-text", form.PrimaryText)

	err = html.WriteHTML(outpath, outHTML)
	if err != nil {
		return
	}

	log.Printf("HTML output written to %s", outpath)

	return
}
