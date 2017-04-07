package controllerView

import (
	"forma_shared/controller"
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

		u.Mail = controller.GetSessionsValues(w, r, "email").(string)
		if err := u.SearchUser().Error; err != nil {
			log.Println(err)
			flashes = controller.SetSessionsFlashes(w, r, "Utilisateur non trouvé")
		} else {
			if pass1 == pass2 {
				u.Password = controller.EncryptionPassword(pass1)
				controller.SetSessionsValues(w, r, "email", u.Mail)
				controller.SetSessionsValues(w, r, "firstname", u.Firstname)
				controller.SetSessionsValues(w, r, "lastname", u.Lastname)
				controller.SetSessionsValues(w, r, "ip", u.IP)
				controller.SetSessionsValues(w, r, "admin", u.Admin)
				controller.SetSessionsValues(w, r, "id", u.ID)
				if u.Active {
					controller.SetSessionsValues(w, r, "active", u.Active)
				}
				u.UpdateUser()
				// flashes = controller.SetSessionsFlashes(w, r, "L'enregistrement c'est bien passé. Vous allez recevoir un email avec un code de validation")
				// go controller.SendEmail(u.Mail)

				http.Redirect(w, r, "/", http.StatusFound)
			} else {
				flashes = controller.SetSessionsFlashes(w, r, "Les mots de passes ne sont pas pareils.")
			}

		}

	}
	m := make(map[string]interface{})
	m["title"] = "New Password"
	if len(flashes) > 0 {
		m["errors"] = flashes[0]
	}
	m["email"] = controller.GetSessionsValues(w, r, "email")
	m["active"] = controller.GetSessionsValues(w, r, "active")
	m["firstname"] = controller.GetSessionsValues(w, r, "firstname")
	m["prof"] = controller.GetSessionsValues(w, r, "prof")
	m["niveau"] = controller.GetSessionsValues(w, r, "niveau")
	m["numberFiles"] = model.COUNTFILES

	tpl.ExecuteTemplate(w, "layout", m)
}
