package controller

import (
	"forma_shared/lib"
	"net/http"
)

// Logout l
func Logout(w http.ResponseWriter, r *http.Request) {
	lib.GetSessionLogout(w, r)
	http.Redirect(w, r, "/", http.StatusFound)
}
