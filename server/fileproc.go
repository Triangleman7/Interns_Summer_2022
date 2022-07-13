package server

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func UploadFile(file multipart.File, header *multipart.FileHeader) (path string, err error) {
	// Construct target path (relative to root directory) to uploaded file
	path = filepath.Join(TEMPDIRECTORY, header.Filename)

	// Create file at target path
	var destination *os.File
	destination, err = os.Create(path)
	if err != nil {
		return
	}

	// Copy received file contents to target file in local directory
	_, err = io.Copy(destination, file)
	if err != nil {
		return
	}

	// Clean-up
	err = destination.Close()
	if err != nil {
		return
	}

	return
}
