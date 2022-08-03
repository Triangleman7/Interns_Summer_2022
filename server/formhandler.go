package server

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Triangleman7/Interns_Summer_2022/resources/msword"
	"github.com/Triangleman7/Interns_Summer_2022/server/docx"
	"github.com/Triangleman7/Interns_Summer_2022/server/html"
	"github.com/Triangleman7/Interns_Summer_2022/server/scss"
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

func (f *Form) OutZIP() (path string) {
	return filepath.Join(OUTPUTDIRECTORY, fmt.Sprintf("%s.zip", f.Name))
}

// SetupOutput creates the output directory for the form f and all necessary subdirectories.
func (f *Form) SetupOutput(name string) {
	f.Name = name

	DirectorySetup(f.OutDir(), FILEMODE)

	var subdirs = []string{"images"}
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

// TemplateSCSS returns the path to the template CSS stylesheet.
func (f *Form) TemplateSCSS() (path string) {
	return f.GetTemplate("styles.scss")
}

// OutSCSS returns the path to the output CSS stylesheet.
func (f *Form) OutSCSS() (path string) {
	return f.GetOut("styles.scss")
}

// OutImages returns the path to the output images/ directory.
func (f *Form) OutImages() (path string) {
	return f.GetOut("images/")
}

type FormPrimary struct {
	form Form

	fileUpload       string
	fileUploadFile   multipart.File
	fileUploadHeader *multipart.FileHeader
	uploadTimestamp  string
	imageScale       int
	imageAlign       string
	captionText      string
	captionAlign     string
	captionCasing    string
	captionStyling   map[string]bool
}

func (f *FormPrimary) handle(w http.ResponseWriter, r *http.Request) (zippath string, err error) {
	f.form.Name = "form-primary"
	log.Printf("%s - Handling form submission", f.form.Name)

	// Parse form submission
	err = r.ParseMultipartForm(0)
	if err != nil {
		return
	}
	log.Printf("%s - Parsed form submission", f.form.Name)

	// Process form <input> fields
	f.fileUploadFile, f.fileUploadHeader, err = r.FormFile("file-upload")
	if err != nil {
		return
	}
	f.uploadTimestamp = r.FormValue("upload-timestamp")
	f.imageScale, err = strconv.Atoi(r.FormValue("image-scale"))
	if err != nil {
		return
	}
	f.imageAlign = r.FormValue("image-align")
	f.captionText = r.FormValue("caption-text")
	f.captionAlign = r.FormValue("caption-align")
	f.captionCasing = r.FormValue("caption-casing")
	f.captionStyling = make(map[string]bool)
	for _, style := range []string{
		"italic", "bold", "underline", "strikethrough",
	} {
		if r.FormValue(fmt.Sprintf("caption-styling-%s", style)) == "on" {
			f.captionStyling[style] = true
		} else {
			f.captionStyling[style] = false
		}
	}
	log.Printf("%s - Processed form <input> fields: %v", f.form.Name, f.form)

	// Process {primary-image} output field
	defer func() {
		var e = f.fileUploadFile.Close()
		if e != nil {
			panic(e)
		}
	}()
	var uploadpath string
	uploadpath, err = UploadFile(f.fileUploadFile, f.fileUploadHeader)
	if err != nil {
		return
	}
	f.fileUpload = filepath.Join(f.form.OutImages(), f.fileUploadHeader.Filename)
	err = CopyFile(uploadpath, f.fileUpload)
	if err != nil {
		return
	}
	log.Printf("%s - Processed output for field {primary-image}", f.form.Name)

	// Process {primary-image-timestamp} output field
	var timeFormat = "2006-01-02T03:04"
	var datetime time.Time
	datetime, err = time.Parse(timeFormat, f.uploadTimestamp)
	if r.FormValue("image-timestamp") == "" {
		f.uploadTimestamp = time.Now().Format(timeFormat)
	} else if err != nil {
		f.uploadTimestamp = time.Now().Format(timeFormat)
	} else {
		f.uploadTimestamp = datetime.Format(timeFormat)
	}
	log.Printf("%s - Processed output for field {primary-images-timestamp}", f.form.Name)

	// Process {primary-text} output field
	f.captionText = FormatValue(f.captionText, f.captionCasing)
	log.Printf("%s - Processed output for field {primary-text}", f.form.Name)

	// Write output
	err = f.outputDOCX()
	if err != nil {
		return
	}
	err = f.outputHTML()
	if err != nil {
		return
	}
	err = f.outputSCSS()
	if err != nil {
		return
	}
	log.Printf("%s - Successfully wrote all output: %s", f.form.Name, OUTPUTDIRECTORY)

	// Transpile output SCSS to CSS
	var cmd = exec.Command(
		"sass", fmt.Sprintf("%s:%s", f.form.OutDir(), f.form.OutDir()),
	)
	err = cmd.Run()
	if err != nil {
		return
	}
	log.Printf("%s - Successfully compiled SCSS to CSS: %s", f.form.Name, f.form.OutDir())

	err = f.outputZIP()
	if err != nil {
		panic(err)
		return
	}
	log.Printf("%s - Successfully compressed output into ZIP archive: %s", f.form.Name, f.form.OutZIP())

	return f.form.OutZIP(), nil
}

func (f *FormPrimary) outputZIP() (err error) {
	var outpath = f.form.OutZIP()

	var file *os.File
	file, err = os.Create(outpath)
	if err != nil {
		return
	}
	defer func() {
		var e = file.Close()
		if e != nil {
			panic(err)
		}
	}()
	log.Printf("%s - Destination file created: %s", f.form.Name, outpath)

	var writer = zip.NewWriter(file)
	defer func() {
		var e = writer.Close()
		if e != nil {
			panic(err)
		}
	}()
	log.Printf("%s - Created ZIP archive writer: %s", f.form.Name, outpath)

	walker := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		
		var source *os.File
		source, err = os.Open(path)
		if err != nil {
			return err
		}
		defer func() {
			var e = source.Close()
			if e != nil {
				panic(e)
			}
		}()
		log.Printf("%s - Opened origin file: %s", f.form.Name, source.Name())
	
		var destination io.Writer
		var relpath string
		relpath, err = filepath.Rel(f.form.OutDir(), path)
		if err != nil {
			return err
		}
		destination, err = writer.Create(relpath)
		if err != nil {
			return err
		}
		log.Printf("%s - Destination file created in ZIP archive: %s", f.form.Name, source.Name())

		_, err = io.Copy(destination, source)
		if err != nil {
			return err
		}
		log.Printf("%s - File contents copied to ZIP archive: %s", f.form.Name, source.Name())

		return nil
	}

	err = filepath.Walk(f.form.OutDir(), walker)
	if err != nil {
		panic(err)
	}
	log.Printf("%s - Walked directory: %s", f.form.Name, f.form.OutDir())

	return nil
}

