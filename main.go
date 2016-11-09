package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type myFile struct {
	Name string
	Size int64
	Time time.Time
}

// DIRFILE constante link download and upload file
const DIRFILE = "./test/"

func main() {
	port := ":9000"
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("Server start : ", time.Now(), " to port 9000")
	router.HandleFunc("/", index)
	router.HandleFunc("/download/{file}", download)
	router.HandleFunc("/upload", upload)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.ListenAndServe(port, router)
}

// func walkpath(path string, f os.FileInfo, err error) error {
// 	if !f.IsDir() {
// 		fmt.Printf("%s with %d bytes\n", path, f.Size())
// 	}
// 	return nil
// }

func index(w http.ResponseWriter, r *http.Request) {
	// files, _ := filepath.Glob("./files/*")
	files := []myFile{}
	filepath.Walk("./files/", func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			file := &myFile{}
			file.Name = f.Name()
			file.Size = f.Size()
			file.Time = f.ModTime()

			files = append(files, *file)
			// files.Name = f.Name()
			// files.Size = strconv.FormatInt(f.Size(), 10)
			// files.Size = f.Size()
			fmt.Printf("%s with %d bytes\n", path, f.Size())
		}
		return nil
	})

	const layout = "Mon 02 Jan 2006"

	funcMap := template.FuncMap{
		"title": strings.Title,
		"time_fr": func(t *time.Time) string {
			return t.Format(layout)
		},
	}

	fmt.Println(files)
	tpl := template.Must(template.New("Partage").Funcs(funcMap).ParseGlob("*.gohtml"))
	// t, _ := template.ParseFiles("index.gohtml")

	m := make(map[string]interface{})
	m["files"] = &files
	tpl.ExecuteTemplate(w, "index.gohtml", m)
}

// upload logic
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gohtml")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		// fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		_, err = io.Copy(f, file)
		if err == nil {
			fmt.Println("Upload de " + handler.Filename + " ok!!")
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

func download(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Filename := vars["file"]
	if Filename == "" {
		//Get not set, send a 400 bad request
		http.Error(w, "Get 'file' not specified in url.", 400)
		return
	}
	fmt.Println("Client requests: " + Filename)

	//Check if file exists and open
	Openfile, err := os.Open("./files/" + Filename)
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

	fmt.Println(FileSize)

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+Filename)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(w, Openfile) //'Copy' the file to the client
	return

}
