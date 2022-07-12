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
	files() []*zip.File
	close() error
}

type ZipFile struct {
	data *zip.ReadCloser
}

func (d ZipFile) files() []*zip.File {
	return d.data.File
}

func (d ZipFile) close() error {
	return d.data.Close()
}

type ReplaceDocx struct {
	zipReader ZipData
	content   string
	links     string
	headers   map[string]string
	footers   map[string]string
	images    map[string]string
}

type Docx struct {
	files   []*zip.File
	content string
	links   string
	headers map[string]string
	footers map[string]string
	images  map[string]string
}

func ReadDocxFile(path string) (*ReplaceDocx, error) {
	reader, err := zip.OpenReader(path)
	if err != nil { return nil, err }

	zipData := ZipFile{data: reader}
	return ReadDocx(zipData)
}

func ReadDocx(reader ZipData) (*ReplaceDocx, error) {
	content, err := readText(reader.files())
	if err != nil { return nil, err }

	links, err := readLinks(reader.files())
	if err != nil { return nil, err }

	headers, footers, _ := readHeaderFooter(reader.files())
	images, _ := retrieveImageFilenames(reader.files())
	return &ReplaceDocx{zipReader: reader, content: content, links: links, headers: headers, footers: footers, images: images}, nil
}

func streamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

// To get Word to recognize a tab character, the previous text element has to be closed off first.
// This means that if there are multiple consecutive tabs, there are empty <w:t></w:t> in between,
// but it still seems to work correctly in the output document, certainly better than with other
// combinations attempted.
const TAB = "</w:t><w:tab/><w:t>"
const NEWLINE = "<w:br/>"

func encode(s string) (string, error) {
	var b bytes.Buffer
	enc := xml.NewEncoder(bufio.NewWriter(&b))
	if err := enc.Encode(s); err != nil { return s, err }

	output := strings.Replace(b.String(), "<string>", "", 1) // remove string tag
	output = strings.Replace(output, "</string>", "", 1)
	output = strings.Replace(output, "&#xD;&#xA;", NEWLINE, -1) // \r\n (Windows newline)
	output = strings.Replace(output, "&#xD;", NEWLINE, -1)      // \r (earlier Mac newline)
	output = strings.Replace(output, "&#xA;", NEWLINE, -1)      // \n (unix/linux/OS X newline)
	output = strings.Replace(output, "&#x9;", TAB, -1)          // \t (tab)
	return output, nil
}
