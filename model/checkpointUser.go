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
	return user
}
