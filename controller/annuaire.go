package controller

import (
	"forma_shared/lib"
	"forma_shared/model"
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Annuaire List user
func Annuaire(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	lib.GetSessionLogin(w, r)

	funcMap := template.FuncMap{
		"title": strings.Title,
		"up":    strings.ToUpper,
	}

	tpl := template.Must(template.New("Annuaire").Funcs(funcMap).ParseFiles("view/annuaire.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml"))

	vars := mux.Vars(r)
	sort, _ := vars["sort"]
	nameCol := vars["nameCol"]

	u := &model.User{}
	allEmail := u.SearchAllUsers(nameCol, sort)

	var mailto string
	var count int
	for _, user := range *allEmail {
		count += 1
		mailto += user.Mail
		if len(*allEmail) >= count {
			mailto += ","
		}
	}

	m := make(map[string]interface{})
	m["title"] = "Annuaire"
	m["users"] = allEmail
	m["mailto"] = mailto
	m["email"] = lib.GetSessionsValues(w, r, "email")
	m["active"] = lib.GetSessionsValues(w, r, "active")
	m["firstname"] = lib.GetSessionsValues(w, r, "firstname")
	m["prof"] = lib.GetSessionsValues(w, r, "prof")
	m["niveau"] = lib.GetSessionsValues(w, r, "niveau")
	m["numberFiles"] = model.COUNTFILES
	// m["ip_name"] = lib.AfficheNom(ip)

	// fmt.Println(model.ReadUserJSON())
	tpl.ExecuteTemplate(w, "layout", m)
}
