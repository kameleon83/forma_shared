package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// DIRFILE constante link download and upload file
const DIRFILE = "./test/"

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/download/{file}", download)
	router.HandleFunc("/upload", upload)

	http.ListenAndServe(":9000", router)
}

func index(w http.ResponseWriter, r *http.Request) {
	files, _ := filepath.Glob("./test/*")

	for k := range files {
		files[k] = (strings.Split(files[k], "/"))[1]
	}
	fmt.Println(files)
	t, _ := template.ParseFiles("index.gohtml")
	m := make(map[string]interface{})
	m["files"] = files
	t.ExecuteTemplate(w, "index.gohtml", m)
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
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		_, err = io.Copy(f, file)
		if err == nil {
			fmt.Println("Upload de " + handler.Filename + " ok!!")
		}
	}
}

func download(w http.ResponseWriter, r *http.Request) {

	fmt.Println("dans download")

	fmt.Println(r.Host)

	link := "http://" + filepath.Clean(r.Host+r.URL.String())

	fmt.Println("link " + link)

	u, _ := url.Parse(r.URL.String())

	if strings.Split(u.Path, "/")[1] == "download" {
		// fmt.Println(strings.Split(u.Path, "/")[2])
		fileName := strings.Split(u.Path, "/")[2]

		output, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Error while creating", fileName, "-", err)
			return
		}
		defer output.Close()

		response, err := http.Get(link)
		if err != nil {
			fmt.Println("Error while downloading", link, "-", err)
			return
		}
		defer response.Body.Close()

		n, err := io.Copy(output, response.Body)
		if err != nil {
			fmt.Println("Error while downloading", link, "-", err)
			return
		}

		fmt.Println(n, "bytes downloaded.")

	} else {
		http.RedirectHandler("/", 302)
	}

}
