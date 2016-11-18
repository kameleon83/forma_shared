package controller

import (
	"forma_shared/model"
	"net"
	"net/http"
)

// Users ip user forma
var Users = []model.User{
	{Name: "Sam", IP: "::1"},
	{Name: "Lucas", IP: "172.16.29.150"},
	{Name: "Francois", IP: "172.16.30.25"},
	{Name: "Laurent", IP: "172.16.29.193"},
	{Name: "Brice", IP: "172.16.27.30"},
	{Name: "Patrice", IP: "172.16.26.224"},
	{Name: "Andre", IP: "172.16.27.182"},
	{Name: "Seb", IP: "172.16.27.49"},
	{Name: "Pierre", IP: "172.16.30.3"},
	{Name: "Mouloud", IP: "172.16.30.66"},
	{Name: "Lea", IP: "172.16.27.113"},
	{Name: "Richard", IP: "172.16.24.172"},
	{Name: "Remi", IP: "172.16.26.224"},
}

// AfficheNom Ip
func AfficheNom(ip string) string {
	for _, v := range Users {
		if v.IP == ip {
			return v.Name
		}
	}
	return "Inconnu"
}

// CheckIP Check Ip Client Connect
func CheckIP(w http.ResponseWriter, r *http.Request) (string, bool) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	for _, v := range Users {
		if v.IP == ip {
			autorized := true
			return ip, autorized
		}
	}
	return ip, false
}
