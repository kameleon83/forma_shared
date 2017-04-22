package controller

import (
	"fmt"
	"forma_shared/lib"
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

	lib.GetSessionLogin(w, r)

	if lib.GetSessionsValues(w, r, "prof") != nil && !lib.GetSessionsValues(w, r, "prof").(bool) {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	user := &model.User{}
	CheckpointUserChange(w, r)
	m := make(map[string]interface{})
	m["title"] = "Checkpoint"
	m["user"] = user.CheckpointUsers()
	m["email"] = lib.GetSessionsValues(w, r, "email")
	m["active"] = lib.GetSessionsValues(w, r, "active")
	m["firstname"] = lib.GetSessionsValues(w, r, "firstname")
	m["prof"] = lib.GetSessionsValues(w, r, "prof")
	m["niveau"] = lib.GetSessionsValues(w, r, "niveau")
	m["numberFiles"] = model.COUNTFILES

	tpl.ExecuteTemplate(w, "layout", m)
}

// CheckpointUser s
func CheckpointUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		vars := mux.Vars(r)
		email := vars["email"]
		user := new(model.User)
		user.Mail = email
		user.SearchUser()

		fmt.Fprint(w, user.Checkpoint)
	}
}

// CheckpointReset r
func CheckpointReset(w http.ResponseWriter, r *http.Request) {
	u := &model.User{}
	u.CheckpointUsersReset()
	CheckpointChange(w, r, 0)
	lib.SetSessionsValues(w, r, "niveau", "")
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
	if lib.GetSessionsValues(w, r, "email") != nil {
		u.Mail = lib.GetSessionsValues(w, r, "email").(string)
	}
	u.SearchUser()
	u.Checkpoint = niveau
	u.UpdateUser()
	switch niveau {
	case 0:
		lib.SetSessionsValues(w, r, "niveau", "")
	case 1:
		lib.SetSessionsValues(w, r, "niveau", " done")
	case 2:
		lib.SetSessionsValues(w, r, "niveau", " in_progress")
	case 3:
		lib.SetSessionsValues(w, r, "niveau", " help")
	}
}

//CheckpointUserChange c
func CheckpointUserChange(w http.ResponseWriter, r *http.Request) {
	u := &model.User{}
	if lib.GetSessionsValues(w, r, "email") != nil {
		u.Mail = lib.GetSessionsValues(w, r, "email").(string)
	}
	u.SearchUser()
	CheckpointChange(w, r, u.Checkpoint)
}
