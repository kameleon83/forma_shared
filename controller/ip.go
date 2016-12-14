package controller

import (
	"forma_shared/model"
	"net"
	"net/http"
)

// Users ip user forma
// var Users = []model.User{
// 	{Firstname: "Sam", IP: "::1"},
// 	{Firstname: "Lucas", IP: "172.16.29.150"},
// 	{Firstname: "Francois", IP: "172.16.30.25"},
// 	{Firstname: "Laurent", IP: "172.16.29.193"},
// 	{Firstname: "Brice", IP: "172.16.27.30"},
// 	{Firstname: "Patrice", IP: "172.16.26.224"},
// 	{Firstname: "Andre", IP: "172.16.27.182"},
// 	{Firstname: "Seb", IP: "172.16.27.49"},
// 	{Firstname: "Pierre", IP: "172.16.30.3"},
// 	{Firstname: "Mouloud", IP: "172.16.30.66"},
// 	{Firstname: "Lea", IP: "172.16.27.113"},
// 	{Firstname: "Richard", IP: "172.16.24.172"},
// 	{Firstname: "Remi", IP: "172.16.26.224"},
// }

// AfficheNom Ip
func AfficheNom(ip string) string {
	user := model.ReadUserJSON(false, "lastname")
	for _, v := range user {
		if v.IP == ip {
			return v.Firstname
		}
	}
	return "Inconnu"
}

// CheckIP Check Ip Client Connect
func CheckIP(w http.ResponseWriter, r *http.Request) (string, bool) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	user := model.ReadUserJSON(false, "lastname")
	for _, v := range user {
		if v.IP == ip {
			autorized := true
			return ip, autorized
		}
	}
	return ip, false
}

// ClientAutorize get ip
func ClientAutorize(w http.ResponseWriter, r *http.Request) {
	// _, autorize := CheckIP(w, r)
	// // fmt.Println(ip, AfficheNom(ip), autorize)
	// if !autorize {
	// 	http.Redirect(w, r, "/not_access", 301)
	// }
}

// ClientAutorizeFormateur get ip
func ClientAutorizeFormateur(w http.ResponseWriter, r *http.Request) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	user := model.ReadUserJSON(false, "lastname")
	autorized := false
	for _, v := range user {
		if v.IP == ip && v.Prof == true {
			autorized = true
		} else if AfficheNom(ip) == "Samuel" {
			autorized = true
		}
	}
	if !autorized {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
