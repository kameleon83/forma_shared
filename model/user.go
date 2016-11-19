package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"
)

// User model
type User struct {
	Firstname string
	Lastname  string
	Function  string
	Phone     string
	MailPro   string
	MailPers  string
	Prof      bool
	IP        string
}

// ByLastname s
type ByLastname []User

func (l ByLastname) Len() int { return len(l) }

func (l ByLastname) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

func (l ByLastname) Less(i, j int) bool { return l[i].Lastname < l[j].Lastname }

// ByFirstName s
type ByFirstName []User

func (f ByFirstName) Len() int { return len(f) }

func (f ByFirstName) Swap(i, j int) { f[i], f[j] = f[j], f[i] }

func (f ByFirstName) Less(i, j int) bool { return f[i].Firstname < f[j].Firstname }

// ReadUserJSON read
func ReadUserJSON(reverse bool, col string) []User {
	var file []User
	raw, err := ioutil.ReadFile(".user.json")
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal(raw, &file)

	if !reverse && col == "" {
		sort.Sort(ByLastname(file))
	}

	switch {
	case !reverse && col == "lastname":
		sort.Sort(ByLastname(file))
	case reverse && col == "lastname":
		sort.Sort(sort.Reverse(ByLastname(file)))
	case !reverse && col == "firstname":
		sort.Sort(ByFirstName(file))
	case reverse && col == "firstname":
		sort.Sort(sort.Reverse(ByFirstName(file)))
	}

	return file
}
