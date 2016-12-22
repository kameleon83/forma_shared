package controllerView

import (
	"forma_shared_dev/controller"
	"forma_shared_dev/model"
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Annuaire List user
func Annuaire(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	controller.GetSessionLogin(w, r)

	funcMap := template.FuncMap{
		"title": strings.Title,
		"up":    strings.ToUpper,
	}

	tpl := template.Must(template.New("Annuaire").Funcs(funcMap).ParseFiles("view/annuaire.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml"))

	vars := mux.Vars(r)
	sort, _ := vars["sort"]
	nameCol := vars["nameCol"]

	u := &model.User{}
	m := make(map[string]interface{})
	m["title"] = "Annuaire"
	m["users"] = u.SearchAllUsers(nameCol, sort)
	m["email"] = controller.GetSessionsValues(w, r, "email")
	m["active"] = controller.GetSessionsValues(w, r, "active")
	m["firstname"] = controller.GetSessionsValues(w, r, "firstname")
	m["prof"] = controller.GetSessionsValues(w, r, "prof")
	m["niveau"] = controller.GetSessionsValues(w, r, "niveau")
	m["numberFiles"] = model.COUNTFILES
	// m["ip_name"] = controller.AfficheNom(ip)

	// fmt.Println(model.ReadUserJSON())
	tpl.ExecuteTemplate(w, "layout", m)
}
