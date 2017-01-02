package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

type LikeAnswer struct {
	gorm.Model
	Good   int
	Bad    int
	Answer *LikeAnswer `gorm:"ForeignKey:LikeAnswerRefer"`
	Users  []User      `gorm:"many2many:users_like;"`
}

// Search c
func (l *LikeAnswer) Search() *[]LikeAnswer {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	LikeAnswer := []LikeAnswer{}
	db.Order("good desc").Find(&LikeAnswer)

	return &LikeAnswer
}

// SearchWithUser c
func (l *LikeAnswer) SearchWithUser() *[]LikeAnswer {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	LikeAnswer := []LikeAnswer{}
	users := []User{}
	db.Model(&LikeAnswer).Related(&users, "Users")

	return &LikeAnswer
}

// Update c
func (l *LikeAnswer) Update() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	return db.Save(&l)
}

// Create c
func (l *LikeAnswer) Create() error {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	return db.Create(&l).Error
}

// SearchByID c
func (l *LikeAnswer) SearchByID() error {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	return db.First(&l).Error
}
