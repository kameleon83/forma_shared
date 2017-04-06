package controllerView

import (
	"forma_shared/controller"
	"forma_shared/model"
	"html/template"
	"net/http"
)

// PasswordForgot controller view inde.gohtml
func PasswordForgot(w http.ResponseWriter, r *http.Request) {

	var flashes interface{}

	tpl := template.Must(template.New("PasswordForgot").ParseFiles("view/passwordForgot.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml"))

	if r.Method == "POST" {
		u := model.User{}

		u.Mail = r.FormValue("email")

		if u.SearchUser().Error != nil {
			flashes = controller.SetSessionsFlashes(w, r, "Cet Utilisateur n'est pas enregistré")
		} else {
			go controller.SendEmailRescue(u.Mail)
			u.Password = controller.EncryptionEmailRescue(u.Mail)
			u.UpdateUser()
			flashes = controller.SetSessionsFlashes(w, r, "Check tes mails pour rentrer ton mot de passe provisoire :-P")
			controller.SetSessionsValues(w, r, "email", u.Mail)
			controller.SetSessionsValues(w, r, "rescue", true)
			http.Redirect(w, r, "/login", http.StatusFound)
		}
	}
	m := make(map[string]interface{})
	m["title"] = "PasswordForgot"
	m["errors"] = flashes
	m["email"] = controller.GetSessionsValues(w, r, "email")
	m["active"] = controller.GetSessionsValues(w, r, "active")
	m["firstname"] = controller.GetSessionsValues(w, r, "firstname")
	m["prof"] = controller.GetSessionsValues(w, r, "prof")
	m["niveau"] = controller.GetSessionsValues(w, r, "niveau")
	m["numberFiles"] = model.COUNTFILES

	tpl.ExecuteTemplate(w, "layout", m)
}
