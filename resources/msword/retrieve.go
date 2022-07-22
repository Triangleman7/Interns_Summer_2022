package msword

import (
	"archive/zip"
	"fmt"
	"log"
	"regexp"
	"strings"
)

// retrieveWordDoc parses through the Word Document ZIP archive (files) to search for the
// XML document "word/document.xml", which contains the formatting for the body text contained in
// the Word Document.
//
// If the target XML document is found in the ZIP archive, a pointer to its corresponding zip.File
// (file) object is returned.
//
// An error (err) is raised if the target XML document cannot be found in the ZIP archive.
func retrieveWordDoc(files []*zip.File) (file *zip.File, err error) {
	var filename string = "word/document.xml"
	log.Printf("Searching for file in ZIP archive: %s", filename)

	// Traverse ZIP archive to search for XML document
	for _, f := range files {
		if f.Name == filename {
			log.Printf("Found file in ZIP archive: %s", f.Name)
			return f, err
		}
	}

	// Target XML document not found
	err = fmt.Errorf("could not find file in ZIP archive: %s", filename)
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
	var filename string = "word/_rels/document.xml.rels"
	log.Printf("Searching for file in ZIP archive: %s", filename)

	// Traverse ZIP archive to search for XML document
	for _, f := range files {
		if f.Name == filename {
			log.Printf("Found file in ZIP archive: %s", f.Name)
			return f, err
		}
	}

	// Target XML document not found
	err = fmt.Errorf("could not find file in ZIP archive: %s", filename)
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
	var regex *regexp.Regexp
	regex, err = regexp.Compile(`(headers|footers)\d+\.xml`)
	if err != nil {
		return
	}
	log.Printf("Searching for files in ZIP archive: %v", regex)

	// Traverse ZIP archive to search for header/footer XML documents
	for _, f := range files {
		var submatch []string = regex.FindStringSubmatch(f.Name)
		if len(submatch) == 0 {
			continue
		}

		switch submatch[1] {
		// Header XML document
		case "header":
			log.Printf("Found header file in ZIP archive: %s", submatch[1])
			headers = append(headers, f)
		// Footer XML document
		case "footer":
			log.Printf("Found footer file in ZIP archive: %s", submatch[1])
			footers = append(footers, f)
		}
	}

	// No header or footer XML documents found
	if len(headers) == 0 && len(footers) == 0 {
		err = fmt.Errorf("could not find files in ZIP archive: %v", regex)
	}
	return
}

// retrieveImageFilenames parses through the Word Document ZIP archive (files) to find the
// "word/media/" directory, which contains the image files contained in the Word Document.
//
// A map of the filenames of all images found in the ZIP archive mapped to empty string is
// returned.
func retrieveImageFilenames(files []*zip.File) (map[string]string, error) {
	var dirname string = "word/media"
	log.Printf("Searching for files in ZIP archive directory: %s", dirname)

	// Traverse ZIP archive to search for image files
	images := make(map[string]string)
	for _, f := range files {
		var filename string = f.Name
		if strings.HasPrefix(filename, dirname) {
			log.Printf("Found image file in ZIP archive: %s", filename)
			images[filename] = ""
		}
	}

	return images, nil
}