func (f *FormPrimary) outputDOCX() (err error) {
	var templatepath = f.form.TemplateDOCX()
	var outpath = f.form.OutDOCX()

	var reader *msword.ReplaceDocx
	var outDOCX *msword.Docx
	reader, outDOCX, err = docx.ReadTemplate(templatepath)
	if err != nil {
		return
	}
	defer func() {
		var e = reader.Close()
		if e != nil {
			panic(e)
		}
	}()

	docx.Image(outDOCX, 1, f.fileUpload)
	docx.Paragraph(outDOCX, "image-timestamp", f.uploadTimestamp)
	docx.Paragraph(outDOCX, "caption-text", f.captionText)

	err = docx.WriteDOCX(outpath, outDOCX)
	if err != nil {
		return
	}

	log.Printf("%s - DOCX output written to %s", f.form.Name, outpath)

	return
}

func (f *FormPrimary) outputHTML() (err error) {
	var templatepath = f.form.TemplateHTML()
	var outpath = f.form.OutHTML()

	var outHTML string
	outHTML, err = html.ReadTemplate(templatepath)
	if err != nil {
		return
	}

	html.Image(&outHTML, "image-upload", f.fileUpload)
	html.Paragraph(&outHTML, "image-timestamp", f.uploadTimestamp)
	html.Paragraph(&outHTML, "caption-text", f.captionText)

	err = html.WriteHTML(outpath, outHTML)
	if err != nil {
		return
	}

	log.Printf("%s - HTML output written to %s", f.form.Name, outpath)

	return
}

func (f *FormPrimary) outputSCSS() (err error) {
	var templatepath = f.form.TemplateSCSS()
	var outpath = f.form.OutSCSS()

	var outSCSS string
	outSCSS, err = scss.ReadTemplate(templatepath)
	if err != nil {
		return
	}

	scss.Rule(
		&outSCSS,
		"div.container",
		map[string]string{
			"display": "block",
			"width":   fmt.Sprintf("%d%%", f.imageScale),
			"margin":  scss.ImgMargin(f.imageAlign),
		},
	)
	scss.Rule(
		&outSCSS,
		"img.image-upload",
		map[string]string{
			"height": "100%",
			"width":  "100%",
		},
	)
	scss.Rule(
		&outSCSS,
		"p.image-timestamp",
		map[string]string{
			"text-align":  "center",
			"font-family": "monospace",
		},
	)
	scss.Rule(
		&outSCSS,
		"p.caption-text",
		map[string]string{
			"text-align":  f.captionAlign,
			"font-style":  scss.PFontStyle(f.captionStyling["italic"]),
			"font-weight": scss.PFontWeight(f.captionStyling["bold"]),
			"text-decoration": scss.PTextDecoration(
				f.captionStyling["strikethrough"], f.captionStyling["underline"],
			),
			"overflow-wrap": "break-word",
		},
	)

	err = scss.WriteSCSS(outpath, outSCSS)
	if err != nil {
		return
	}

	log.Printf("%s - SCSS output written to %s", f.form.Name, outpath)

	return
}
