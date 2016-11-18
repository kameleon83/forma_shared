package controllerView

import (
	"forma_shared/controller"
	"html/template"
	"net/http"
)

// NotAccess error Ip
func NotAccess(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("view/not_access.gohtml")
	ip, _ := controller.CheckIP(w, r)
	m := make(map[string]interface{})
	m["ip"] = ip
	tpl.ExecuteTemplate(w, "not_access.gohtml", m)
}
