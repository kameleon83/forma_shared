package controllerView

import (
	"forma_shared/controller"
	"forma_shared/model"
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("samestunbeaugosse"))

// Register controller view inde.gohtml
func Register(w http.ResponseWriter, r *http.Request) {

	// controller.GetSessionLogin(w, r)

	tpl := template.Must(template.New("Partage").ParseFiles("view/register.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml"))

	// foldersAndFiles := controller.ReadJSON(&model.FolderFile{})

	// Get a session.
	session, err := store.Get(r, "formation_PHP")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var flashes = session.Flashes()

	if r.Method == "POST" {
		itsOk := true
		u := &model.User{}
		u.Lastname = strings.TrimSpace(r.FormValue("lastname"))
		u.Firstname = strings.TrimSpace(r.FormValue("firstname"))
		pass1 := strings.TrimSpace(r.FormValue("mdp"))
		pass2 := strings.TrimSpace(r.FormValue("mdp_verif"))

		if !u.ValidFirstname() {
			if len(flashes) == 0 {
				session.AddFlash("Pas assez de caractères pour le prénom")
				session.Save(r, w)
				itsOk = false
				http.Redirect(w, r, "/register", http.StatusFound)
			}
		}
		if !u.ValidLastname() {
			if len(flashes) == 0 {
				session.AddFlash("Pas assez de caractères pour le nom")
				session.Save(r, w)
				itsOk = false
				http.Redirect(w, r, "/register", http.StatusFound)
			}
		}

		if len(pass1) < 6 || len(pass2) < 6 {
			session.AddFlash("Pas assez de caractères pour le mot de passe")
			session.Save(r, w)
			itsOk = false
			http.Redirect(w, r, "/register", http.StatusFound)
		}

		if controller.CheckEqualsPasswordString(pass1, pass2) != "" {
			u.Password = controller.CheckEqualsPasswordString(pass1, pass2)
			u.Function = strings.TrimSpace(r.FormValue("function"))
			u.Phone = strings.TrimSpace(r.FormValue("phone"))
			if !u.ValidPhone() {
				session.AddFlash("Numéro de téléphone incorrect")
				session.Save(r, w)
				itsOk = false
				http.Redirect(w, r, "/register", http.StatusFound)
			}
			u.Mail = r.FormValue("email")
			if !u.ValidEmail() {
				session.AddFlash("Email incorrect")
				session.Save(r, w)
				itsOk = false
				http.Redirect(w, r, "/register", http.StatusFound)
			}
			if r.FormValue("prof") == "true" {
				u.Prof = true
			} else if r.FormValue("prof") == "false" {
				u.Prof = false
			}
			u.IP, _ = controller.CheckIP(w, r)
			u.Admin = 0
			u.Active = false

			// log.Println(u)

			if itsOk {
				if err := u.CreateUser(); err != nil {
					// if len(flashes) == 0 {
					session.AddFlash("Cet utilsateur existe déjà.")
					// }
					session.Save(r, w)
					http.Redirect(w, r, "/register", http.StatusFound)

				} else {
					if len(flashes) == 0 {
						session.AddFlash("L'enregistrement c'est bien passé. Vous allez recevoir un email avec un code de validation")
					}
					session.Save(r, w)
					go controller.SendEmail(u.Mail)
					session.Values["email"] = u.Mail
					session.Save(r, w)
					http.Redirect(w, r, "/valid", http.StatusFound)

				}
			}
		} else {
			if len(flashes) == 0 {
				session.AddFlash("les mdp ne sont pas pareil")
			}
			session.Save(r, w)
			http.Redirect(w, r, "/register", http.StatusFound)
		}
	}

	m := make(map[string]interface{})
	m["title"] = "See or Download"
	// m["errors"] = session.Flashes()
	m["errors"] = flashes
	// m["files"] = &foldersAndFiles.File
	// m["folder"] = &foldersAndFiles.Folder
	// m["ip_name"] = controller.AfficheNom(ip)
	tpl.ExecuteTemplate(w, "layout", m)
	session.Save(r, w)
}
