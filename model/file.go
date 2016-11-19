package model

import "time"

// File struct
type File struct {
	Name   string
	Folder string
	Size   float64
	Ext    string
	Time   time.Time
}
