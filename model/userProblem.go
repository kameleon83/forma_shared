package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
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
func ResetFollowJSON() {
	jsonFile := "./static/.user_followed.json"
	// user := &User{}
	if _, err := os.Stat(jsonFile); os.IsNotExist(err) {
		f, err := os.Create(jsonFile)
		log.Println(err)
		defer f.Close()
		b, _ := json.MarshalIndent(userjson, "", "\t")
		ioutil.WriteFile(jsonFile, b, 0644)
	} else {
		os.Remove(jsonFile)
		ResetFollowJSON()
		// str := `{user: ` + name + `, level: ` + strconv.Itoa(level) + `}`
		// t := json.Unmarshal([]byte(str), &u)
	}

}

// ChangeLevelByName c
// func (f *Followed) ChangeLevelByName(name string, level int) {
// 	follow := f.ReadProblemJSON()
// 	log.Println(name, level)
// 	log.Println(follow)
// }

// ReadProblemJSON read
func (f *Followed) ReadProblemJSON() *Followed {
	raw, err := ioutil.ReadFile("./static/.user_followed.json")
	if err != nil {
		log.Println(err)
	} else if os.IsNotExist(err) {
		ResetFollowJSON()
	}

	json.Unmarshal(raw, &f)

	return f
}

func (f *Followed) WriteProblemJSON() {
	jsonFile := "./static/.user_followed.json"

	b, _ := json.MarshalIndent(f, "", "\t")

	ioutil.WriteFile(jsonFile, b, 0644)
}

var userjson = &Followed{
	Follow: []OneFollowed{
		OneFollowed{"Samuel", 0},
		OneFollowed{"Léa", 0},
		OneFollowed{"Landa", 0},
		OneFollowed{"André", 0},
		OneFollowed{"Brice", 0},
		OneFollowed{"Sébastien", 0},
		OneFollowed{"Patrice", 0},
		OneFollowed{"François", 0},
		OneFollowed{"Laurent", 0},
		OneFollowed{"Mouloud", 0},
		OneFollowed{"Abdelmounaim", 0},
		OneFollowed{"Richard", 0},
		OneFollowed{"Pierre-Arnaud", 0},
		OneFollowed{"Lucas", 0},
	},
}

// var userjson = &Followed{
// 	Follow: []OneFollowed{Firstname: "Samuel", Level: 0},
// 	Follow: []OneFollowed{Firstname: "Léa", Level: 0},
// 	Follow: []OneFollowed{Firstname: "Linda", Level: 0},
// 	Follow: []OneFollowed{Firstname: "André", Level: 0},
// 	Follow: []OneFollowed{Firstname: "Brice", Level: 0},
// 	Follow: []OneFollowed{Firstname: "Sébastien", Level: 0},
// 	Follow: []OneFollowed{Firstname: "Patrice", Level: 0},
// 	Follow: []OneFollowed{Firstname: "François", Level: 0},
// 	Follow: []OneFollowed{Firstname: "Laurent", Level: 0},
// 	Follow: []OneFollowed{Firstname: "Mouloud", Level: 0},
// 	Follow: []OneFollowed{Firstname: "Abdelmounaim", Level: 0},
// 	Follow: []OneFollowed{Firstname: "Richard", Level: 0},
// 	Follow: []OneFollowed{Firstname: "Pierre-Arnaud", Level: 0},
// 	Follow: []OneFollowed{Firstname: "Lucas", Level: 0},
// }
