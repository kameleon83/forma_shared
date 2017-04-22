package controller

import (
	"forma_shared/lib"
	"forma_shared/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func QAndALike(w http.ResponseWriter, r *http.Request) {
	lib.GetSessionLogin(w, r)
	if r.Method == "POST" {
		vars := mux.Vars(r)
		postID, _ := strconv.ParseUint(vars["postID"], 10, 32)

		a := model.Answer{}
		a.ID = uint(postID)
		a.SearchByID()
		if vars["like"] == "good" {
			a.Good++
		} else if vars["like"] == "bad" {
			a.Bad++
		}
		a.Update()

		w.Write([]byte(`
                <div>` + strconv.Itoa(a.Good) + `</div>
                <img src="/images/like.png" alt="like" class="img-like" data-like="good">

                <div>` + strconv.Itoa(a.Bad) + `</div>
                <img src="/images/dont-like.png" alt="dont-like" class="img-dont-like" data-like="bad">
            `))
		// log.Println(postID, like)
	}
}
