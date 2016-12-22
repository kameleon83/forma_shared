package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Folder model
type Folder struct {
	gorm.Model
	Name  string
	Empty bool
	Files []File `gorm:"ForeignKey:FolderRefer"`
}

// Search c
func (f *Folder) Search() *[]Folder {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	folder := []Folder{}
	db.Find(&folder)

	return &folder
}

// SearchWithName c
func (f *Folder) SearchWithName() error {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	db.Where("name = ?", f.Name).First(&f)

	return db.Error
}

// Create c
func (f *Folder) Create() error {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	db.Create(&f)

	return db.Error
}

// Update u
func (f *Folder) Update() error {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	db.Save(&f)

	return db.Error
}
