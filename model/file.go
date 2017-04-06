package model

import (
	"os"

	"github.com/jinzhu/gorm"
)

// File struct
type File struct {
	gorm.Model
	Name        string
	Size        float64
	Ext         string
	Path        string
	FolderRefer uint
}

// COUNTFILES c
var COUNTFILES int64

// CreateFile c
func (f *File) CreateFile() error {

	if err := f.SearchFile().Error; err != nil {
		db.Create(&f)
	} else {
		file, _ := os.Stat(f.Path)
		if file.ModTime().Unix() > f.UpdatedAt.Unix() {
			db.Save(&f)
		}
	}

	f.CountFiles()
	return db.Error
}

// CountFiles c
func (f *File) CountFiles() error {

	db.Model(&f).Count(&COUNTFILES)

	return db.Error
}

// SearchAllFiles read
func (f *File) SearchAllFiles() *[]File {

	file := []File{}
	folder := []Folder{}

	db.Model(&folder).Order("folder_refer").Order("created_at desc").Find(&file)
	// db.Model(&folder).Find(&file)
	return &file
}

// SearchFile s
func (f *File) SearchFile() *gorm.DB {

	return db.Where("name = ? AND folder_refer = ?", f.Name, f.FolderRefer).First(&f)
}

// // SearchFolder s
// func (f *File) SearchFolder() []string {
// 	files := f.SearchAllFiles()
// 	folders := []string{}
// 	var folder string
// 	for _, v := range *files {
// 		if folder != v.Folder {
// 			folders = append(folders, v.Folder)
// 		}
// 		folder = v.Folder
// 	}
// 	return folders
// }
