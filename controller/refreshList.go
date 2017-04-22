package controller

import (
	"forma_shared/lib"
	"forma_shared/model"
	"net/http"
)

// RefreshList refresh
func RefreshList(w http.ResponseWriter, r *http.Request) {

	lib.GetSessionLogin(w, r)

	// log.Println("DIRFILE : " + lib.DIRFILE)
	config := new(model.Config)
	lib.ListFiles(config.SendDirectory())

	http.Redirect(w, r, "/", http.StatusFound)
}
