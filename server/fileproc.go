package server

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func UploadFile(file multipart.File, header *multipart.FileHeader) string {
	var err error

	// Construct target path (relative to root directory) to uploaded file
	var path string = filepath.Join(TEMPDIRECTORY, header.Filename)

	// Create file at target path
	destination, err := os.Create(path)
	if err != nil { panic(err) }

	// Copy received file contents to target file in local directory
	_, err = io.Copy(destination, file)
	if err != nil { panic(err) }

	// Clean-up
	err = destination.Close()
	if err != nil { panic(err) }

	return path
}
