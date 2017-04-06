package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Answer a
type Answer struct {
	gorm.Model
	Post          string
	Good          int
	Bad           int
	UserRefer     uint
	QuestionRefer uint
}

// Search c
func (a *Answer) Search() *[]Answer {
	Answer := []Answer{}
	db.Order("good desc").Find(&Answer)

	return &Answer
}

// Update c
func (a *Answer) Update() *gorm.DB {
	return db.Save(&a)
}

// Create c
func (a *Answer) Create() error {
	return db.Create(&a).Error
}

// SearchByID c
func (a *Answer) SearchByID() error {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	return db.First(&a).Error
}
