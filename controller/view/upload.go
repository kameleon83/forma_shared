package controllerView

import (
	"forma_shared/controller"
	"forma_shared/model"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

// Upload file user
func Upload(w http.ResponseWriter, r *http.Request) {
	controller.GetSessionLogin(w, r)

	if r.Method == "GET" {
		f := &model.File{}
		tpl := template.Must(template.New("Upload").ParseFiles("view/upload.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml"))

		m := make(map[string]interface{})
		m["title"] = "Upload"
		m["folder"] = f.SearchFolder()
		m["email"] = controller.GetSessionsValues(w, r, "email")
		m["active"] = controller.GetSessionsValues(w, r, "active")
		m["firstname"] = controller.GetSessionsValues(w, r, "firstname")
		m["prof"] = controller.GetSessionsValues(w, r, "prof")
		m["niveau"] = controller.GetSessionsValues(w, r, "niveau")
		m["numberFiles"] = model.COUNTFILES

		// m["ip_name"] = controller.AfficheNom(ip)

		tpl.ExecuteTemplate(w, "layout", m)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")

		handler.Filename = controller.GetSessionsValues(w, r, "firstname").(string) + "_" + handler.Filename

		// handler.Filename = handler.Filename
		folder := r.FormValue("folder")
		if err != nil {
			log.Println(err)
			return
		}
		defer file.Close()
		// fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile(controller.DIRFILE+folder+"/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Println(err)
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
