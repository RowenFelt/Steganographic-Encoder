package main

import (
    "net/http"
	"io/ioutil"
    "fmt"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

    // Parse our multipart form, 10 << 20 specifies a maximum
    // upload of 10 MB files.
    r.ParseMultipartForm(10 << 20)
    // FormFile returns the first file for the given key `myFile`
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file
	for _, element := range [2]string{"file1", "file2"} {
		fmt.Printf("file refered to as %s\n", element)
		file, handler, err := r.FormFile(element)
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		// Create a temporary file within our temp-images directory that follows
		// a particular naming pattern
		tempFile, err := ioutil.TempFile("temp-images", "upload-*.JPG")
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		// write this byte array to our temporary file
		tempFile.Write(fileBytes)
		// return that we have successfully uploaded our file!
	}
	fmt.Fprintf(w, "Successfully Uploaded Files\n")
}

func setupRoutes() {
    http.HandleFunc("/encode", uploadFile)
    http.ListenAndServe(":5000", nil)
}

func main() {
    fmt.Println("hello World")
	setupRoutes()
}
