package controllerView

import (
	"fmt"
	"forma_shared/controller"
	"forma_shared/model"
	"html/template"
	"net/http"
	"strconv"
)

// Register controller view inde.gohtml
func Register(w http.ResponseWriter, r *http.Request) {

	ip, autorize := controller.CheckIP(w, r)
	controller.WriteLog("Index : ", ip, controller.AfficheNom(ip), strconv.FormatBool(autorize))
	controller.ClientAutorize(w, r)

	controller.GetSessions(w, r)

	tpl := template.Must(template.New("Partage").ParseFiles("view/register.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml"))

	// foldersAndFiles := controller.ReadJSON(&model.FolderFile{})

	m := make(map[string]interface{})
	m["title"] = "See or Download"
	// m["files"] = &foldersAndFiles.File
	// m["folder"] = &foldersAndFiles.Folder
	// m["ip_name"] = controller.AfficheNom(ip)

	if r.Method == "POST" {
		u := &model.User{}
		u.Lastname = r.FormValue("lastname")
		u.Firstname = r.FormValue("firstname")
		pass1 := r.FormValue("mdp")
		pass2 := r.FormValue("mdp_verif")
		if controller.CheckEqualsPasswordString(pass1, pass2) != "" {
			u.Password = controller.CheckEqualsPasswordString(pass1, pass2)
		} else {
			fmt.Println("les mdp ne sont pas pareil")
		}
		u.Function = r.FormValue("function")
		u.Phone = r.FormValue("phone")
		u.Mail = r.FormValue("mail_perso")
		if r.FormValue("prof") == "true" {
			u.Prof = true
		} else if r.FormValue("prof") == "false" {
			u.Prof = false
		}
		u.IP, _ = controller.CheckIP(w, r)
		u.Admin = 0

		// log.Println(u)

		u.CreateUser()
	}

	tpl.ExecuteTemplate(w, "layout", m)
}
