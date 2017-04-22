package controller

import (
	"forma_shared/lib"
	"forma_shared/model"
	"html/template"
	"log"
	"net/http"
)

// NewPassword controller view inde.gohtml
func NewPassword(w http.ResponseWriter, r *http.Request) {

	var flashes []interface{}

	tpl := template.Must(template.New("NewPassword").ParseFiles("view/newPassword.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml"))

	if r.Method == "POST" {
		u := &model.User{}

		pass1 := r.FormValue("pass1")
		pass2 := r.FormValue("pass2")

		u.Mail = lib.GetSessionsValues(w, r, "email").(string)
		if err := u.SearchUser().Error; err != nil {
			log.Println(err)
			flashes = lib.SetSessionsFlashes(w, r, "Utilisateur non trouvé")
		} else {
			if pass1 == pass2 {
				u.Password = lib.EncryptionPassword(pass1)
				lib.SetSessionsValues(w, r, "email", u.Mail)
				lib.SetSessionsValues(w, r, "firstname", u.Firstname)
				lib.SetSessionsValues(w, r, "lastname", u.Lastname)
				lib.SetSessionsValues(w, r, "ip", u.IP)
				lib.SetSessionsValues(w, r, "admin", u.Admin)
				lib.SetSessionsValues(w, r, "id", u.ID)
				if u.Active {
					lib.SetSessionsValues(w, r, "active", u.Active)
				}
				u.UpdateUser()
				// flashes = lib.SetSessionsFlashes(w, r, "L'enregistrement c'est bien passé. Vous allez recevoir un email avec un code de validation")
				// go lib.SendEmail(u.Mail)

				http.Redirect(w, r, "/", http.StatusFound)
			} else {
				flashes = lib.SetSessionsFlashes(w, r, "Les mots de passes ne sont pas pareils.")
			}

		}

	}
	m := make(map[string]interface{})
	m["title"] = "New Password"
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
