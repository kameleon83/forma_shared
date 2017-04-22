package controller

import (
	"forma_shared/lib"
	"forma_shared/model"
	"html/template"
	"net/http"
)

// Login controller view inde.gohtml
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	tpl := template.Must(template.New("Login").ParseFiles("view/login.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml"))

	var flashes []interface{}

	if r.Method == "POST" {

		u := model.User{}
		u.Mail = r.FormValue("email")

		if rescue := lib.GetSessionsValues(w, r, "rescue"); rescue != nil && rescue.(bool) {
			if r.FormValue("password") == lib.EncryptionEmailRescue(u.Mail) {
				http.Redirect(w, r, "/newPassword", http.StatusFound)
			}
			flashes = lib.SetSessionsFlashes(w, r, "Vous avez perdu votre mot de passe. Un mot de passe provisoire vous a été envoyé. Veuillez le rentrer dans le champ mot de passe")
		}

		password := lib.EncryptionPassword(r.FormValue("password"))
		u.SearchUser()

		if u.Mail != "" {
			if u.Password == password {
				lib.SetSessionsValues(w, r, "email", u.Mail)
				lib.SetSessionsValues(w, r, "firstname", u.Firstname)
				lib.SetSessionsValues(w, r, "lastname", u.Lastname)
				lib.SetSessionsValues(w, r, "ip", u.IP)
				lib.SetSessionsValues(w, r, "admin", u.Admin)
				lib.SetSessionsValues(w, r, "id", u.ID)
				lib.SetSessionsValues(w, r, "prof", u.Prof)
				if u.Active {
					lib.SetSessionsValues(w, r, "active", u.Active)

					http.Redirect(w, r, "/", http.StatusFound)
				} else {
					flashes = lib.SetSessionsFlashes(w, r, "Votre compte n'est pas actif!! Vérifier vos courriels")
					http.Redirect(w, r, "/valid", http.StatusFound)
				}

				flashes = lib.SetSessionsFlashes(w, r, "Il y a une erreur d'authentification. Est-ce le bon mot de passe? Le bon utilisateur?")
			}
		}
		flashes = lib.SetSessionsFlashes(w, r, "Il y a une erreur d'authentification. Est-ce le bon mot de passe? Le bon utilisateur?")
	}

	m := make(map[string]interface{})
	m["title"] = "Login"
	if len(flashes) > 0 {
		m["errors"] = flashes[0]
	}
	m["email"] = lib.GetSessionsValues(w, r, "email")
	m["active"] = lib.GetSessionsValues(w, r, "active")
	m["firstname"] = lib.GetSessionsValues(w, r, "firstname")
	m["prof"] = lib.GetSessionsValues(w, r, "prof")
	m["niveau"] = lib.GetSessionsValues(w, r, "niveau")
	m["numberFiles"] = model.COUNTFILES

	tpl.ExecuteTemplate(w, "layout", m)
}
