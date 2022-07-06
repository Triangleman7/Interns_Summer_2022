package server

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func UploadFile(file multipart.File, header *multipart.FileHeader) string {
	var err error

	var path string = filepath.Join(TEMPDIRECTORY, header.Filename)

	destination, err := os.Create(path)
	if err != nil { panic(err) }

	_, err = io.Copy(destination, file)
	if err != nil { panic(err) }

	destination.Close()

	return path
}
