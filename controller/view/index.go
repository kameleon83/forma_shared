package controllerView

import (
	"fmt"
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

	ip, autorize := controller.CheckIP(w, r)
	fmt.Println(ip, controller.AfficheNom(ip), autorize)
	// if !autorize {
	// 	http.Redirect(w, r, "/not_access", 301)
	// }

	const layout = "Mon 02 Jan 2006"

	funcMap := template.FuncMap{
		"title": strings.Title,
		"time_fr": func(t *time.Time) string {
			return t.Format(layout)
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

	tpl := template.Must(template.New("Partage").Funcs(funcMap).ParseFiles("view/index.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml"))

	foldersAndFiles := controller.ReadJSON(".config.json", &model.FolderFile{})

	m := make(map[string]interface{})
	m["title"] = "See or Download"
	m["files"] = &foldersAndFiles.File
	m["folder"] = &foldersAndFiles.Folder
	m["ip_name"] = controller.AfficheNom(ip)

	tpl.ExecuteTemplate(w, "layout", m)
}
