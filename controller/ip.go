package controller

import (
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

// CheckIP Check Ip Client Connect
func CheckIP(w http.ResponseWriter, r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}
