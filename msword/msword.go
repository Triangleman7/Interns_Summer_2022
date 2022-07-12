// Package msword handles Microsoft Word Documents.
package msword

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/xml"
	"io"
	"strings"
)

type ZipData interface {
	// files serves the content from a ZIP archive.
	files() []*zip.File

	// close closes the ZIP file.
	close() error
}

type ZipFile struct {
	// data is a reader for the ZIP archive.
	data *zip.ReadCloser
}

func (d ZipFile) files() []*zip.File {
	return d.data.File
}

func (d ZipFile) close() error {
	return d.data.Close()
}

// ReplaceDocx represents a Word Document, the document ZIP archive, and the document contents.
type ReplaceDocx struct {
	// zipReader is a reader of the Word Document ZIP archive.
	zipReader ZipData

	// content is the text content of the body of the Word Document.
	content   string

	// links is the text content of the hyperlinks of the Word Document.
	links     string

	// headers is a map of header style names to the corresponding text content of the Word
	// Document.
	headers   map[string]string

	// footers is a map of footer style names to the corresponding text content of the Word
	// Document.
	footers   map[string]string

	// images is a map of the filenames of the images of the Word Document to empty strings.
	images    map[string]string
}

// Docx represents a Word Document, the document ZIP archive files, and the document contents.
type Docx struct {
	// files is an array of individual files contained in the ZIP archive.
	files   []*zip.File

	// content is the text content of the body of the Word Document.
	content string

	// links is the text content of the hyperlinks of the Word Document.
	links   string

	// headers is a map of header style names to the corresponding text content of the Word
	// Document.
	headers map[string]string

	// footers is a map of footer style names to the corresponding text content of the Word
	// Document.
	footers map[string]string

	// images is a map of the filenames of the images of the Word Document to empty strings.
	images  map[string]string
}

// Close closes the ZIP archive of the Word Document d.
func (r *ReplaceDocx) Close() error {
	return r.zipReader.close()
}

// GetContent returns the content of the Word Document (d) body.
func (d *Docx) GetContent() string {
	return d.content
}

// SetContent sets the content of the Word Document (d) body to content.
func (d *Docx) SetContent(content string) {
	d.content = content
}

// ReadDocxFile reads the Word Document specified by path. The Word Document is read by opening the
// ZIP archive located at path. A Word Document object (type ReplaceDocx) is returned.
//
// Raises any errors encountered while opening the ZIP archive.
func ReadDocxFile(path string) (*ReplaceDocx, error) {
	reader, err := zip.OpenReader(path)
	if err != nil { return nil, err }

	zipData := ZipFile{data: reader}
	return ReadDocx(zipData)
}

// ReadDocx reads the contents of the Word Document ZIP archive (ZipData) and processes each aspect
// of the Word Document (body, hyperlinks, headers, footers, images). A Word Document object (type
// Replace Docx) is returned.
//
// Raises any errors encountered while reading the contents of the Word Document ZIP archive.
func ReadDocx(reader ZipData) (*ReplaceDocx, error) {
	content, err := readText(reader.files())
	if err != nil { return nil, err }

	links, err := readLinks(reader.files())
	if err != nil { return nil, err }

	headers, footers, _ := readHeaderFooter(reader.files())
	images, _ := retrieveImageFilenames(reader.files())
	return &ReplaceDocx{zipReader: reader, content: content, links: links, headers: headers, footers: footers, images: images}, nil
}

// streamToByte returns an array of the bytes contained in the file reader.
func streamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}
