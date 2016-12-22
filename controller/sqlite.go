package controller

import (
	"forma_shared_dev/model"
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// // GetSqlite sqlite3
// func GetSqlite() {
//
// 	db, err := gorm.Open("sqlite3", "./gorm.db")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer db.Close()
// }

// CreateDatabase c
func CreateDatabase() {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	defer db.Close()
	db.AutoMigrate(&model.User{}, &model.Config{}, &model.File{}, &model.Folder{}, &model.Question{}, &model.Answer{})
	if err != nil {
		log.Println(err)
	}
}
