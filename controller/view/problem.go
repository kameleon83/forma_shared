package controllerView

import (
	"forma_shared/controller"
	"forma_shared/model"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ProblemFollowed s
func ProblemFollowed(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("view/problem.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml")

	controller.GetSessionLogin(w, r)

	if !controller.GetSessionsValues(w, r, "prof").(bool) {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	user := &model.User{}

	m := make(map[string]interface{})
	m["title"] = "Checkpoint"
	// m["ip_name"] = controller.AfficheNom(ip)
	m["user"] = user.ReadProblemJSON()
	m["email"] = controller.GetSessionsValues(w, r, "email")
	m["active"] = controller.GetSessionsValues(w, r, "active")
	m["firstname"] = controller.GetSessionsValues(w, r, "firstname")
	m["prof"] = controller.GetSessionsValues(w, r, "prof")

	tpl.ExecuteTemplate(w, "layout", m)
}

// ResetFollow r
func ResetFollow(w http.ResponseWriter, r *http.Request) {
	u := &model.User{}
	u.ResetFollowJSON()
	http.Redirect(w, r, "/checkpoint", http.StatusFound)
}

// ChangeLevelByName c
func ChangeLevelByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// name := vars["user"]
	niveau, _ := strconv.Atoi(vars["niveau"])

	u := model.User{}
	u.Mail = controller.GetSessionsValues(w, r, "email").(string)
	u.SearchUser()
	u.Checkpoint = niveau
	u.ActiveUser()
}
