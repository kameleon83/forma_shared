package controller

import (
	"forma_shared/lib"
	"forma_shared/model"
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("samestunbeaugosse"))

// Register controller view inde.gohtml
func Register(w http.ResponseWriter, r *http.Request) {

	// lib.GetSessionLogin(w, r)

	tpl := template.Must(template.New("Register").ParseFiles("view/register.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml"))

	var flashes []interface{}

	if r.Method == "POST" {
		itsOk := true
		u := &model.User{}
		u.Lastname = strings.TrimSpace(r.FormValue("lastname"))
		u.Firstname = strings.TrimSpace(r.FormValue("firstname"))
		pass1 := strings.TrimSpace(r.FormValue("mdp"))
		pass2 := strings.TrimSpace(r.FormValue("mdp_verif"))

		if !u.ValidFirstname() {
			flashes = lib.SetSessionsFlashes(w, r, "Pas assez de caractères pour le prénom")
			itsOk = false
			// http.Redirect(w, r, "/register", http.StatusFound)
		}
		if !u.ValidLastname() {
			flashes = lib.SetSessionsFlashes(w, r, "Pas assez de caractères pour le nom")
			itsOk = false
			// http.Redirect(w, r, "/register", http.StatusFound)
		}

		if len(pass1) < 6 || len(pass2) < 6 {
			flashes = lib.SetSessionsFlashes(w, r, "Pas assez de caractères pour le mot de passe")
			itsOk = false
			// http.Redirect(w, r, "/register", http.StatusFound)
		}

		if lib.CheckEqualsPasswordString(pass1, pass2) != "" {
			u.Password = lib.CheckEqualsPasswordString(pass1, pass2)
			u.Function = strings.TrimSpace(r.FormValue("function"))
			u.Phone = strings.TrimSpace(r.FormValue("phone"))
			// if !u.ValidPhone() {
			// 	flashes = lib.SetSessionsFlashes(w, r, "Numéro de téléphone incorrect")
			// 	itsOk = false
			// 	http.Redirect(w, r, "/register", http.StatusFound)
			// } else {
			// 	u.Phone = "none"
			// }
			u.Mail = r.FormValue("email")
			if !u.ValidEmail() {
				flashes = lib.SetSessionsFlashes(w, r, "Email incorrect")
				itsOk = false
				// http.Redirect(w, r, "/register", http.StatusFound)
			}

			var sendMailForma = false
			if r.FormValue("prof") == "false" {
				u.Prof = false
			} else {
				sendMailForma = true
			}
			u.IP = lib.CheckIP(w, r)
			u.Admin = 0
			u.Active = false
			u.Checkpoint = 0

			// log.Println(u)

			if itsOk {
				if err := u.CreateUser(); err != nil {
					flashes = lib.SetSessionsFlashes(w, r, "Cet utilsateur existe déjà.")
					// http.Redirect(w, r, "/register", http.StatusFound)

				} else {
					flashes = lib.SetSessionsFlashes(w, r, "L'enregistrement c'est bien passé. Vous allez recevoir un email avec un code de validation")
					go lib.SendEmail(u.Mail)
					if sendMailForma {
						go lib.SendEmailFormer(u.Mail)
					}
					lib.SetSessionsValues(w, r, "email", u.Mail)
					http.Redirect(w, r, "/valid", http.StatusFound)

				}
			}
		} else {
			flashes = lib.SetSessionsFlashes(w, r, "les mdp ne sont pas pareil")
			// http.Redirect(w, r, "/register", http.StatusFound)
		}
	}

	m := make(map[string]interface{})
	m["title"] = "See or Download"
	if len(flashes) > 0 {
		m["errors"] = flashes[0]
	}
	tpl.ExecuteTemplate(w, "layout", m)
}
