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

// {
// 	"File": [
// 		{
// 			"Name": "invite_LightRP.zip",
// 			"Folder": "lony",
// 			"Size": 94164,
// 			"Ext": "zip",
// 			"Time": "2016-11-19T22:17:03.3836775+01:00"
// 		},
// 		{
// 			"Name": "invite_invite_invite_invite_Samuel_invite_invite_invite_invite_CaptureÉcranDeepin20161119182340.png",
// 			"Folder": "lony",
// 			"Size": 183717,
// 			"Ext": "png",
// 			"Time": "2016-11-19T22:06:02.888856173+01:00"
// 		},
// 		{
// 			"Name": "CV-Samuel-Michaux-Dev.pdf",
// 			"Folder": "sam",
// 			"Size": 709429,
// 			"Ext": "pdf",
// 			"Time": "2016-11-19T14:46:01.582276486+01:00"
// 		},
// 		{
// 			"Name": "Sam_gtk-recordMyDesktop-crash.log",
// 			"Folder": "sam",
// 			"Size": 989,
// 			"Ext": "log",
// 			"Time": "2016-11-19T16:03:59.400404341+01:00"
// 		},
// 		{
// 			"Name": "appart lille fives.sh3d",
// 			"Folder": "sam",
// 			"Size": 42905,
// 			"Ext": "sh3d",
// 			"Time": "2016-11-19T14:46:01.582276486+01:00"
// 		},
// 		{
// 			"Name": "CV-Samuel-Michaux-Dev.pdf",
// 			"Folder": "sev",
// 			"Size": 709429,
// 			"Ext": "pdf",
// 			"Time": "2016-11-19T14:45:58.685604403+01:00"
// 		},
// 		{
// 			"Name": "Lony_Mario-mario-wallpaper-hd-games-1920x1080.jpg",
// 			"Folder": "sev",
// 			"Size": 378868,
// 			"Ext": "jpg",
// 			"Time": "2016-11-19T22:20:52.187530542+01:00"
// 		},
// 		{
// 			"Name": "Sam_dead.letter",
// 			"Folder": "sev",
// 			"Size": 1028,
// 			"Ext": "letter",
// 			"Time": "2016-11-19T16:16:00.771897085+01:00"
// 		},
// 		{
// 			"Name": "appart lille fives.sh3d",
// 			"Folder": "sev",
// 			"Size": 42905,
// 			"Ext": "sh3d",
// 			"Time": "2016-11-19T14:45:58.685604403+01:00"
// 		},
// 		{
// 			"Name": "invite_invite_invite_invite_Samuel_invite_invite_invite_invite_CaptureÉcranDeepin20161119182340.png",
// 			"Folder": "sev",
// 			"Size": 183717,
// 			"Ext": "png",
// 			"Time": "2016-11-19T22:07:24.959040202+01:00"
// 		},
// 		{
// 			"Name": "CV-Samuel-Michaux-Dev.pdf",
// 			"Folder": "test",
// 			"Size": 709429,
// 			"Ext": "pdf",
// 			"Time": "2016-11-19T14:45:54.008929013+01:00"
// 		},
// 		{
// 			"Name": "Sam_mail.html",
// 			"Folder": "test",
// 			"Size": 291,
// 			"Ext": "html",
// 			"Time": "2016-11-18T21:57:50.213196593+01:00"
// 		},
// 		{
// 			"Name": "appart lille fives.sh3d",
// 			"Folder": "test",
// 			"Size": 42905,
// 			"Ext": "sh3d",
// 			"Time": "2016-11-19T14:45:53.998928996+01:00"
// 		},
// 		{
// 			"Name": "invite_CaptureÉcranDeepin20161119182239.png",
// 			"Folder": "test",
// 			"Size": 130401,
// 			"Ext": "png",
// 			"Time": "2016-11-19T22:09:52.659371209+01:00"
// 		}
// 	],
// 	"Folder": [
// 		{
// 			"Name": "lony",
// 			"Empty": false
// 		},
// 		{
// 			"Name": "sam",
// 			"Empty": false
// 		},
// 		{
// 			"Name": "sev",
// 			"Empty": false
// 		},
// 		{
// 			"Name": "test",
// 			"Empty": false
// 		}
// 	]
// }
