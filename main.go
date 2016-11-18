package main

import (
	"fmt"
	"html/template"
	"io"
	"math"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type myFile struct {
	Name   string
	Folder string
	Size   float64
	Ext    string
	Time   time.Time
}

type user struct {
	Name string
	IP   string
}

var users = []user{
	{"Sam", "::1"},
	{"Lucas", "172.16.29.150"},
	{"Francois", "172.16.30.25"},
	{"Laurent", "172.16.29.193"},
	{"Brice", "172.16.27.30"},
	{"Patrice", "172.16.26.224"},
	{"Andre", "172.16.27.182"},
	{"Seb", "172.16.27.49"},
	{"Pierre", "172.16.30.3"},
	{"Mouloud", "172.16.30.66"},
	{"Lea", "172.16.27.113"},
	{"Richard", "172.16.24.172"},
	{"Remi", "172.16.26.224"},
}

// DIRFILE constante link download and upload file
const DIRFILE = "/home/sam/go/src/forma_shared/"

var autorized = false

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for i, addr := range addrs {
		fmt.Printf("%d %v\n", i, addr)
	}

	port := ":9000"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/download/{folder}/{file}", download)
	router.HandleFunc("/upload", upload)
	router.HandleFunc("/not_access", notAccess)
	router.HandleFunc("/annuaire", annuaire)

	router.PathPrefix("/files").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir("files"))))
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server start : ", time.Now(), " to port 9000")

	http.ListenAndServe(port, router)
}

// func walkpath(path string, f os.FileInfo, err error) error {
// 	if !f.IsDir() {
// 		fmt.Printf("%s with %d bytes\n", path, f.Size())
// 	}
// 	return nil
// }

func checkFilesInFolder(folder string) bool {
	count := 0
	filepath.Walk(folder, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			count++
		}
		return nil
	})
	if count > 0 {
		return true
	} else {
		return false
	}
}

func annuaire(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("annuaire.gohtml")
	ip, autorize := checkIp(w, r)
	fmt.Println(ip, autorize)
	if !autorize {
		http.Redirect(w, r, "/not_access", 301)
	}
	m := make(map[string]interface{})
	m["title"] = "Annuaire"
	tpl.ExecuteTemplate(w, "annuaire.gohtml", m)
}

func checkIp(w http.ResponseWriter, r *http.Request) (string, bool) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	for _, v := range users {
		if v.IP == ip {
			autorized = true
			return ip, autorized
		}
	}
	return ip, false
}

func notAccess(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("not_access.gohtml")
	ip, _ := checkIp(w, r)
	m := make(map[string]interface{})
	m["ip"] = ip
	tpl.ExecuteTemplate(w, "not_access.gohtml", m)
}

func index(w http.ResponseWriter, r *http.Request) {
	// ip, port, err := net.SplitHostPort(r.RemoteAddr)
	ip, autorize := checkIp(w, r)
	fmt.Println(ip, autorize)
	// if !autorize {
	// 	http.Redirect(w, r, "/not_access", 301)
	// }

	// files, _ := filepath.Glob("DIRFILE + files/*")
	files := []myFile{}
	filepath.Walk(DIRFILE+"files/", func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			// fmt.Println(path)
			file := &myFile{}
			file.Name = f.Name()
			// fmt.Println(path[len(DIRFILE+"files/"):])
			if f.Name() != strings.Split(path[len(DIRFILE):], "/")[1] {
				file.Folder = strings.Split(path[len(DIRFILE):], "/")[1]
				// fmt.Println(file.Folder)
			} else {
				file.Folder = ""
			}
			file.Size = float64(f.Size())
			file.Time = f.ModTime()
			file.Ext = filepath.Ext(f.Name())[1:]

			files = append(files, *file)
			// fmt.Printf("%s with %d bytes\n", path, f.Size())
		} else if f.IsDir() && f.Name() != "files" {
			if checkFilesInFolder(path) {
				// fmt.Println(f.Size())
				folder := &myFile{}

				folder.Folder = strings.Split(path[len(DIRFILE):], "/")[1]

				files = append(files, *folder)
			}
		}
		return nil
	})

	const layout = "Mon 02 Jan 2006"

	funcMap := template.FuncMap{
		"title": strings.Title,
		"time_fr": func(t *time.Time) string {
			return t.Format(layout)
		},
		"exp": func(i float64) string {
			if i > math.Pow(10, 12) {
				return strconv.FormatFloat(i/math.Pow(10, 12), 'f', 2, 64) + " To"
			} else if i > math.Pow(10, 9) {
				return strconv.FormatFloat(i/math.Pow(10, 9), 'f', 2, 64) + " Go"
			} else if i > math.Pow(10, 6) {
				return strconv.FormatFloat(i/math.Pow(10, 6), 'f', 2, 64) + " Mo"
			} else if i > math.Pow(10, 3) {
				return strconv.FormatFloat(i/math.Pow(10, 3), 'f', 2, 64) + " ko"
			} else {
				ip, _, _ := net.SplitHostPort(r.RemoteAddr)
				fmt.Println(ip)
				return strconv.FormatFloat(i, 'f', -1, 64) + " octets"
			}
		},
	}

	tpl := template.Must(template.New("Partage").Funcs(funcMap).ParseGlob("*.gohtml"))

	m := make(map[string]interface{})
	m["files"] = &files
	tpl.ExecuteTemplate(w, "index.gohtml", m)
}

func upload(w http.ResponseWriter, r *http.Request) {
	ip, autorize := checkIp(w, r)
	fmt.Println(ip, autorize)
	// if !autorize {
	// 	http.Redirect(w, r, "/not_access", 301)
	// }

	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		files := []myFile{}
		filepath.Walk(DIRFILE+"files/", func(path string, f os.FileInfo, err error) error {
			if f.IsDir() && f.Name() != "files" {
				folder := &myFile{}

				folder.Folder = strings.Split(path[len(DIRFILE):], "/")[1]

				files = append(files, *folder)
			}
			return nil
		})
		tpl, _ := template.ParseFiles("upload.gohtml")

		m := make(map[string]interface{})
		m["folder"] = &files
		tpl.ExecuteTemplate(w, "upload.gohtml", m)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")

		for _, v := range users {
			fmt.Println(ip)
			if v.IP == ip {
				handler.Filename = v.Name + "_" + handler.Filename
			} else if v.IP == "" {
				handler.Filename = "invite_" + handler.Filename
			}
		}

		// handler.Filename = handler.Filename
		folder := r.FormValue("folder")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		// fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile(DIRFILE+"files/"+folder+"/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
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
	ip, autorize := checkIp(w, r)
	fmt.Println(ip, autorize)
	// if !autorize {
	// 	// http.RedirectHandler("/not_access", 403)
	// 	http.Redirect(w, r, "/not_access", 301)
	// }
	vars := mux.Vars(r)
	folder := vars["folder"]
	fileName := vars["file"]

	fmt.Println("Filename :", fileName)
	if fileName == "" {
		//Get not set, send a 400 bad request
		http.Error(w, "Get 'file' not specified in url.", 400)
		return
	}
	fmt.Println("Client requests: " + fileName)

	//Check if file exists and open
	Openfile, err := os.Open(DIRFILE + "files/" + folder + "/" + fileName)
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
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(w, Openfile) //'Copy' the file to the client
	return
}
