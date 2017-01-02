package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Answer a
type Answer struct {
	gorm.Model
	Post            string
	UserRefer       uint
	QuestionRefer   uint
	LikeAnswerRefer uint
}

// Search c
func (a *Answer) Search() *[]Answer {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	Answer := []Answer{}
	db.Find(&Answer)

	return &Answer
}

// Update c
func (a *Answer) Update() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	return db.Save(&a)
}

// Create c
func (a *Answer) Create() error {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

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
