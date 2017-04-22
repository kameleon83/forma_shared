package controller

import (
	"fmt"
	"forma_shared/lib"
	"forma_shared/model"
	"html/template"
	"net/http"
	"strings"
)

// Valid v
func Valid(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("Valide").ParseFiles("view/valid.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml"))

	var flashes []interface{}

	email := lib.GetSessionsValues(w, r, "email")

	if r.Method == "POST" {
		key := strings.TrimSpace(r.FormValue("validKey"))
		if email != nil {
			if key == lib.EncryptionEmail(email.(string)) {
				u := model.User{}
				u.Mail = email.(string)
				u.SearchUser()

				u.Active = true
				u.UpdateUser()
				http.Redirect(w, r, "/login", http.StatusFound)
			} else {
				flashes = lib.SetSessionsFlashes(w, r, "La clef rentrée n'est pas la bonne")
				fmt.Println(flashes)
			}
		}
	}

	if email == nil {
		email = "Aucun email"
	}

	m := make(map[string]interface{})
	m["title"] = "Validation Compte"
	m["email"] = email.(string)
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
