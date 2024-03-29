package msword

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"strings"
)

// WriteToFile writes the contents of the Word Document d to a file at path.
//
// Raises any errors encountered while writing the Word Document contents.
func (d *Docx) WriteToFile(path string) (err error) {
	// Create target file at path
	var file *os.File
	file, err = os.Create(path)
	if err != nil {
		return
	}
	defer file.Close()
	log.Printf("Destination file created: %s", path)

	// Write document contents to target file
	err = d.Write(file)

	log.Printf("Word Document contents successfully written to %s", path)
	return
}

// Write writes each file in the Word Document ZIP archive to the target file ioWriter.
//
// A new ZIP archive of Word Document files is created at the target path via ioWriter. Each aspect
// of the Word Document d (e.g., body, header, footer, hyperlink, image) and other miscellaneous
// files is written to the appropriate file in the new ZIP archive.
//
// Raises any errors encountered while writing the Word Document contents.
func (d *Docx) Write(ioWriter io.Writer) (err error) {
	var zipWriter *zip.Writer = zip.NewWriter(ioWriter)
	log.Print("Created ZIP archive writer")

	// Iterate through ZIP archive Word Document XML files
	for _, file := range d.Files {
		var writer io.Writer
		var closer io.ReadCloser

		// Create target file within the ZIP archive
		writer, err = zipWriter.Create(file.Name)
		if err != nil {
			return err
		}
		log.Printf("Destination file created in ZIP archive: %s", file.Name)

		// Open the target file
		closer, err = file.Open()
		if err != nil {
			return err
		}
		log.Printf("Opened origin file: %s", file.Name)

		if file.Name == "word/document.xml" {
			// Write content of Word Document body to appropriate XML document
			writer.Write([]byte(d.Content))
		} else if file.Name == "word/_rels/document.xml.rels" {
			// Write content of Word Document hyperlinks to appropriate XML document
			writer.Write([]byte(d.Links))
		} else if strings.Contains(file.Name, "header") && d.Headers[file.Name] != "" {
			// Write content of Word Document headers to appropriate XML document
			writer.Write([]byte(d.Headers[file.Name]))
		} else if strings.Contains(file.Name, "footer") && d.Footers[file.Name] != "" {
			// Write content of Word Document footers to appropriate XML document
			writer.Write([]byte(d.Footers[file.Name]))
		} else if strings.HasPrefix(file.Name, "word/media/") && d.Images[file.Name] != "" {
			// Write content of Word Document image to appropriate directory
			log.Printf("Replacing image %s: %s", file.Name, d.Images[file.Name])
			var new *os.File
			new, err = os.Open(d.Images[file.Name])
			log.Printf("Opened new image: %s", d.Images[file.Name])
			if err != nil {
				return err
			}
			writer.Write(streamToByte(new))
			new.Close()
			log.Printf("Contents of new image written: %s", d.Images[file.Name])
		} else {
			// Write content of miscellaneous file to appropriate file
			writer.Write(streamToByte(closer))
		}

		log.Printf("File contents successfully written to ZIP archive: %s", file.Name)
	}
	zipWriter.Close()
	log.Print("Closed ZIP archive writer")

	log.Print("Word Document files successfully written to ZIP archive")
	return
}
