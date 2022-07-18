package server

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

// UploadFile saves the contents of file to the local file system, at the path determined by header
// (header.Filename).
//
// Raises any errors encountered while copying the file contents to the local file system.
func UploadFile(file multipart.File, header *multipart.FileHeader) (path string, err error) {
	log.Printf("Uploading origin file to local file system: %s", header.Filename)

	// Construct target path (relative to root directory) to uploaded file
	path = filepath.Join(TEMPDIRECTORY, header.Filename)
	log.Printf("Destination path created: %s", path)

	// Create file at target path
	var destination *os.File
	destination, err = os.Create(path)
	if err != nil {
		return
	}
	defer destination.Close()
	log.Printf("Destination file created: %s", path)

	// Copy received file contents to target file in local directory
	_, err = io.Copy(destination, file)
	if err != nil {
		return
	}
	log.Printf("Contents of origin file copied to destination file: %s", path)

	log.Printf("Successfully uploaded %s to file system at %s", header.Filename, path)
	return
}
