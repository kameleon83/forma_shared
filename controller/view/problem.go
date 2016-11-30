package controllerView

import (
	"forma_shared/controller"
	"forma_shared/model"
	"html/template"
	"net/http"
	"strconv"

	"github.com/0xAX/notificator"
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
	m["title"] = "Checkpoint"
	m["ip_name"] = controller.AfficheNom(ip)
	m["followed"] = follow.ReadProblemJSON().Follow

	tpl.ExecuteTemplate(w, "layout", m)
}

// ResetFollow r
func ResetFollow(w http.ResponseWriter, r *http.Request) {
	model.ResetFollowJSON()
	http.Redirect(w, r, "/checkpoint", http.StatusFound)
}

var notify *notificator.Notificator

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

	var niveau_text string
	switch niveau {
	case 1:
		niveau_text = "Tout est bon pour moi"
	case 2:
		niveau_text = "En train de coder comme un porc!!!"
	case 3:
		niveau_text = "Help Me, Please!!!!"
	}
	notify = notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     "My test App",
	})

	notify.Push("title", niveau_text, "/home/user/icon.png", notificator.UR_CRITICAL)

	f.WriteProblemJSON()
}
