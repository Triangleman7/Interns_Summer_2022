package outputdata

import (
	"os"
)

var OutputDirectory string = "out/"
var PermissionBits os.FileMode = 0755

func DirectorySetup() {
	// Remove the target directory and its children
	os.RemoveAll(OutputDirectory)		// Returns `nil` if directory exists

	// Create an empty output directory
	err := os.Mkdir(OutputDirectory, PermissionBits)
	if err != nil { panic(err) }
}
