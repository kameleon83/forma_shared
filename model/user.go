package model

import (
	"log"
	"regexp"

	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Firstname  string
	Lastname   string
	Function   string
	Phone      string
	Password   string
	Mail       string `gorm:"unique_index"`
	Prof       bool
	IP         string
	Admin      int
	Active     bool
	Checkpoint int
	Questions  []Question `gorm:"ForeignKey:UserRefer"`
	Answers    []Answer   `gorm:"ForeignKey:UserRefer"`
}

// CreateUser c
func (u *User) CreateUser() error {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	return db.Create(u).Error
}

// SearchAllUsers read
func (u *User) SearchAllUsers(col, sort string) *[]User {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	user := []User{}
	if sort == "asc" {
		sort = ""
	}
	db.Order(col + " " + sort).Find(&user)
	return &user
}

// SearchUser s
func (u *User) SearchUser() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	return db.Where("mail = ?", u.Mail).First(&u)
}

// SearchUserByID s
func (u *User) SearchUserByID() *User {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	db.First(&u)
	return u
}

//UpdateUser u
func (u *User) UpdateUser() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	return db.Save(u)
}

// ValidFirstname v
func (u *User) ValidFirstname() bool {
	if len(u.Firstname) <= 2 {
		return false
	}
	return true
}

// ValidLastname v
func (u *User) ValidLastname() bool {
	if len(u.Lastname) <= 2 {
		return false
	}
	return true
}

// ValidEmail v
func (u *User) ValidEmail() bool {
	if m, _ := regexp.MatchString(`(?im)^([\w\.\-_]+)?\w+@[\w-_]+(\.\w+){1,}$`, u.Mail); !m {
		return false
	}
	return true
}

// ValidPass v
func (u *User) ValidPass() bool {
	if len(u.Password) <= 6 {
		return false
	}
	return true
}

// ValidPhone v
func (u *User) ValidPhone() bool {
	if m, _ := regexp.MatchString(`^0[1-68]([-. ]?[0-9]{2}){4}$`, u.Phone); !m {
		return false
	}
	return true
}
