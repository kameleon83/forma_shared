package controllerView

import (
	"fmt"
	"forma_shared/controller"
	"forma_shared/model"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
)

// Upload file user
func Upload(w http.ResponseWriter, r *http.Request) {
	ip, autorize := controller.CheckIP(w, r)
	controller.WriteLog("Upload : ", ip, controller.AfficheNom(ip), strconv.FormatBool(autorize))
	// if !autorize {
	// 	http.Redirect(w, r, "/not_access", 301)
	// }
	if r.Method == "GET" {
		folderAndFile := controller.ReadJSON(".config.json", &model.FolderFile{})
		tpl, _ := template.ParseFiles("view/upload.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml")

		m := make(map[string]interface{})
		m["title"] = "Upload"
		m["folder"] = &folderAndFile.Folder
		m["ip_name"] = controller.AfficheNom(ip)

		tpl.ExecuteTemplate(w, "layout", m)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		users := model.ReadUserJSON(false, "lastname")
		nameIP, name := false, ""
		for _, v := range users {
			// fmt.Println(v.IP, ip)
			if v.IP == ip {
				nameIP = true
				name = v.Firstname
				break
			}
		}

		if nameIP {
			handler.Filename = name + "_" + handler.Filename
		} else {
			handler.Filename = "invite_" + handler.Filename
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
			controller.WriteLog("Upload de " + handler.Filename + " ok!!")
			http.Redirect(w, r, "/refresh", http.StatusFound)
		}
	}
}
