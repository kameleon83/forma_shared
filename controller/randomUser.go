package controller

import (
	"fmt"
	"forma_shared/model"
	"math/rand"
	"net/http"
)

type users struct {
	user []model.User
}

var usersIsChecked *users

func RandomUser(w http.ResponseWriter, r *http.Request) {

	formatUser := searchUsers()
	rand := rand.Intn(len(formatUser.user))

	randUser := formatUser.user[rand]
	randUser.Checkpoint = 666
	randUser.UpdateUser()

	fmt.Fprint(w, randUser.Firstname)
}

func searchUsers() *users {
	u := &model.User{}
	us := u.CheckpointUsers()

	var idUsers = &users{}
	for _, v := range us {
		if v.Prof == false || v.Admin == 0 {
			if v.Firstname != "" {
				if v.Checkpoint != 666 || v.Checkpoint != 999 {
					idUsers.user = append(idUsers.user, v)
				}
			}
		}
	}
	return idUsers
}
