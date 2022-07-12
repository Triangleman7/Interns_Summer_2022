package msword

import (
	"fmt"
	"strings"
)

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

func (r *ReplaceDocx) Close() error {
	return r.zipReader.close()
}

func (d *Docx) GetContent() string {
	return d.content
}

func (d *Docx) SetContent(content string) {
	d.content = content
}

func (d *Docx) ReplaceRaw(oldString string, newString string, num int) {
	d.content = strings.Replace(d.content, oldString, newString, num)
}

func (d *Docx) Replace(oldString string, newString string, num int) (err error) {
	oldString, err = encode(oldString)
	if err != nil { return err }

	newString, err = encode(newString)
	if err != nil { return err }

	d.content = strings.Replace(d.content, oldString, newString, num)

	return nil
}

func (d *Docx) ReplaceHeader(oldString string, newString string) (err error) {
	return replaceHeaderFooter(d.headers, oldString, newString)
}

func (d *Docx) ReplaceFooter(oldString string, newString string) (err error) {
	return replaceHeaderFooter(d.footers, oldString, newString)
}

func (d *Docx) ReplaceLink(oldString string, newString string, num int) (err error) {
	oldString, err = encode(oldString)
	if err != nil { return err }

	newString, err = encode(newString)
	if err != nil { return err }

	d.links = strings.Replace(d.links, oldString, newString, num)

	return nil
}

func (d *Docx) ReplaceImage(oldImage string, newImage string) (err error) {
	if _, ok := d.images[oldImage]; ok {
		d.images[oldImage] = newImage
		return nil
	}
	return fmt.Errorf("old image: %q, file not found", oldImage)
}

func replaceHeaderFooter(headerFooter map[string]string, oldString string, newString string) (err error) {
	oldString, err = encode(oldString)
	if err != nil { return err }

	newString, err = encode(newString)
	if err != nil { return err }

	for k := range headerFooter {
		headerFooter[k] = strings.Replace(headerFooter[k], oldString, newString, -1)
	}

	return nil
}