package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

// CheckPointUsersReset Write file json File Model
func (u *User) CheckPointUsersReset() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	db.Model(u).Update("checkpoint", 0)

	return db

}

// CheckPointUsers read
func (u *User) CheckPointUsers() []User {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	user := []User{}

	db.Find(&user)
	return user
}
