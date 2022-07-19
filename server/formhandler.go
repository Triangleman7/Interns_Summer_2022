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

type FormOutput struct {
	TemplateDOCX string		// Filename of template DOCX file
	TemplateHTML string		// Filename of template HTML file

	OutDOCX string		// Filename of output DOCX file
	OutHTML string		// Filename of output HTML file
}

type FormPrimary struct {
	Output FormOutput		// Output specifications

	PrimaryImage string				// {primary-image}
	PrimaryText string				// {primary-text}
}

// HandleFormPrimary handles form submission to form#primary.
//
// Raises any errors encountered while handling the form or procesing form input.
func HandleFormPrimary(w http.ResponseWriter, r *http.Request) (err error) {
	var form FormPrimary
	form.Output = FormOutput{"template.docx", "template.html", "form-primary.docx", "form-primary.html"}
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
	var templatepath = filepath.Join(TEMPLATEDIRECTORY, form.Output.TemplateDOCX)
	var outpath = filepath.Join(OUTPUTDIRECTORY, form.Output.OutDOCX)

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
	var outpath = filepath.Join(OUTPUTDIRECTORY, form.Output.OutHTML)
	var templatepath = filepath.Join(TEMPLATEDIRECTORY, form.Output.TemplateHTML)

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
