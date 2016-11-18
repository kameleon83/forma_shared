package controller

import (
	"forma_shared/model"
	"os"
	"path/filepath"
	"strings"
)

// DIRFILE folder list files
const DIRFILE = "/home/sam/go/src/forma_shared/"

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

// ListFiles index
func ListFiles(folder string) []model.MyFile {
	files := []model.MyFile{}
	filepath.Walk(folder, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			// fmt.Println(path)
			file := &model.MyFile{}
			file.Name = f.Name()
			// fmt.Println(path[len(DIRFILE+"files/"):])
			if f.Name() != strings.Split(path[len(DIRFILE):], "/")[1] {
				file.Folder = strings.Split(path[len(DIRFILE):], "/")[1]
				// fmt.Println(file.Folder)
			} else {
				file.Folder = ""
			}
			file.Size = float64(f.Size())
			file.Time = f.ModTime()
			file.Ext = filepath.Ext(f.Name())[1:]

			files = append(files, *file)

			// fmt.Printf("%s with %d bytes\n", path, f.Size())
		} else if f.IsDir() && f.Name() != "files" {
			if CheckFilesInFolder(path) {
				// fmt.Println(f.Size())
				folder := &model.MyFile{}

				folder.Folder = strings.Split(path[len(DIRFILE):], "/")[1]

				files = append(files, *folder)
			}
		}
		return nil
	})

	return files
}
