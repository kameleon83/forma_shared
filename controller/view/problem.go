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
	ip, autorize := controller.CheckIP(w, r)
	controller.WriteLog("Problem : ", ip, controller.AfficheNom(ip), strconv.FormatBool(autorize))
	controller.ClientAutorizeFormateur(w, r)

	follow := &model.Followed{}

	m := make(map[string]interface{})
	m["title"] = "Followed"
	m["ip_name"] = controller.AfficheNom(ip)
	m["followed"] = follow.ReadProblemJSON().Follow

	tpl.ExecuteTemplate(w, "layout", m)
}

// ResetFollow r
func ResetFollow(w http.ResponseWriter, r *http.Request) {
	model.ResetFollowJSON()
	http.Redirect(w, r, "/followed", http.StatusFound)
}

// ChangeLevelByName c
func ChangeLevelByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["user"]
	niveau, _ := strconv.Atoi(vars["niveau"])

	f := &model.Followed{}
	p := f.ReadProblemJSON().Follow

	for i := range p {
		if p[i].Firstname == name {
			p[i].Level = niveau
		}
	}

	f.WriteProblemJSON()
}
