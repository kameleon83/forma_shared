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
	log.Println(dirActual)

	if _, err := os.Stat("./config_directory"); os.IsNotExist(err) {

		reader := bufio.NewReader(os.Stdin)
		fmt.Println()
		fmt.Println("Le fichier indiquant le répertoire à partager n'est pas présent !!!!!!!!!!!!")
		fmt.Println()
		fmt.Println("Veux-tu en rentrer un ??? ex => /home/user/dlna/shared")
		fmt.Println()
		fmt.Println("Si oui, rentres le chemin du dossier à partager, puis valides avec la touche entrer.")
		fmt.Println()
		fmt.Println("Si non, appuies de suite sur entrer !! Et dans ce cas, le dossier de partage sera " + dirActual)
		fmt.Println()
		text, _ := reader.ReadString('\n')

		text = strings.TrimSpace(text)

		f, _ := os.Create("./config_directory")
		// Create a new writer.
		w := bufio.NewWriter(f)

		// Write a string to the file.
		if len(text) == 0 {
			w.WriteString("directory: " + filepath.Clean(dirActual) + string(filepath.Separator))
		} else {
			w.WriteString("directory: " + filepath.Clean(text) + string(filepath.Separator))
		}

		// Flush.
		w.Flush()

		ReadDir()
	}
	f, _ := os.Open("./config_directory")
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		line := scanner.Text()
		tabLine := strings.Split(line, ": ")
		switch {
		case tabLine[0] == "directory":
			DIRFILE = tabLine[1]
		default:
			DIRFILE = dirActual
		}
	}
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
func ListFiles(folder string) *model.FolderFile {
	files := &model.FolderFile{}
	dir := deleteEmpty(strings.Split(DIRFILE, "/"))

	filepath.Walk(folder, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			// fmt.Println(path)
			file := &model.File{}
			file.Name = f.Name()

			if f.Name() != strings.Split(path[len(DIRFILE)-1:], "/")[1] {
				file.Folder = strings.Split(path[len(DIRFILE)-1:], "/")[1]
				// fmt.Println(file.Folder)
			} else {
				file.Folder = ""
			}
			file.Size = float64(f.Size())
			file.Time = f.ModTime()
			file.Ext = filepath.Ext(f.Name())[1:]

			files.File = append(files.File, *file)

		} else if f.IsDir() && f.Name() != dir[len(dir)-1] {
			folder := &model.Folder{}
			folder.Name = strings.Split(path[len(DIRFILE)-1:], "/")[1]

			if CheckFilesInFolder(path) {
				folder.Empty = false
			} else {
				folder.Empty = true
			}
			files.Folder = append(files.Folder, *folder)
		}
		return nil
	})

	return files
}
