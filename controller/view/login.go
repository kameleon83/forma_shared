package controllerView

import (
	"forma_shared/controller"
	"forma_shared/model"
	"html/template"
	"net/http"
)

// Login controller view inde.gohtml
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	tpl := template.Must(template.New("Login").ParseFiles("view/login.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml"))

	var flashes string

	if r.Method == "POST" {

		if rescue := controller.GetSessionsValues(w, r, "rescue"); rescue != nil && rescue.(bool) {
			http.Redirect(w, r, "/newPassword", http.StatusFound)
		}

		password := controller.EncryptionPassword(r.FormValue("password"))

		u := model.User{}
		u.Mail = r.FormValue("email")
		u.SearchUser()

		if u.Mail != "" {
			if u.Password == password {
				controller.SetSessionsValues(w, r, "email", u.Mail)
				controller.SetSessionsValues(w, r, "firstname", u.Firstname)
				controller.SetSessionsValues(w, r, "lastname", u.Lastname)
				controller.SetSessionsValues(w, r, "ip", u.IP)
				controller.SetSessionsValues(w, r, "admin", u.Admin)
				controller.SetSessionsValues(w, r, "id", u.ID)
				controller.SetSessionsValues(w, r, "prof", u.Prof)
				if u.Active {
					controller.SetSessionsValues(w, r, "active", u.Active)

					http.Redirect(w, r, "/", http.StatusFound)
				} else {
					flashes = controller.SetSessionsFlashes(w, r, "Votre compte n'est pas actif!! Vérifier vos courriels")
					http.Redirect(w, r, "/valid", http.StatusFound)
				}

				flashes = controller.SetSessionsFlashes(w, r, "Il y a une erreur d'authentification. Est-ce le bon mot de passe? Le bon utilisateur?")
			}
		}
		flashes = controller.SetSessionsFlashes(w, r, "Il y a une erreur d'authentification. Est-ce le bon mot de passe? Le bon utilisateur?")
	}

	m := make(map[string]interface{})
	m["title"] = "Login"
	m["errors"] = flashes
	m["email"] = controller.GetSessionsValues(w, r, "email")
	m["active"] = controller.GetSessionsValues(w, r, "active")
	m["firstname"] = controller.GetSessionsValues(w, r, "firstname")
	m["prof"] = controller.GetSessionsValues(w, r, "prof")
	m["niveau"] = controller.GetSessionsValues(w, r, "niveau")
	m["numberFiles"] = model.COUNTFILES

	tpl.ExecuteTemplate(w, "layout", m)
}
