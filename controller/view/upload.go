package controllerView

import (
	"fmt"
	"forma_shared/controller"
	"forma_shared/model"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Upload file user
func Upload(w http.ResponseWriter, r *http.Request) {
	ip, autorize := controller.CheckIP(w, r)
	fmt.Println(ip, controller.AfficheNom(ip), autorize)
	// if !autorize {
	// 	http.Redirect(w, r, "/not_access", 301)
	// }

	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		files := []model.MyFile{}
		filepath.Walk(controller.DIRFILE+"files/", func(path string, f os.FileInfo, err error) error {
			if f.IsDir() && f.Name() != "files" {
				folder := &model.MyFile{}

				folder.Folder = strings.Split(path[len(controller.DIRFILE):], "/")[1]

				files = append(files, *folder)
			}
			return nil
		})
		tpl, _ := template.ParseFiles("view/upload.gohtml")

		m := make(map[string]interface{})
		m["folder"] = &files
		tpl.ExecuteTemplate(w, "upload.gohtml", m)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")

		for _, v := range controller.Users {
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
		f, err := os.OpenFile(controller.DIRFILE+"files/"+folder+"/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
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
