package model

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

// ConnDB function test connexion to sqlite
func ConnDB() {
	db, err = gorm.Open("sqlite3", "forma_shared.db")
	// db, err = gorm.Open("mysql", "root:@/mediatheque?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
	}
	// defer db.Close()
	db.LogMode(false)
	// db.Debug()
	db.DB()
	db.AutoMigrate(&User{}, &Config{}, &File{}, &Folder{}, &Question{}, &Answer{})
	if err != nil {
		log.Println(err)
	}
}
