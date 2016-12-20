package controller

import (
	"bufio"
	"fmt"
	"forma_shared/model"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// DIRFILE folder list files
var DIRFILE = ""

// ReadDir see config_directory
func ReadDir() {
	dirActual, _ := os.Getwd()
	// log.Println(dirActual)

	c := &model.Config{}
	c.SearchConfigDirectory()

	if c.Directory == "" {
		// c.CreateConfig()

		reader := bufio.NewReader(os.Stdin)
		fmt.Println()
		fmt.Println("Le répertoire de partage n'est pas configuré !")
		fmt.Println()
		fmt.Println("Veux-tu en rentrer un un répertoire ? ex => /home/user/dlna/shared")
		fmt.Println()
		fmt.Println("Si oui, rentres le chemin du dossier à partager, puis valides avec la touche entrer.")
		fmt.Println()
		fmt.Println("Si non, appuies de suite sur entrer !! Et dans ce cas, le dossier de partage sera " + dirActual)
		fmt.Println()
		text, _ := reader.ReadString('\n')

		text = strings.TrimSpace(text)

		c.Directory = filepath.Clean(text) + string(filepath.Separator)

		err := c.CreateConfig()

		log.Println(err)

	}

	DIRFILE = c.Directory

}

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
			file.Name = f.Name()

			if f.Name() != strings.Split(path[len(folder)-1:], "/")[1] {
				file.Folder = strings.Split(path[len(folder)-1:], "/")[1]
				// fmt.Println(file.Folder)
			} else {
				file.Folder = ""
			}
			file.Size = float64(f.Size())
			file.Ext = filepath.Ext(f.Name())[1:]
			file.Path = path

			file.CreateFile()
		}
		return nil
	})

	// return files
}
