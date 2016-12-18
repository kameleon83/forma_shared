package model

import "github.com/jinzhu/gorm"

// Folder model
type Folder struct {
	gorm.Model
	Name        string
	Empty       bool
	FolderRefer uint
}
