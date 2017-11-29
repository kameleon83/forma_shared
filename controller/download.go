package controller

import (
	"forma_shared/lib"
	"io"
	"net/http"
	"os"
	"strconv"

	"forma_shared/model"

	"github.com/gorilla/mux"
)

// Download View
func Download(w http.ResponseWriter, r *http.Request) {
	lib.GetSessionLogin(w, r)

	vars := mux.Vars(r)
	folder := vars["folder"]
	fileName := vars["file"]

	// fmt.Println("Filename :", fileName)
	if fileName == "" {
		//Get not set, send a 400 bad request
		http.Error(w, "Get 'file' not specified in url.", 400)
		return
	}
	// fmt.Println("Client requests: " + fileName)

	//Check if file exists and open
	config := new(model.Config)
	Openfile, err := os.Open(config.SendDirectory() + folder + os.PathListSeparator + fileName)
	defer Openfile.Close() //Close after function return
	if err != nil {
		//File not found, send 404
		http.Error(w, "File not found.", 404)
		return
	}

	//File is found, create and send the correct headers

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	// fmt.Println(FileSize)

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(w, Openfile) //'Copy' the file to the client
	return
}
