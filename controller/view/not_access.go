package controllerView

import (
	"forma_shared/controller"
	"html/template"
	"net/http"
	"strconv"
)

// NotAccess error Ip
func NotAccess(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("view/not_access.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml")

	ip, autorize := controller.CheckIP(w, r)
	controller.WriteLog("Not_Access : ", ip, controller.AfficheNom(ip), strconv.FormatBool(autorize))

	m := make(map[string]interface{})
	m["ip"] = ip
	m["title"] = "Not Access Site"
	tpl.ExecuteTemplate(w, "layout", m)
}
