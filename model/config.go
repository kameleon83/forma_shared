package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Config c
type Config struct {
	Directory string
}

// CreateConfig c
func (c *Config) CreateConfig() error {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	return db.Create(c).Error
}

// SearchConfigDirectory s
func (c *Config) SearchConfigDirectory() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	return db.First(&c)
}

//UpdateConfig u
func (c *Config) UpdateConfig() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	return db.Save(&c)
}
