package controller

import (
	"bufio"
	"forma_shared/model"
	"os"
	"path/filepath"
	"strings"
)

// DIRFILE folder list files
var DIRFILE = ReadDir()

// ReadDir see config_directory
func ReadDir() string {
	if _, err := os.Stat("./config_directory"); os.IsNotExist(err) {
		f, _ := os.Create("./config_directory")

		// Create a new writer.
		w := bufio.NewWriter(f)

		// Write a string to the file.
		w.WriteString("directory: /home/sam/go/src/forma_shared/files/")

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
			return tabLine[1]
		default:
			return "/home/sam/go/src/forma_shared/files"
		}
	}
	return ""
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
