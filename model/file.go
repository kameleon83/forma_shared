package model

import "time"

type MyFile struct {
	Name   string
	Folder string
	Size   float64
	Ext    string
	Time   time.Time
}
