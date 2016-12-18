package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

// FolderFile group tab files and folder
type FolderFile struct {
	gorm.Model
	File   []File   `gorm:"ForeignKey:FileRefer"`
	Folder []Folder `gorm:"ForeignKey:FolderRefer"`
}

// CreateFolderFile c
func (f *FolderFile) CreateFolderFile() error {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	return db.Create(f).Error
}

// SearchAllFolderFiles read
func (f *FolderFile) SearchAllFolderFiles(col, sort string) *[]FolderFile {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	folderFile := []FolderFile{}
	db.Find(&folderFile)
	return &folderFile
}

// SearchFolderFile s
// func (f *FolderFile) SearchFolderFile() *gorm.DB {
// 	db, err := gorm.Open("sqlite3", "./gorm.db")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer db.Close()
//
// 	return db.Where("mail = ?", u.Mail).First(&u)
// }

//UpdateFolderFile u
func (f *FolderFile) UpdateFolderFile() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	return db.Save(f)
}
