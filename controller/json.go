package controller

import (
	"encoding/json"
	"forma_shared_dev/model"
	"io/ioutil"
	"log"
	"os"
)

// WriteJSON Write file json File Model
func WriteJSON(file *model.FolderFile) {
	jsonFile := ".folders_files.json"
	if _, err := os.Stat(jsonFile); os.IsNotExist(err) {
		f, err := os.Create(jsonFile)
		log.Println(err)
		defer f.Close()
		WriteJSON(file)
	} else {
		b, _ := json.MarshalIndent(file, "", "\t")
		ioutil.WriteFile(jsonFile, b, 0644)
	}
}

// ReadJSON read
func ReadJSON(file *model.FolderFile) *model.FolderFile {
	raw, err := ioutil.ReadFile(".folders_files.json")
	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(raw, &file)

	// sort.Sort(model.ByDate(file))

	// sort.Sort(sort.Reverse(model.ByDate(file.File)))

	return file
}

// // ReadUserJSON read
// func ReadUserJSON(reverse bool, col string) []User {
// 	var file []User
// 	raw, err := ioutil.ReadFile(".user.json")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	json.Unmarshal(raw, &file)
//
// 	if !reverse && col == "" {
// 		sort.Sort(ByLastname(file))
// 	}
//
// 	switch {
// 	case !reverse && col == "lastname":
// 		sort.Sort(ByLastname(file))
// 	case reverse && col == "lastname":
// 		sort.Sort(sort.Reverse(ByLastname(file)))
// 	case !reverse && col == "firstname":
// 		sort.Sort(ByFirstName(file))
// 	case reverse && col == "firstname":
// 		sort.Sort(sort.Reverse(ByFirstName(file)))
// 	}
//
// 	return file
// }
