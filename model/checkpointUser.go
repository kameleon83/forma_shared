package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

// CheckpointUsersReset Write file json File Model
func (u *User) CheckpointUsersReset() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	db.Model(&u).Update("checkpoint", 0)

	return db

}

// CheckpointUsers read
func (u *User) CheckpointUsers() []User {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	user := []User{}

	db.Find(&user)
	return user
}
