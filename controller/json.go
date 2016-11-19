package controller

import (
	"encoding/json"
	"forma_shared/model"
	"io/ioutil"
	"log"
	"os"
)

// WriteJSON Write file json File Model
func WriteJSON(jsonFile string, file *model.FolderFile) {
	if _, err := os.Stat(jsonFile); os.IsNotExist(err) {
		f, err := os.Create(jsonFile)
		log.Println(err)
		defer f.Close()
		WriteJSON(jsonFile, file)
	} else {
		b, _ := json.MarshalIndent(file, "", "\t")
		ioutil.WriteFile(jsonFile, b, 0644)
	}
}

// ReadJSON read
func ReadJSON(jsonFile string, file *model.FolderFile) *model.FolderFile {
	raw, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(raw, &file)

	return file
}
