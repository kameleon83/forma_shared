package controllerView

import (
	"fmt"
	"forma_shared/controller"
	"forma_shared/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// QAndALike q
func QAndALike(w http.ResponseWriter, r *http.Request) {
	controller.GetSessionLogin(w, r)

	// l := model.LikeAnswer{}

	if r.Method == "POST" {
		vars := mux.Vars(r)
		postID, _ := strconv.ParseUint(vars["postID"], 10, 32)

		a := model.Answer{}
		l := model.LikeAnswer{}
		// u := []model.User{}
		u := l.Users
		fmt.Println(u)
		user := model.User{}
		user.Mail = vars["user"]
		user.SearchUser()
		l.Users = append(l.Users, user)

		log.Println(postID, l.Users, vars["like"])

		a.ID = uint(postID)
		a.SearchByID()
		l.ID = a.LikeAnswerRefer
		l.SearchByID()

		fmt.Println(u)

		if vars["like"] == "good" {
			l.Good++
		} else if vars["like"] == "bad" {
			l.Bad++
		}
		// userId, _ := strconv.ParseUint(a.UserRefer, 10, 32)
		// users := l.Users

		log.Println(l.SearchWithUser())

		l.Update()

		w.Write([]byte(`
		        <div class="good" data-good="` + strconv.Itoa(l.Good) + `" data-user="` + user.Mail + `">` + strconv.Itoa(l.Good) + `</div>
		        <img src="/images/like.png" alt="like" class="img-like" data-like="good">

		        <div>` + strconv.Itoa(l.Bad) + `</div>
		        <img src="/images/dont-like.png" alt="dont-like" class="img-dont-like" data-like="bad">
		    `))
		// log.Println(postID, like)
	}
}
