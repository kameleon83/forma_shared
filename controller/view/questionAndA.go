package controllerView

import (
	"forma_shared/controller"
	"forma_shared_dev/model"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// QuestionAndA q
func QuestionAndA(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	controller.GetSessionLogin(w, r)

	question := &model.Question{}
	answer := &model.Answer{}
	user := &model.User{}

	if r.Method == "POST" {
		question.Title = r.FormValue("title")
		question.Post = r.FormValue("post")
		user.Mail = r.FormValue("email")
		user.SearchUser()
		question.UserRefer = user.ID

		question.Create()

		http.Redirect(w, r, "/question&a", http.StatusFound)

	}

	funcMap := template.FuncMap{
		"title": strings.Title,
		"up":    strings.ToUpper,
		"searchEmailUser": func(i uint) string {
			user.ID = i
			log.Println(user)
			user.SearchUserByID()
			return user.Mail
		},
	}

	tpl := template.Must(template.New("QuestionAndA").Funcs(funcMap).ParseFiles("view/questionAndA.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml"))

	m := make(map[string]interface{})
	m["title"] = "Q&A"
	m["question"] = question.Search()
	m["answer"] = answer.Search()
	m["user"] = user.SearchUserByID()
	m["email"] = controller.GetSessionsValues(w, r, "email")
	m["active"] = controller.GetSessionsValues(w, r, "active")
	m["firstname"] = controller.GetSessionsValues(w, r, "firstname")
	m["prof"] = controller.GetSessionsValues(w, r, "prof")
	m["niveau"] = controller.GetSessionsValues(w, r, "niveau")
	// m["numberFiles"] = model.COUNTFILES
	// m["ip_name"] = controller.AfficheNom(ip)

	// fmt.Println(model.ReadUserJSON())
	tpl.ExecuteTemplate(w, "layout", m)
}
