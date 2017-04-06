package model

import (
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
	folder := []Folder{}
	db.Find(&folder)

	return &folder
}

// SearchWithName c
func (f *Folder) SearchWithName() error {

	db.Where("name = ?", f.Name).First(&f)

	return db.Error
}

// Create c
func (f *Folder) Create() error {

	db.Create(&f)

	return db.Error
}

// Update u
func (f *Folder) Update() error {

	db.Save(&f)

	return db.Error
}
