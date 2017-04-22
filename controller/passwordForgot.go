package controller

import (
	"forma_shared/lib"
	"forma_shared/model"
	"html/template"
	"net/http"
)

// PasswordForgot controller view inde.gohtml
func PasswordForgot(w http.ResponseWriter, r *http.Request) {

	var flashes []interface{}

	tpl := template.Must(template.New("PasswordForgot").ParseFiles("view/passwordForgot.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml"))

	if r.Method == "POST" {
		u := model.User{}

		u.Mail = r.FormValue("email")

		if u.SearchUser().Error != nil {
			flashes = lib.SetSessionsFlashes(w, r, "Cet Utilisateur n'est pas enregistré")
		} else {
			go lib.SendEmailRescue(u.Mail)
			u.Password = lib.EncryptionEmailRescue(u.Mail)
			flashes = lib.SetSessionsFlashes(w, r, "Check tes mails pour rentrer ton mot de passe provisoire :-P")
			lib.SetSessionsValues(w, r, "email", u.Mail)
			lib.SetSessionsValues(w, r, "rescue", true)
			u.UpdateUser()
			http.Redirect(w, r, "/login", http.StatusFound)
		}
	}
	m := make(map[string]interface{})
	m["title"] = "PasswordForgot"
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
