package model

import "time"

// File struct
type File struct {
	Name   string
	Folder string
	Size   float64
	Ext    string
	Time   time.Time
}

// ByLastname s
type ByDate []File

func (d ByDate) Len() int { return len(d) }

func (d ByDate) Swap(i, j int) { d[i], d[j] = d[j], d[i] }

func (d ByDate) Less(i, j int) bool { return d[i].Time.String() < d[j].Time.String() }

// // ReadUserJSON read
// func ReadFileJSON(reverse bool, col string) []File {
// 	var file []File
// 	raw, err := ioutil.ReadFile(".folders_files.json")
// 	if err != nil {
// 		log.Println(err)
// 	}
//
// 	json.Unmarshal(raw, &file)
//
// 	sort.Sort(ByDate(file))
//
// 	return file
// }
