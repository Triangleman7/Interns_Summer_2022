package msword

import (
	"archive/zip"
	"errors"
	"strings"
)

// retrieveWordDoc parses through the Word Document ZIP archive (files) to find the
// XML document "word/document.xml", which contains the formatting for the body text contained in
// the Word Document.
//
// If the target XML document is found in the ZIP archive, a pointer to its corresponding zip.File
// (file) object is returned.
//
// An error (err) is raised if the target XML document cannot be found in the ZIP archive.
func retrieveWordDoc(files []*zip.File) (file *zip.File, err error) {
	for _, f := range files {
		if f.Name == "word/document.xml" {
			file = f
			break
		}
	}

	if file == nil {
		err = errors.New("document.xml file not found")
	}
	return
}

// retrieveHeaderFooterDoc parses through the Word Document ZIP archive (files) to find the XML
// documents that contain the formatting for the headers and footers contained in the Word
// Document.
// 
// The target XML documents have filenames that ahere to the following regular expressions:
// - ^(headers|footers)\d\.xml$
//
// If the target XML documents are found in the ZIP archive, an array of pointers to each
// corresponding zip.File object is returned. Pointers to zip.File objects corresponding to XML
// documents containing formatting for document headers are contained in the array headers;
// Pointers to zip.File objects corresonding to XML documents containing formatting for document
// footers are contained in the array footers.
//
// An error (err) is raised if the target XML documents cannot be found in the ZIP archive.
func retrieveHeaderFooterDoc(files []*zip.File) (headers []*zip.File, footers []*zip.File, err error) {
	for _, f := range files {
		if strings.Contains(f.Name, "header") { headers = append(headers, f) }
		if strings.Contains(f.Name, "footer") { footers = append(footers, f) }
	}

	// No header XML documents or footer XML documents found
	if len(headers) == 0 && len(footers) == 0 {
		err = errors.New("headers[1-3].xml file not found and footers[1-3].xml file not found")
	}
	return
}

// retrieveLinkDoc parses through the Word Document ZIP archive (files) to find the XML document
// "word/_rels/document.xml.rels", which contains the formatting for the hyperlinks contained in
// the Word Document.
//
// If the target XML document is found in the ZIP archive, a pointer to its corresponding zip.File
// (file) object is returned.
//
// An error (err) is raised if the target XML document cannot be found in the ZIP archive.
func retrieveLinkDoc(files []*zip.File) (file *zip.File, err error) {
	for _, f := range files {
		if f.Name == "word/_rels/document.xml.rels" {
			file = f
			break
		}
	}

	if file == nil {
		err = errors.New("document.xml.rels file not found")
	}
	return
}

// retrieveImageFilenames parses through the Word Document ZIP archive (files) to find the
// "word/media/" directory, which contains the image files contained in the Word Document.
//
// A map of the filenames of all images found in the ZIP archive mapped to empty string is
// returned.
func retrieveImageFilenames(files []*zip.File) (map[string]string, error) {
	images := make(map[string]string)
	for _, f := range files {
		if strings.HasPrefix(f.Name, "word/media/") { images[f.Name] = "" }
	}
	return images, nil
}
