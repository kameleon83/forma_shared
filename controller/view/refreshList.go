package controllerView

import (
	"forma_shared/controller"
	"net/http"
)

// RefreshList refresh
func RefreshList(w http.ResponseWriter, r *http.Request) {

	controller.GetSessionLogin(w, r)

	FolderAndFiles := controller.ListFiles(controller.DIRFILE)

	controller.WriteJSON(FolderAndFiles)

	http.Redirect(w, r, "/", http.StatusFound)
}
