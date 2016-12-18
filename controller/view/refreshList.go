package controllerView

import (
	"forma_shared/controller"
	"net/http"
)

// RefreshList refresh
func RefreshList(w http.ResponseWriter, r *http.Request) {

	controller.GetSessionLogin(w, r)

	// log.Println("DIRFILE : " + controller.DIRFILE)

	controller.ListFiles(controller.DIRFILE)

	http.Redirect(w, r, "/", http.StatusFound)
}
