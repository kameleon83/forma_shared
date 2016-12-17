package controllerView

import (
	"forma_shared/controller"
	"forma_shared/model"
	"html/template"
	"net/http"
	"strings"
)

// Valid v
func Valid(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("view/valid.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml")

	var flashes interface{}

	email := controller.GetSessionsValues(w, r, "email")

	if r.Method == "POST" {
		key := strings.TrimSpace(r.FormValue("validKey"))
		if email != nil {
			if key == controller.EncryptionEmail(email.(string)) {
				u := model.User{}
				u.Mail = email.(string)
				u.SearchUser()

				u.Active = true
				u.ActiveUser()
				http.Redirect(w, r, "/login", http.StatusFound)
			}
		}
	}

	m := make(map[string]interface{})
	m["title"] = "Validation Compte"
	m["email"] = email.(string)
	m["errors"] = flashes

	tpl.ExecuteTemplate(w, "layout", m)
}
