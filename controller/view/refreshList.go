package controllerView

import (
	"forma_shared/controller"
	"net/http"
	"strconv"
)

// RefreshList refresh
func RefreshList(w http.ResponseWriter, r *http.Request) {
	ip, autorize := controller.CheckIP(w, r)
	controller.WriteLog("Annuaire : ", ip, controller.AfficheNom(ip), strconv.FormatBool(autorize))
	controller.ClientAutorize(w, r)

	FolderAndFiles := controller.ListFiles(controller.DIRFILE)

	controller.WriteJSON(FolderAndFiles)

	http.Redirect(w, r, "/", http.StatusFound)
}
