package controllerView

import (
	"fmt"
	"forma_shared/controller"
	"forma_shared/model"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Annuaire List user
func Annuaire(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("view/annuaire.gohtml")
	ip, autorize := controller.CheckIP(w, r)
	fmt.Println(ip, autorize)
	// if !autorize {
	// 	http.Redirect(w, r, "/not_access", 301)
	// }

	vars := mux.Vars(r)
	reverse, _ := strconv.ParseBool(vars["reverse"])
	col := vars["col"]

	fmt.Println(reverse, col)

	m := make(map[string]interface{})
	m["title"] = "Annuaire"
	m["users"] = model.ReadUserJSON(reverse, col)
	m["ip_name"] = controller.AfficheNom(ip)

	// fmt.Println(model.ReadUserJSON())
	tpl.ExecuteTemplate(w, "annuaire.gohtml", m)
}
