package controllerView

import (
	"forma_shared/controller"
	"html/template"
	"net/http"
)

// NotAccess error Ip
func NotAccess(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("view/not_access.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml")
	ip, _ := controller.CheckIP(w, r)
	m := make(map[string]interface{})
	m["ip"] = ip
	m["title"] = "Not Access Site"
	tpl.ExecuteTemplate(w, "content", m)
}
