package msword

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"strings"
)

// Editable exposes the file contents of the ZIP archive. A Word Document object (type Docx) is
// returned.
func (r *ReplaceDocx) Editable() *Docx {
	return &Docx{
		Files:   r.ZipReader.files(),
		Content: r.Content,
		Links:   r.Links,
		Headers: r.Headers,
		Footers: r.Footers,
		Images:  r.Images,
	}
}

// ReplaceRaw replaces the first num instances of old found in the content of the Word Document (d)
// body with new.
func (d *Docx) ReplaceRaw(old string, new string, num int) {
	d.Content = strings.Replace(d.Content, old, new, num)
}

// Replace replaces the first num instances of old found in the content of the Word Document (d)
// body with new.
//
// Raises any errors encountered while replacing the body of the Word Document body (while encoding
// old/new).
func (d *Docx) Replace(old string, new string, num int) (err error) {
	log.Printf("Body text replacement: %s => %s (%d)", old, new, num)

	old, err = encode(old)
	if err != nil {
		return err
	}

	new, err = encode(new)
	if err != nil {
		return err
	}

	d.Content = strings.Replace(d.Content, old, new, num)
	log.Printf("Successfully performed %d replacements", num)

	return nil
}

// ReplaceLink replaces the first num instances of old found in the content of the Word Document
// (d) hyperlinks with new.
//
// Raises any errors encountered while replacing the content of the Word Document hyperlinks (while
// encoding old/new).
func (d *Docx) ReplaceLink(old string, new string, num int) (err error) {
	log.Printf("Hyperlink replacement: %s => %s (%d)", old, new, num)

	old, err = encode(old)
	if err != nil {
		return err
	}

	new, err = encode(new)
	if err != nil {
		return err
	}

	d.Links = strings.Replace(d.Links, old, new, num)
	log.Printf("Successfully performed %d replacements", num)

	return nil
}

// ReplaceHeader replaces all instances of old found in the content of the Word Document (d)
// headers with new.
func (d *Docx) ReplaceHeader(old string, new string) (err error) {
	return replaceHeaderFooter(d.Headers, old, new)
}

// ReplaceFooter replaces all instances of old found in the content of the Word Document (d)
// footers with new.
func (d *Docx) ReplaceFooter(old string, new string) (err error) {
	return replaceHeaderFooter(d.Footers, old, new)
}

// ReplaceImage replaces all instances of old found in the Word Document (d) images with new.
//
// Raises an error if old cannot be found in the Word Document images.
func (d *Docx) ReplaceImage(old string, new string) (err error) {
	log.Printf("Image replacement: %s => %s", old, new)
	_, exists := d.Images[old]
	if !exists {
		return fmt.Errorf("old image: %q, file not found", old)
	}

	d.Images[old] = new
	
	log.Print("Succesfully performed replacement")
	return nil
}

// replaceHeaderFooter replaces all instances of old found in the content of the Word
// Document (d) headers/footers (headerFooter) with new.
//
// Raises any errors encountered while replacing the content of the Word Document headers/footers
// (while encoding old/new).
func replaceHeaderFooter(headerFooter map[string]string, old string, new string) (err error) {
	log.Printf("Header/Footer text replacement: %s => %s", old, new)
	old, err = encode(old)
	if err != nil {
		return err
	}

	new, err = encode(new)
	if err != nil {
		return err
	}

	for k := range headerFooter {
		headerFooter[k] = strings.Replace(headerFooter[k], old, new, -1)
	}

	log.Print("Successfully performed replacements")
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
	log.Printf("Encoding %s", s)

	// Create a new XML encoding
	var b bytes.Buffer
	enc := xml.NewEncoder(bufio.NewWriter(&b))
	err = enc.Encode(s)
	if err != nil {
		return
	}

	output = b.String()
	// Remove <string> tag so that the tag name ("string") is not replaced
	output = strings.Replace(output, "<string>", "", 1)  // opening string tag
	output = strings.Replace(output, "</string>", "", 1) // closing string tag
	// Replace special characters
	output = strings.Replace(output, "&#xD;&#xA;", NEWLINE, -1) // `\r\n` - Newline (Windows)
	output = strings.Replace(output, "&#xD;", NEWLINE, -1)      // `\r` - Newline (OS X, earlier version)
	output = strings.Replace(output, "&#xA;", NEWLINE, -1)      // `\n` - Newline (Unix/Linux/OS X)
	output = strings.Replace(output, "&#x9;", TAB, -1)          // `\t` - Tab

	log.Printf("Successfully encoded %s: %s", s, output)
	return
}
