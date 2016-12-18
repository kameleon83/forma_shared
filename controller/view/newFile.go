package controllerView

import (
	"forma_shared/model"
	"html/template"
	"net/http"
)

// NewFileCheck controller view inde.gohtml
func NewFileCheck(w http.ResponseWriter, r *http.Request) {

	tpl := template.Must(template.New("Countfiles").ParseFiles("view/countFiles.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml"))

	f := &model.File{}
	f.CountFiles()

	m := make(map[string]interface{})
	m["countfiles"] = model.COUNTFILES

	tpl.ExecuteTemplate(w, "layout", m)
}
