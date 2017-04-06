package controller

import (
	"forma_shared/model"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CheckFilesInFolder pour append select to upload
func CheckFilesInFolder(folder string) bool {
	count := 0
	filepath.Walk(folder, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			count++
		}
		return nil
	})
	if count > 0 {
		return true
	}
	return false

}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

// ListFiles index
func ListFiles(folder string) {
	// files := &model.FolderFile{}
	filepath.Walk(folder, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {

			// fmt.Println(path)
			file := &model.File{}
			fold := &model.Folder{}
			file.Name = f.Name()

			fold.Name = strings.Split(path[len(folder)-1:], string(os.PathSeparator))[1]
			fold.SearchWithName()

			fold.Empty = FilesorNotFolder(path)
			fold.Update()

			file.FolderRefer = fold.ID
			// fmt.Println(file.Folder)

			file.Size = float64(f.Size())
			file.Ext = filepath.Ext(f.Name())[1:]
			file.Path = path

			file.CreateFile()
		} else {
			if path != DIRFILE {
				fold := &model.Folder{}

				fold.Name = f.Name()

				if fold.SearchWithName() != nil {
					fold.Empty = FilesorNotFolder(path)
					fold.Create()
				} else {
					fold.Empty = FilesorNotFolder(path)
					fold.Update()
				}
			}

		}
		return nil
	})

	// return files
}

// FilesorNotFolder f
func FilesorNotFolder(path string) bool {
	var count int

	filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			count++
		}
		return nil
	})
	if count > 0 {
		return false
	}
	return true
}

func folderIsExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0666)
		log.Printf("Dossier %s Créé", path)
	}
	log.Printf("Dossier %s déjà existant", path)

}

// DeleteFile :
func DeleteFile(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idFile, _ := vars["fileID"]
	id, _ := strconv.Atoi(idFile)

	file := new(model.File)
	file.ID = uint(id)
	file.FindById()
	err := os.Remove(file.Path)
	if err == nil {
		file.Delete()
	}
}
