package controllerView

import (
	"forma_shared/controller"
	"forma_shared/model"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Annuaire List user
func Annuaire(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("view/annuaire.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml")
	ip, autorize := controller.CheckIP(w, r)
	controller.WriteLog("Annuaire : ", ip, controller.AfficheNom(ip), strconv.FormatBool(autorize))
	// if !autorize {
	// 	http.Redirect(w, r, "/not_access", 301)
	// }

	vars := mux.Vars(r)
	reverse, _ := strconv.ParseBool(vars["reverse"])
	col := vars["col"]

	m := make(map[string]interface{})
	m["title"] = "Annuaire"
	m["users"] = model.ReadUserJSON(reverse, col)
	m["ip_name"] = controller.AfficheNom(ip)

	// fmt.Println(model.ReadUserJSON())
	tpl.ExecuteTemplate(w, "layout", m)
}
