package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Question q
type Question struct {
	gorm.Model
	Title     string `gorm:"unique_index"`
	Post      string
	Answers   []Answer `gorm:"ForeignKey:QuestionRefer"`
	UserRefer uint
}

// Search c
func (q *Question) Search() *[]Question {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	question := []Question{}

	db.Order("created_at desc").Find(&question)

	return &question
}

// Update c
func (q *Question) Update() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	return db.Save(&q)
}

// Create c
func (q *Question) Create() error {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	return db.Create(&q).Error
}
