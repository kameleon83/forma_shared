package model

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
)

// Config c
type Config struct {
	Directory string `gorm:"unique_index"`
}

// CreateConfig c
func (c *Config) CreateConfig() error {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	if c.Directory == "" {
		return db.AddError(errors.New("Le chemin d'accès est vide"))
	}

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
