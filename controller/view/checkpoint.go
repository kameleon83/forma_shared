package controllerView

import (
	"forma_shared/controller"
	"forma_shared/model"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Checkpoint s
func Checkpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tpl := template.Must(template.New("Checkpoint").ParseFiles("view/checkpoint.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml"))

	controller.GetSessionLogin(w, r)

	if controller.GetSessionsValues(w, r, "prof") != nil && !controller.GetSessionsValues(w, r, "prof").(bool) {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	user := &model.User{}
	CheckpointUserChange(w, r)
	m := make(map[string]interface{})
	m["title"] = "Checkpoint"
	m["user"] = user.CheckpointUsers()
	m["email"] = controller.GetSessionsValues(w, r, "email")
	m["active"] = controller.GetSessionsValues(w, r, "active")
	m["firstname"] = controller.GetSessionsValues(w, r, "firstname")
	m["prof"] = controller.GetSessionsValues(w, r, "prof")
	m["niveau"] = controller.GetSessionsValues(w, r, "niveau")
	m["numberFiles"] = model.COUNTFILES

	tpl.ExecuteTemplate(w, "layout", m)
}

// CheckpointReset r
func CheckpointReset(w http.ResponseWriter, r *http.Request) {
	u := &model.User{}
	u.CheckpointUsersReset()
	CheckpointChange(w, r, 0)
	controller.SetSessionsValues(w, r, "niveau", "")
	http.Redirect(w, r, "/checkpoint", http.StatusFound)
}

// ChangeLevelByName c
func ChangeLevelByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// name := vars["user"]
	niveau, _ := strconv.Atoi(vars["niveau"])
	CheckpointChange(w, r, niveau)
}

// CheckpointChange c
func CheckpointChange(w http.ResponseWriter, r *http.Request, niveau int) {
	u := model.User{}
	if controller.GetSessionsValues(w, r, "email") != nil {
		u.Mail = controller.GetSessionsValues(w, r, "email").(string)
	}
	u.SearchUser()
	u.Checkpoint = niveau
	u.UpdateUser()
	switch niveau {
	case 0:
		controller.SetSessionsValues(w, r, "niveau", "")
	case 1:
		controller.SetSessionsValues(w, r, "niveau", " done")
	case 2:
		controller.SetSessionsValues(w, r, "niveau", " in_progress")
	case 3:
		controller.SetSessionsValues(w, r, "niveau", " help")
	}
}

//CheckpointUserChange c
func CheckpointUserChange(w http.ResponseWriter, r *http.Request) {
	u := &model.User{}
	if controller.GetSessionsValues(w, r, "email") != nil {
		u.Mail = controller.GetSessionsValues(w, r, "email").(string)
	}
	u.SearchUser()
	CheckpointChange(w, r, u.Checkpoint)
}
