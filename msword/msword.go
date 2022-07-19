/* Package msword handles Microsoft Word Documents.

The code in this package has been adapted from the github.com/nguyenthenguyen/docx package:
- GitHub: https://github.com/nguyenthenguyen/docx
- Go Package Index: https://pkg.go.dev/github.com/nguyenthenguyen/docx
*/
package msword

import (
	"archive/zip"
	"bytes"
	"io"
	"log"
)

type ZipData interface {
	// files serves the content from a ZIP archive.
	files() []*zip.File

	// close closes the ZIP file.
	close() error
}

type ZipFile struct {
	// Data is a reader for the ZIP archive.
	Data *zip.ReadCloser
}

func (d ZipFile) files() []*zip.File {
	return d.Data.File
}

func (d ZipFile) close() error {
	return d.Data.Close()
}

// ReplaceDocx represents a Word Document, the document ZIP archive, and the document contents.
type ReplaceDocx struct {
	// ZipReader is a reader of the Word Document ZIP archive.
	ZipReader ZipData

	// Content is the text content of the body of the Word Document.
	Content string

	// Links is the text content of the hyperlinks of the Word Document.
	Links string

	// Headers is a map of header style names to the corresponding text content of the Word
	// Document.
	Headers map[string]string

	// Footers is a map of footer style names to the corresponding text content of the Word
	// Document.
	Footers map[string]string

	// Images is a map of the filenames of the images of the Word Document to the filename of its
	// replacement file. The file extension of the two files must match.
	Images map[string]string
}

// Docx represents a Word Document, the document ZIP archive files, and the document contents.
type Docx struct {
	// Files is an array of individual files contained in the ZIP archive.
	Files []*zip.File

	// Content is the text content of the body of the Word Document.
	Content string

	// links is the text content of the hyperlinks of the Word Document.
	Links string

	// Headers is a map of header style names to the corresponding text content of the Word
	// Document.
	Headers map[string]string

	// Footers is a map of footer style names to the corresponding text content of the Word
	// Document.
	Footers map[string]string

	// Images is a map of the filenames of the images of the Word Document to the filename of its
	// replacement file. The file extension of the two files must match.
	Images map[string]string
}

// Close closes the ZIP archive of the Word Document d.
func (r *ReplaceDocx) Close() error {
	return r.ZipReader.close()
}

// GetContent returns the content of the Word Document (d) body.
func (d *Docx) GetContent() string {
	return d.Content
}

// SetContent sets the content of the Word Document (d) body to content.
func (d *Docx) SetContent(content string) {
	d.Content = content
}

// ReadDocxFile reads the Word Document specified by path. The Word Document is read by opening the
// ZIP archive located at path. A Word Document object (type ReplaceDocx) is returned.
//
// Raises any errors encountered while opening the ZIP archive.
func ReadDocxFile(path string) (*ReplaceDocx, error) {
	log.Printf("Reading Word Document: %s", path)

	reader, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	log.Printf("Opened Word Document ZIP archive: %s", path)

	zipData := ZipFile{Data: reader}
	return ReadDocx(zipData)
}

// ReadDocx reads the contents of the Word Document ZIP archive (ZipData) and processes each aspect
// of the Word Document (body, hyperlinks, headers, footers, images). A Word Document object (type
// Replace Docx) is returned.
//
// Raises any errors encountered while reading the contents of the Word Document ZIP archive.
func ReadDocx(reader ZipData) (*ReplaceDocx, error) {
	log.Print("Reading Word Document ZIP archive")

	content, err := readText(reader.files())
	if err != nil {
		return nil, err
	}
	log.Print("Read Word Document body text")

	links, err := readLinks(reader.files())
	if err != nil {
		return nil, err
	}
	log.Printf("Read Word Document hyperlinks")

	headers, footers, _ := readHeaderFooter(reader.files())
	log.Printf("Read Word Document headers/footers")
	images, _ := retrieveImageFilenames(reader.files())
	log.Printf("Read Word Document images")
	return &ReplaceDocx{ZipReader: reader, Content: content, Links: links, Headers: headers, Footers: footers, Images: images}, nil
}

// streamToByte returns an array of the bytes contained in the file reader.
func streamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}
