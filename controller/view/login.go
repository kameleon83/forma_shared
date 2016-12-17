package controllerView

import (
	"forma_shared/controller"
	"forma_shared/model"
	"html/template"
	"net/http"
)

// Login controller view inde.gohtml
func Login(w http.ResponseWriter, r *http.Request) {

	controller.GetSessions(w, r)

	tpl := template.Must(template.New("Partage").ParseFiles("view/login.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml"))

	// foldersAndFiles := controller.ReadJSON(&model.FolderFile{})

	// Get a session.
	session, err := store.Get(r, "formation_PHP")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	flashes := session.Flashes()

	if r.Method == "POST" {
		password := controller.EncryptionPassword(r.FormValue("password"))

		u := model.User{}
		u.Mail = r.FormValue("email")
		u.SearchUser()

		if u.Password == password {
			session.Values["email"] = u.Mail
			session.Values["firstname"] = u.Firstname
			session.Values["lastname"] = u.Lastname
			session.Values["ip"] = u.IP
			session.Values["admin"] = u.Admin
			session.Values["id"] = u.ID
			session.Save(r, w)
			if u.Active {
				session.Values["active"] = u.Active
				session.Save(r, w)
				http.Redirect(w, r, "/", http.StatusFound)
			} else {
				if flashes = session.Flashes(); len(flashes) > 0 {

				} else {
					session.AddFlash("Votre compte n'est pas actif!! Vérifier vos courriels")
					session.Save(r, w)
				}
				http.Redirect(w, r, "/valid", http.StatusFound)

			}

		} else {
			if flashes = session.Flashes(); len(flashes) > 0 {

			} else {
				session.AddFlash("Il y a une erreur d'authentification. Est-ce le bon mot de passe? Le bon utilisateur?")
				session.Save(r, w)
			}
			http.Redirect(w, r, "/login", http.StatusFound)
		}

	}
	m := make(map[string]interface{})
	m["title"] = "Login"
	if session.Values["email"] != nil {
		m["email"] = session.Values["email"]
	}
	m["errors"] = flashes
	if session.Values["active"] != nil {
		m["active"] = session.Values["active"]
	}
	// m["errors"] = flashes
	// m["files"] = &foldersAndFiles.File
	// m["folder"] = &foldersAndFiles.Folder
	// m["ip_name"] = controller.AfficheNom(ip)
	tpl.ExecuteTemplate(w, "layout", m)
}
