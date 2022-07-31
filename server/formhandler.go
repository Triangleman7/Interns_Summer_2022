package server

import (
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"time"

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

	imageUpload    string // {primary-image}
	imageTimestamp string // {primary-image-timestamp}
	captionText    string // {primary-text}
}

func (f *FormPrimary) handle(w http.ResponseWriter, r *http.Request) (err error) {
	f.form.Name = "form-primary"
	log.Printf("%s - Handling form submission", f.form.Name)

	// Parse form submission
	err = r.ParseMultipartForm(0)
	if err != nil {
		return
	}
	log.Printf("%s - Parsed form submission", f.form.Name)

	// Process form <input> fields
	var captionText string = r.FormValue("caption-text")
	var captionCasing string = r.FormValue("caption-casing")
	var imageUploadFile multipart.File
	var imageUploadHeader *multipart.FileHeader
	imageUploadFile, imageUploadHeader, err = r.FormFile("image-upload")
	if err != nil {
		var css string = "input[name='image-upload']"
		log.Printf("%s - Empty form <input> field: %s", f.form.Name, css)
		return
	}
	var imageTimestamp string = r.FormValue("image-timestamp")
	log.Printf("%s - Processed form <input> fields", f.form.Name)

	// Process {primary-image} output field
	defer imageUploadFile.Close()
	var uploadpath string
	uploadpath, err = UploadFile(imageUploadFile, imageUploadHeader)
	if err != nil {
		return
	}
	f.imageUpload = filepath.Join(f.form.OutImages(), imageUploadHeader.Filename)
	err = CopyFile(uploadpath, f.imageUpload)
	if err != nil {
		return
	}

	// Process {primary-image-timestamp} output field
	var timeFormat = "2006-01-02T03:04"
	_, err = time.Parse(timeFormat, imageTimestamp)
	if err != nil {
		f.imageTimestamp = imageTimestamp
	} else {
		f.imageTimestamp = time.Now().Format(timeFormat)
	}

	// Process {primary-text} output field
	f.captionText = FormatValue(captionText, captionCasing)

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
	log.Printf("%s - Successfully wrote all output: %s", f.form.Name, OUTPUTDIRECTORY)

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

	docx.Image(outDOCX, 1, f.imageUpload)
	docx.Paragraph(outDOCX, "caption-text", f.captionText)

	err = docx.WriteDOCX(outpath, outDOCX)
	if err != nil {
		return
	}

	log.Printf("%s - DOCX output written to %s", f.form.Name, outpath)

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

	html.Image(&outHTML, "image-upload", f.imageUpload)
	html.Paragraph(&outHTML, "image-timestamp", f.imageTimestamp)
	html.Paragraph(&outHTML, "caption-text", f.captionText)

	err = html.WriteHTML(outpath, outHTML)
	if err != nil {
		return
	}

	log.Printf("%s - HTML output written to %s", f.form.Name, outpath)

	return
}

func (f *FormPrimary) copyCSS() (err error) {
	return CopyFile(f.form.TemplateCSS(), f.form.OutCSS())
}
