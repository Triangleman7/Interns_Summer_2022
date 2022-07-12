package msword

import (
	"archive/zip"
	"io"
	"io/ioutil"
)

func readText(files []*zip.File) (text string, err error) {
	var documentFile *zip.File
	documentFile, err = retrieveWordDoc(files)
	if err != nil { return text, err }

	var documentReader io.ReadCloser
	documentReader, err = documentFile.Open()
	if err != nil { return text, err }

	text, err = wordDocToString(documentReader)
	return
}

func readHeaderFooter(files []*zip.File) (headerText map[string]string, footerText map[string]string, err error) {

	h, f, err := retrieveHeaderFooterDoc(files)

	if err != nil {
		return map[string]string{}, map[string]string{}, err
	}

	headerText, err = buildHeaderFooter(h)
	if err != nil {
		return map[string]string{}, map[string]string{}, err
	}

	footerText, err = buildHeaderFooter(f)
	if err != nil {
		return map[string]string{}, map[string]string{}, err
	}

	return headerText, footerText, err
}

func buildHeaderFooter(headerFooter []*zip.File) (map[string]string, error) {

	headerFooterText := make(map[string]string)
	for _, element := range headerFooter {
		documentReader, err := element.Open()
		if err != nil {
			return map[string]string{}, err
		}

		text, err := wordDocToString(documentReader)
		if err != nil {
			return map[string]string{}, err
		}

		headerFooterText[element.Name] = text
	}

	return headerFooterText, nil
}

func readLinks(files []*zip.File) (text string, err error) {
	var documentFile *zip.File
	documentFile, err = retrieveLinkDoc(files)
	if err != nil { return text, err }

	var documentReader io.ReadCloser
	documentReader, err = documentFile.Open()
	if err != nil { return text, err }

	text, err = wordDocToString(documentReader)
	return
}

func wordDocToString(reader io.Reader) (string, error) {
	b, err := ioutil.ReadAll(reader)
	if err != nil { return "", err }

	return string(b), nil
}