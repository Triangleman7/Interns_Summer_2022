package server

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/Triangleman7/Interns_Summer_2022/resources/msword"
	"github.com/Triangleman7/Interns_Summer_2022/server/docx"
	"github.com/Triangleman7/Interns_Summer_2022/server/html"
)

type FormOutput interface {
	handle(http.ResponseWriter, *http.Request) error // Handles submissions to the form element

	outputDOCX() error // Outputs to a Word Document (.docx)
	outputHTML() error // Outputs to an HTML Document (.html)

	copyCSS() error // Copies CSS stylesheet to output directory
}

type Form struct {
	Name string // Field element 'name' attribute
}

// OutDir constructs the path (relative to the root directory) to the output directory for the form
// f.
func (f *Form) OutDir() (path string) {
	return filepath.Join(OUTPUTDIRECTORY, f.Name)
}

// SetupOutput creates the output directory for the form f and all necessary subdirectories.
func (f *Form) SetupOutput(name string) {
	f.Name = name

	DirectorySetup(f.OutDir(), FILEMODE)

	var subdirs []string = []string{"images"}
	for _, sd := range subdirs {
		DirectorySetup(filepath.Join(f.OutDir(), sd), FILEMODE)
	}
}

// GetTemplate returns the path to the template file under the name filename.
func (f *Form) GetTemplate(filename string) (path string) {
	return filepath.Join(TEMPLATEDIRECTORY, f.Name, filename)
}

// GetOut returns the path to the output file under the name filename.
func (f *Form) GetOut(filename string) (path string) {
	return filepath.Join(OUTPUTDIRECTORY, f.Name, filename)
}

// TemplateDOCX returns the path to the template Word Document (DOCX).
func (f *Form) TemplateDOCX() (path string) {
	return f.GetTemplate("index.docx")
}

// OutDOCX returns the path to the output Word Document (DOCX).
func (f *Form) OutDOCX() (path string) {
	return f.GetOut("index.docx")
}

// TemplateHTML returns the path to the template HTML Document.
func (f *Form) TemplateHTML() (path string) {
	return f.GetTemplate("index.html")
}

// OutHTML returns the path to the output HTML Document.
func (f *Form) OutHTML() (path string) {
	return f.GetOut("index.html")
}

// TemplateCSS returns the path to the template CSS stylesheet.
func (f *Form) TemplateCSS() (path string) {
	return f.GetTemplate("styles.css")
}

// OutCSS returns the path to the output CSS stylesheet.
func (f *Form) OutCSS() (path string) {
	return f.GetOut("styles.css")
}

// OutImages returns the path to the output images/ directory.
func (f *Form) OutImages() (path string) {
	return f.GetOut("images/")
}

type FormPrimary struct {
	form Form

	primaryImage string // {primary-image}
	primaryText  string // {primary-text}
}

func (f *FormPrimary) handle(w http.ResponseWriter, r *http.Request) (err error) {
	if f.form.Name == "" {
		return fmt.Errorf("output directory for element 'form#primary' not set up")
	}
	log.Print("Handling form submission to form#primary")

	// Parse form submission
	err = r.ParseMultipartForm(0)
	if err != nil {
		return
	}
	log.Print("Parsed form submission")

	// Process form <input> fields
	var captionText string = r.FormValue("caption-text")
	var captionCasing string = r.FormValue("caption-casing")
	var imageUploadFile multipart.File
	var imageUploadHeader *multipart.FileHeader
	imageUploadFile, imageUploadHeader, err = r.FormFile("image-upload")
	if err != nil {
		var css string = "form#primary input[name='image-upload']"
		log.Printf("Empty form <input> field: %s", css)
		return
	}
	log.Printf("Processed form <input> fields")

	// Process {primary-image} output field
	defer imageUploadFile.Close()
	var uploadpath string
	uploadpath, err = UploadFile(imageUploadFile, imageUploadHeader)
	if err != nil {
		return
	}
	f.primaryImage = filepath.Join(f.form.OutImages(), imageUploadHeader.Filename)
	err = CopyFile(uploadpath, f.primaryImage)
	if err != nil {
		return
	}

	// Process {primary-text} output field
	f.primaryText = FormatValue(captionText, captionCasing)

	// Write output
	err = f.outputDOCX()
	if err != nil {
		return
	}
	err = f.outputHTML()
	if err != nil {
		return
	}
	err = f.copyCSS()
	if err != nil {
		return
	}
	log.Printf("Successfully wrote all output: %s", OUTPUTDIRECTORY)

	return
}

func (f *FormPrimary) outputDOCX() (err error) {
	var templatepath string = f.form.TemplateDOCX()
	var outpath string = f.form.OutDOCX()

	var reader *msword.ReplaceDocx
	var outDOCX *msword.Docx
	reader, outDOCX, err = docx.ReadTemplate(templatepath)
	if err != nil {
		return
	}
	defer reader.Close()

	docx.Image(outDOCX, 1, f.primaryImage)
	docx.Paragraph(outDOCX, "caption-text", f.primaryText)

	err = docx.WriteDOCX(outpath, outDOCX)
	if err != nil {
		return
	}

	log.Printf("DOCX output written to %s", outpath)

	return
}

func (f *FormPrimary) outputHTML() (err error) {
	var templatepath string = f.form.TemplateHTML()
	var outpath string = f.form.OutHTML()

	var outHTML string
	outHTML, err = html.ReadTemplate(templatepath)
	if err != nil {
		return
	}

	html.Image(&outHTML, "image-upload", f.primaryImage)
	html.Paragraph(&outHTML, "caption-text", f.primaryText)

	err = html.WriteHTML(outpath, outHTML)
	if err != nil {
		return
	}

	log.Printf("HTML output written to %s", outpath)

	return
}

func (f *FormPrimary) copyCSS() (err error) {
	return CopyFile(f.form.TemplateCSS(), f.form.OutCSS())
}
