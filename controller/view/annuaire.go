package controllerView

import (
	"fmt"
	"forma_shared/controller"
	"html/template"
	"net/http"
)

// Annuaire List user
func Annuaire(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("view/annuaire.gohtml")
	ip, autorize := controller.CheckIP(w, r)
	fmt.Println(ip, autorize)
	if !autorize {
		http.Redirect(w, r, "/not_access", 301)
	}
	m := make(map[string]interface{})
	m["title"] = "Annuaire"
	tpl.ExecuteTemplate(w, "annuaire.gohtml", m)
}
