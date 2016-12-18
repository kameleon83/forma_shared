package model

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

// File struct
type File struct {
	gorm.Model
	Name   string
	Folder string
	Size   float64
	Ext    string
	Path   string
}

// CreateFile c
func (f *File) CreateFile() error {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	if err := f.SearchFile().Error; err != nil {
		db.Create(&f)
	} else {
		file, _ := os.Stat(f.Path)
		if file.ModTime().Unix() > f.UpdatedAt.Unix() {
			db.Save(&f)
		}
	}

	return db.Error
}

// SearchAllFiles read
func (f *File) SearchAllFiles() *[]File {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	file := []File{}
	db.Order("folder").Order("updated_at desc").Find(&file)
	return &file
}

// SearchFile s
func (f *File) SearchFile() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	return db.Where("name = ? AND folder = ?", f.Name, f.Folder).First(&f)
}

// SearchFolder s
func (f *File) SearchFolder() []string {
	files := f.SearchAllFiles()
	folders := []string{}
	var folder string
	for _, v := range *files {
		if folder != v.Folder {
			folders = append(folders, v.Folder)
		}
		folder = v.Folder
	}
	return folders
}
