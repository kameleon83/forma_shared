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

	// Get a session.
	session, err := store.Get(r, "formation_PHP")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	email := session.Values["email"]

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

	tpl.ExecuteTemplate(w, "layout", m)
}
