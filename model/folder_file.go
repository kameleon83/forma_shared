package model

import (
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
	return db.Create(f).Error
}

// SearchAllFolderFiles read
func (f *FolderFile) SearchAllFolderFiles(col, sort string) *[]FolderFile {
	folderFile := []FolderFile{}
	db.Find(&folderFile)
	return &folderFile
}

//UpdateFolderFile u
func (f *FolderFile) UpdateFolderFile() *gorm.DB {

	return db.Save(f)
}
