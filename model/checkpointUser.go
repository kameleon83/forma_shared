package model

import (
	"github.com/jinzhu/gorm"
)

// CheckpointUsersReset Write file json File Model
func (u *User) CheckpointUsersReset() *gorm.DB {

	db.Model(&u).Update("checkpoint", 0)

	return db

}

// CheckpointUsers read
func (u *User) CheckpointUsers() []User {

	user := []User{}

	db.Find(&user)

	tmp := []User{}
	for _, v := range user {
		if v.Checkpoint != 999 {
			tmp = append(tmp, v)
		}
	}

	return tmp
}
