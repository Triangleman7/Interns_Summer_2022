package msword

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func (d *Docx) WriteToFile(path string) (err error) {
	var target *os.File
	target, err = os.Create(path)
	if err != nil { return }

	defer target.Close()
	err = d.Write(target)
	return
}

func (d *Docx) Write(ioWriter io.Writer) (err error) {
	w := zip.NewWriter(ioWriter)
	for _, file := range d.files {
		var writer io.Writer
		var readCloser io.ReadCloser

		writer, err = w.Create(file.Name)
		if err != nil { return err }

		readCloser, err = file.Open()
		if err != nil { return err }

		if file.Name == "word/document.xml" {
			writer.Write([]byte(d.content))
		} else if file.Name == "word/_rels/document.xml.rels" {
			writer.Write([]byte(d.links))
		} else if strings.Contains(file.Name, "header") && d.headers[file.Name] != "" {
			writer.Write([]byte(d.headers[file.Name]))
		} else if strings.Contains(file.Name, "footer") && d.footers[file.Name] != "" {
			writer.Write([]byte(d.footers[file.Name]))
		} else if strings.HasPrefix(file.Name, "word/media/") && d.images[file.Name] != "" {
			newImage, err := os.Open(d.images[file.Name])
			if err != nil { return err }
			writer.Write(streamToByte(newImage))
			newImage.Close()
		} else {
			writer.Write(streamToByte(readCloser))
		}
	}
	w.Close()
	return
}