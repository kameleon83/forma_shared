package model

import (
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
	db.Find(&Answer)

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
