package controllerView

import (
	"forma_shared/controller"
	"net/http"
)

// RefreshList refresh
func RefreshList(w http.ResponseWriter, r *http.Request) {
	FolderAndFiles := controller.ListFiles(controller.DIRFILE)

	controller.WriteJSON(".config.json", FolderAndFiles)

	http.Redirect(w, r, "/", http.StatusFound)
}
