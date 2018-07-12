package lib

import (
	"net"
	"net/http"
)

// CheckIP Check Ip Client Connect
func CheckIP(w http.ResponseWriter, r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}
