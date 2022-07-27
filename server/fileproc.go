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
	path = filepath.Join(TEMPDIRECTORY, header.Filename)

	// Create destination file
	var destination *os.File
	destination, err = os.Create(path)
	if err != nil {
		return
	}
	defer destination.Close()
	log.Printf("Destination file created: %s", path)

	// Copy contents of received file to destination file
	_, err = io.Copy(destination, file)
	if err != nil {
		return
	}
	log.Printf("Contents of origin file copied to destination file: %s", path)

	log.Printf("Successfully uploaded %s to file system at %s", header.Filename, path)
	return
}

// CopyFile copies the contents of the file located at source to a new file created at destination.
//
// Raises any errors encountered while copying the source file contents to the destination file.
func CopyFile(source string, destination string) (err error) {
	log.Printf("Moving file: %s => %s", source, destination)
	var fSource, fDestination *os.File

	// Read source file
	fSource, err = os.Open(source)
	if err != nil {
		return
	}
	defer fSource.Close()
	log.Printf("Opened source file: %s", source)

	// Create destination file
	fDestination, err = os.Create(destination)
	if err != nil {
		return
	}
	defer fDestination.Close()
	log.Printf("Created destination file: %s", destination)

	// Copy source file contents to destination file
	_, err = io.Copy(fDestination, fSource)
	if err != nil {
		return
	}
	log.Printf("Contents of origin file copied to destination file: %s", destination)

	log.Printf("Successfully moved %s to %s", source, destination)
	return
}
