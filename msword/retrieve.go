package msword

import (
	"archive/zip"
	"errors"
	"strings"
)

func retrieveWordDoc(files []*zip.File) (file *zip.File, err error) {
	for _, f := range files {
		if f.Name == "word/document.xml" { file = f }
	}
	if file == nil {
		err = errors.New("document.xml file not found")
	}
	return
}

func retrieveHeaderFooterDoc(files []*zip.File) (headers []*zip.File, footers []*zip.File, err error) {
	for _, f := range files {

		if strings.Contains(f.Name, "header") {
			headers = append(headers, f)
		}
		if strings.Contains(f.Name, "footer") {
			footers = append(footers, f)
		}
	}
	if len(headers) == 0 && len(footers) == 0 {
		err = errors.New("headers[1-3].xml file not found and footers[1-3].xml file not found")
	}
	return
}

func retrieveLinkDoc(files []*zip.File) (file *zip.File, err error) {
	for _, f := range files {
		if f.Name == "word/_rels/document.xml.rels" { file = f }
	}
	if file == nil {
		err = errors.New("document.xml.rels file not found")
	}
	return
}

func retrieveImageFilenames(files []*zip.File) (map[string]string, error) {
	images := make(map[string]string)
	for _, f := range files {
		if strings.HasPrefix(f.Name, "word/media/") { images[f.Name] = "" }
	}
	return images, nil
}
