package controller

import (
	"forma_shared/lib"
	"forma_shared/model"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Machiel/slugify"
)

// Upload file user
func Upload(w http.ResponseWriter, r *http.Request) {
	lib.GetSessionLogin(w, r)

	if r.Method == "GET" {
		folder := &model.Folder{}
		tpl := template.Must(template.New("Upload").ParseFiles("view/upload.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml"))

		m := make(map[string]interface{})
		m["title"] = "Upload"
		m["folder"] = folder.Search()
		m["email"] = lib.GetSessionsValues(w, r, "email")
		m["active"] = lib.GetSessionsValues(w, r, "active")
		m["firstname"] = lib.GetSessionsValues(w, r, "firstname")
		m["prof"] = lib.GetSessionsValues(w, r, "prof")
		m["niveau"] = lib.GetSessionsValues(w, r, "niveau")
		m["numberFiles"] = model.COUNTFILES

		// m["ip_name"] = lib.AfficheNom(ip)

		tpl.ExecuteTemplate(w, "layout", m)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")

		handler.Filename = lib.GetSessionsValues(w, r, "firstname").(string) + "_" + slug(handler.Filename)

		// handler.Filename = handler.Filename
		folder := r.FormValue("folder")
		if err != nil {
			log.Println(err)
			return
		}
		defer file.Close()
		// fmt.Fprintf(w, "%v", handler.Header)
		config := new(model.Config)
		f, err := os.OpenFile(config.SendDirectory()+folder+"/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Println(err)
			return
		}
		defer f.Close()
		_, err = io.Copy(f, file)
		if err == nil {
			lib.WriteLog("Upload de " + handler.Filename + " ok!!")
			http.Redirect(w, r, "/refresh", http.StatusFound)
		}
	}
}

func slug(fileName string) string {
	ext := filepath.Ext(fileName)
	fileName = fileName[:len(fileName)-len(ext)]
	fileName = slugify.Slugify(fileName)
	ext = slugify.Slugify(ext[1:])
	return fileName + "." + ext
}
