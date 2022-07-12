package msword

import (
	"archive/zip"
	"io"
	"io/ioutil"
)

// readText returns the content of the body of the target Word Document. The return value text is
// the text content of body of the target Word Document.
//
// The ZIP archive, files, of the Word Document files is traversed to search for the appropriate
// XML document, which is then read from to obtain the body text of the target Word Document.
//
// Raises any errors encountered while reading the Word Document text content.
func readText(files []*zip.File) (text string, err error) {
	// Get Word Document body XML document
	var documentFile *zip.File
	documentFile, err = retrieveWordDoc(files)
	if err != nil { return text, err }

	// Read target XML document
	var documentReader io.ReadCloser
	documentReader, err = documentFile.Open()
	if err != nil { return text, err }

	// Get text content of Word Document body
	text, err = wordDocToString(documentReader)

	return
}

// readLinks returns the content of the hyperlinks of the target Word Document. The return value
// text is the text content of all referenced hyperlinks of the target Word Document.
//
// The ZIP archive, files, of the Word Document files is traversed to search for the appropriate
// XML document, which is then read from to obtain the hyperlinks of the target Word Document.
//
// Raises any errors encountered while reading the Word Document text content.
func readLinks(files []*zip.File) (text string, err error) {
	// Get Word Document hyperlinks XML document
	var documentFile *zip.File
	documentFile, err = retrieveLinkDoc(files)
	if err != nil { return text, err }

	// Read target XML document
	var documentReader io.ReadCloser
	documentReader, err = documentFile.Open()
	if err != nil { return text, err }

	// Get text content of Word Document hyperlinks
	text, err = wordDocToString(documentReader)

	return
}

// readHeaderFooter returns the content of each header/footer type of the target Word Document. The
// return value textHeader is a map of each header type and its corresponding text content; the
// return value textFooter is a map of each footer type and its corresponding text content.
//
// The ZIP archive, files, of the Word Document files is traversed to search for the appropriate
// XML documents, which are then individually read from to obtain the headers/footers of the target
// Word Document.
//
// Raises any errors encountered while reading the Word Documents headers and footers.
func readHeaderFooter(files []*zip.File) (textHeader map[string]string, textFooter map[string]string, err error) {
	// Get Word Document header/footer XML documents
	h, f, err := retrieveHeaderFooterDoc(files)
	if err != nil { return map[string]string{}, map[string]string{}, err }

	// Compile headers from Word Document header XML documents
	textHeader, err = buildHeaderFooter(h)
	if err != nil { return map[string]string{}, map[string]string{}, err }

	// Compile footers from Word Document footer XML documents
	textFooter, err = buildHeaderFooter(f)
	if err != nil { return map[string]string{}, map[string]string{}, err }

	return textHeader, textFooter, err
}

// buildHeaderFooter returns a map (textHeaderFooter) of each header/footer type and its
// corresponding text content.
// 
// Raises any errors encountered while reading and processing the XML documents.
func buildHeaderFooter(files []*zip.File) (textHeaderFooter map[string]string, err error) {
	// Iterate through header/footer XML documents
	// (Each iteration equates to processing one type of header/footer)
	for _, element := range files {
		// Read Word Document header/footer XML document
		var documentReader io.ReadCloser
		documentReader, err = element.Open()
		if err != nil { return }

		// Get text content of header/footer
		var text string
		text, err = wordDocToString(documentReader)
		if err != nil { return }

		// Map header/footer type to corresponding text content
		textHeaderFooter[element.Name] = text
	}

	return
}

// wordDocToString reads the contents of reader, which wraps the contents of the XML document that
// contains the formatting for an aspect of the target Word Document. The contents of the XML
// Document are casted to type 'string' and returned.
//
// Raises any errors encountered while reading from reader.
func wordDocToString(reader io.Reader) (string, error) {
	b, err := ioutil.ReadAll(reader)
	if err != nil { return "", err }

	return string(b), nil
}