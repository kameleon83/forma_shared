package model

import (
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

// Todo: Faire une table entre user and like afin de ne pouvoir checker qu'une seule fois

// Search c
func (q *Question) Search() *[]Question {

	question := []Question{}

	db.Order("created_at desc").Find(&question)

	return &question
}

// Update c
func (q *Question) Update() *gorm.DB {

	return db.Save(&q)
}

// Create c
func (q *Question) Create() error {

	return db.Create(&q).Error
}
