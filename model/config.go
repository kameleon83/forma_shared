package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Config c
type Config struct {
	Token        string `gorm:"unique_index"`
	Directory    string `gorm:"unique_index"`
	MailServer   string `gorm:"unique_index"`
	MailPassword string `gorm:"unique_index"`
	MailSender   string `gorm:"unique_index"`
}

// CreateConfig c
func (c *Config) Create() error {

	if c.Directory == "" {
		return db.AddError(errors.New("Le chemin d'accès est vide"))
	}

	return db.Create(c).Error
}

// SearchConfigDirectory s
func (c *Config) Find() *gorm.DB {

	return db.First(&c)
}

// Count s
func (c *Config) Count() int {
	var count int
	db.Find(&c).Count(&count)
	return count
}

//UpdateConfig u
func (c *Config) Update() *gorm.DB {

	return db.Save(&c)
}

func (c *Config) SendDirectory() string {
	c.Find()
	return c.Directory
}
