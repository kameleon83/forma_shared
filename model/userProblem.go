package model

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/jinzhu/gorm"
)

// Followed u
type Followed struct {
	Follow []OneFollowed
}

// OneFollowed u
type OneFollowed struct {
	Firstname string
	Level     int
}

// ResetFollowJSON Write file json File Model
func (u *User) ResetFollowJSON() *gorm.DB {
	// jsonFile := "./static/.user_followed.json"
	// // user := &User{}
	// if _, err := os.Stat(jsonFile); os.IsNotExist(err) {
	// 	f, err := os.Create(jsonFile)
	// 	log.Println(err)
	// 	defer f.Close()
	// 	b, _ := json.MarshalIndent(userjson, "", "\t")
	// 	ioutil.WriteFile(jsonFile, b, 0644)
	// } else {
	// 	os.Remove(jsonFile)
	// 	ResetFollowJSON()
	// 	// str := `{user: ` + name + `, level: ` + strconv.Itoa(level) + `}`
	// 	// t := json.Unmarshal([]byte(str), &u)
	// }

	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	db.Model(u).Update("checkpoint", 0)

	return db

}

// ChangeLevelByName c
// func (f *Followed) ChangeLevelByName(name string, level int) {
// 	follow := f.ReadProblemJSON()
// 	log.Println(name, level)
// 	log.Println(follow)
// }

// ReadProblemJSON read
func (u *User) ReadProblemJSON() []User {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	user := []User{}

	db.Find(&user)
	return user
}

// WriteProblemJSON w
func (f *Followed) WriteProblemJSON() {
	jsonFile := "./static/.user_followed.json"

	b, _ := json.MarshalIndent(f, "", "\t")

	ioutil.WriteFile(jsonFile, b, 0644)
}
