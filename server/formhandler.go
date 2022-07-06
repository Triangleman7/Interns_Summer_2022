package server

func HandleFormPrimary(w http.ResponseWriter, r *http.Request) {
	var res string

	// Parse form submission
	err = r.ParseMultipartForm(0);
	if err != nil { panic(err) }

	// Get values from form submission
	var valueTextField string = r.FormValue("primary-text")
	var valueMenu string = r.FormValue("primary-text-operation")
	file, header, err := r.FormFile("primary-image")
	if err != nil { panic(err) }

	// Format value of text field
	err, res = FormatValue(valueTextField, valueMenu)
	if err != nil { panic(err) }

	// Upload `file`
	var uploadpath string = UploadFile(file, header)

	// Clean-up
	err = file.Close()
	if err != nil { panic(err) }

	// Debugging
	fmt.Fprintf(w, "<form name=\"primary\">: Text Field value = \"%v\"\n", valueTextField)
	fmt.Fprintf(w, "<form name=\"primary\">: Dropdown Menu value = \"%v\"\n", valueMenu)
	fmt.Fprintf(w, "<form name=\"primary\">: Image Field value = \"%v\"\n", header.Filename)
	fmt.Fprintf(w, "\tUploaded: \"%v\"\n", uploadpath)
	fmt.Fprintf(w, "\tFormatted: \"%v\"\n", res)

	// Write formatted value of text field to output files
	var htmlpath string = filepath.Join(OUTPUTDIRECTORY, "primary-text.html")
	var docpath string = filepath.Join(OUTPUTDIRECTORY, "primary-text.doc")
	outputdata.WriteOutput(htmlpath, "outputdata/templates/template.html", PERMISSIONBITS, res)
	outputdata.WriteOutput(docpath, "outputdata/templates/template.doc", PERMISSIONBITS, res)
}