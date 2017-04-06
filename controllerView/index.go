package controllerView

import (
	"forma_shared/controller"
	"forma_shared/model"
	"html/template"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Index controller view inde.gohtml
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	controller.GetSessionLogin(w, r)

	const layout = "Mon 02 Jan 2006 - 15:04:05"

	funcMap := template.FuncMap{
		"title": strings.Title,
		"upper": func(str string) string { return strings.ToUpper(str) },
		"lower": func(str string) string { return strings.ToLower(str) },
		"firstLetter": func(str string) string {
			return string(str[0])
		},
		"lastLetter": func(str string) string {
			return string(str[1:])
		},
		"time_fr": func(t *time.Time) string {
			return t.Format(layout)
		},
		"time_diff": func(t *time.Time) int {
			now := time.Now()
			now.Format(layout)
			diff := now.Sub(t.UTC())
			return int(diff.Minutes() / 10)
		},
		"exp": func(i float64) string {
			if i > math.Pow(10, 12) {
				return strconv.FormatFloat(i/math.Pow(10, 12), 'f', 2, 64) + " To"
			} else if i > math.Pow(10, 9) {
				return strconv.FormatFloat(i/math.Pow(10, 9), 'f', 2, 64) + " Go"
			} else if i > math.Pow(10, 6) {
				return strconv.FormatFloat(i/math.Pow(10, 6), 'f', 2, 64) + " Mo"
			} else if i > math.Pow(10, 3) {
				return strconv.FormatFloat(i/math.Pow(10, 3), 'f', 2, 64) + " ko"
			} else {
				return strconv.FormatFloat(i, 'f', -1, 64) + " octets"
			}
		},
	}

	tpl := template.Must(template.New("Index").Funcs(funcMap).ParseFiles("view/index.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml"))

	file := &model.File{}
	folder := &model.Folder{}

	CheckpointUserChange(w, r)

	m := make(map[string]interface{})
	m["title"] = "See or Download"
	m["files"] = file.SearchAllFiles()
	m["folder"] = folder.Search()
	m["email"] = controller.GetSessionsValues(w, r, "email")
	m["active"] = controller.GetSessionsValues(w, r, "active")
	m["firstname"] = controller.GetSessionsValues(w, r, "firstname")
	m["prof"] = controller.GetSessionsValues(w, r, "prof")
	m["admin"] = controller.GetSessionsValues(w, r, "admin")
	m["niveau"] = controller.GetSessionsValues(w, r, "niveau")
	// m["ip_name"] = controller.AfficheNom(ip)

	tpl.ExecuteTemplate(w, "layout", m)
}
