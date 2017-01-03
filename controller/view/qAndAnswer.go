package controllerView

import (
	"forma_shared/controller"
	"forma_shared/model"
	"net/http"
	"strconv"
)

// QAndAnswer q
func QAndAnswer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	controller.GetSessionLogin(w, r)

	if r.Method == "POST" {
		answer := &model.Answer{}
		user := &model.User{}
		answer.Post = r.FormValue("post")
		user.Mail = r.FormValue("email")
		questionID, _ := strconv.ParseUint(r.FormValue("question_id"), 10, 32)
		answer.QuestionRefer = uint(questionID)
		user.SearchUser()
		answer.UserRefer = user.ID

		answer.Create()

		controller.CountAnswer = 1
		if len(answer.Post) > 30 {
			controller.AnswerName = answer.Post[0:30] + "..."
		} else {
			controller.AnswerName = answer.Post
		}

		http.Redirect(w, r, "/question&a", http.StatusFound)

	}

}
