package controllerView

import (
	"forma_shared_dev/controller"
	"net/http"
)

// Logout l
func Logout(w http.ResponseWriter, r *http.Request) {
	controller.GetSessionLogout(w, r)
	http.Redirect(w, r, "/", http.StatusFound)
}
