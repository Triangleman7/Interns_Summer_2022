package msword

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"
)

// Editable exposes the file contents of the ZIP archive. A Word Document object (type Docx) is
// returned.
func (r *ReplaceDocx) Editable() *Docx {
	return &Docx{
		files:   r.zipReader.files(),
		content: r.content,
		links:   r.links,
		headers: r.headers,
		footers: r.footers,
		images:  r.images,
	}
}

// ReplaceRaw replaces the first num instances of oldString found in the content of the Word
// Document (d) body with newString.
func (d *Docx) ReplaceRaw(oldString string, newString string, num int) {
	d.content = strings.Replace(d.content, oldString, newString, num)
}

// Replace replaces the first num instances of oldString found in the content of the Word Document
// (d) body with newString.
//
// Raises any errors encountered while replacing the body of the Word Document body (while encoding
// oldString/newString).
func (d *Docx) Replace(oldString string, newString string, num int) (err error) {
	oldString, err = encode(oldString)
	if err != nil { return err }

	newString, err = encode(newString)
	if err != nil { return err }

	d.content = strings.Replace(d.content, oldString, newString, num)

	return nil
}

// ReplaceLink replaces the first num instances of oldString found in the content of the Word
// Document (d) hyperlinks with newString.
//
// Raises any errors encountered while replacing the content of the Word Document hyperlinks (while
// encoding oldString/newString).
func (d *Docx) ReplaceLink(oldString string, newString string, num int) (err error) {
	oldString, err = encode(oldString)
	if err != nil { return err }

	newString, err = encode(newString)
	if err != nil { return err }

	d.links = strings.Replace(d.links, oldString, newString, num)

	return nil
}

// ReplaceHeader replaces all instances of oldString found in the content of the Word Document (d)
// headers with newString.
func (d *Docx) ReplaceHeader(oldString string, newString string) (err error) {
	return replaceHeaderFooter(d.headers, oldString, newString)
}

// ReplaceFooter replaces all instances of oldString found in the content of the Word Document (d)
// footers with newString.
func (d *Docx) ReplaceFooter(oldString string, newString string) (err error) {
	return replaceHeaderFooter(d.footers, oldString, newString)
}

// ReplaceImage replaces all instances of oldImage found in the Word Document (d) images with
// newImage.
//
// Raises an error if oldImage cannot be found in the Word Document images
func (d *Docx) ReplaceImage(oldImage string, newImage string) (err error) {
	_, exists := d.images[oldImage]
	if exists {
		d.images[oldImage] = newImage
		return
	}

	return fmt.Errorf("old image: %q, file not found", oldImage)
}

// replaceHeaderFooter replaces all instances of oldString found in the content of the Word
// Document (d) headers/footers (headerFooter) with newString.
//
// Raises any errors encountered while replacing the content of the Word Document headers/footers
// (while encoding oldString/newString).
func replaceHeaderFooter(headerFooter map[string]string, oldString string, newString string) (err error) {
	oldString, err = encode(oldString)
	if err != nil { return err }

	newString, err = encode(newString)
	if err != nil { return err }

	for k, _ := range headerFooter {
		headerFooter[k] = strings.Replace(headerFooter[k], oldString, newString, -1)
	}

	return nil
}

const NEWLINE = "<w:br/>"
// To get Microsoft Word to recognize a tab character, the previous text element first has to be
// closed off. Thus, if there are multiple consecutive tabs, there are empty <w:t></w:t> in
// between, but it still seems to work correctly in the output document, certainly better than with
// other combinations attempted.
const TAB = "</w:t><w:tab/><w:t>"

// encode modifies s so that Word Document content replacement functions (Replace, ReplaceLink,
// ReplaceHeader, ReplaceFooter) can properly perform the replacements. The string s is written to
// an XML encoder, removes the <string> tag, and special characters (tab, newline) are replaced
// with the corresponding XML representation of the character. The resultant string is then
// returned.
//
// Raises any errors encountered while 
func encode(s string) (output string, err error) {
	// Create a new XML encoding
	var b bytes.Buffer
	enc := xml.NewEncoder(bufio.NewWriter(&b))
	err = enc.Encode(s)
	if err != nil { return }

	output = b.String()
	// Remove <string> tag so that the tag name ("string") is not replaced
	output = strings.Replace(output, "<string>", "", 1)				// opening string tag
	output = strings.Replace(output, "</string>", "", 1)			// closing string tag
	// Replace special characters
	output = strings.Replace(output, "&#xD;&#xA;", NEWLINE, -1)		// `\r\n` - Newline (Windows)
	output = strings.Replace(output, "&#xD;", NEWLINE, -1)			// `\r` - Newline (OS X, earlier version)
	output = strings.Replace(output, "&#xA;", NEWLINE, -1)			// `\n` - Newline (Unix/Linux/OS X)
	output = strings.Replace(output, "&#x9;", TAB, -1)				// `\t` - Tab
	
	return
}