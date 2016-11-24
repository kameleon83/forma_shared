package controllerView

import (
	"encoding/json"
	"forma_shared/controller"
	"forma_shared/model"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ProblemFollowed s
func ProblemFollowed(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("view/problem.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml")
	ip, autorize := controller.CheckIP(w, r)
	controller.WriteLog("Problem : ", ip, controller.AfficheNom(ip), strconv.FormatBool(autorize))
	controller.ClientAutorize(w, r)

	follow := &model.Followed{}

	m := make(map[string]interface{})
	m["title"] = "Followed"
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

	for _, v := range f.ReadProblemJSON().Follow {
		if v.Firstname == name {
			log.Println(niveau)
			v.Level = niveau
		}
	}
	s, _ := json.MarshalIndent(f, "", "\n")
	log.Println(string(s))

	log.Println(name, niveau)
}
