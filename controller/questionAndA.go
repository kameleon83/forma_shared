package controller

import (
	"forma_shared/lib"
	"forma_shared/model"
	"html/template"
	"net/http"
	"strings"
	"time"
)

// QuestionAndA q
func QuestionAndA(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	lib.GetSessionLogin(w, r)

	question := &model.Question{}
	answer := &model.Answer{}
	user := &model.User{}

	const layout = "02-Jan-06 15h04"

	funcMap := template.FuncMap{
		"title": strings.Title,
		"up":    strings.ToUpper,
		"searchEmailUser": func(i uint) string {
			user.ID = i
			user.SearchUserByID()
			return user.Mail
		},
		"time_fr": func(t *time.Time) string {
			return t.Format(layout)
		},
	}

	tpl := template.Must(template.New("QuestionAndA").Funcs(funcMap).ParseFiles("view/questionAndA.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml"))

	if r.Method == "POST" {
		question.Title = r.FormValue("title")
		question.Post = r.FormValue("post")
		user.Mail = r.FormValue("email")
		user.SearchUser()
		question.UserRefer = user.ID

		question.Create()

		lib.CountQuestion = 1
		if len(question.Title) > 30 {
			lib.QuestionName = question.Title[0:30] + "..."
		} else {
			lib.QuestionName = question.Title
		}

		http.Redirect(w, r, "/question&a", http.StatusFound)

	}

	m := make(map[string]interface{})
	m["title"] = "Q&A"
	m["question"] = question.Search()
	m["answer"] = answer.Search()
	m["user"] = user.SearchUserByID()
	m["email"] = lib.GetSessionsValues(w, r, "email")
	m["active"] = lib.GetSessionsValues(w, r, "active")
	m["firstname"] = lib.GetSessionsValues(w, r, "firstname")
	m["prof"] = lib.GetSessionsValues(w, r, "prof")
	m["niveau"] = lib.GetSessionsValues(w, r, "niveau")
	// m["numberFiles"] = model.COUNTFILES
	// m["ip_name"] = lib.AfficheNom(ip)

	// fmt.Println(model.ReadUserJSON())
	tpl.ExecuteTemplate(w, "layout", m)
}
