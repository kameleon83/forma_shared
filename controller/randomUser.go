package controller

import (
	"fmt"
	"forma_shared/model"
	"math/rand"
	"net/http"
	"time"
)

type users struct {
	user []model.User
}

var usersIsChecked *users

func RandomUser(w http.ResponseWriter, r *http.Request) {

	formatUser := searchUsers()
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	random := r1.Intn(len(formatUser.user))

	randUser := formatUser.user[random]
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
