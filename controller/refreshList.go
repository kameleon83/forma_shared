package controller

import (
	"forma_shared/lib"
	"net/http"
)

// RefreshList refresh
func RefreshList(w http.ResponseWriter, r *http.Request) {

	lib.GetSessionLogin(w, r)

	// log.Println("DIRFILE : " + lib.DIRFILE)

	lib.ListFiles(lib.DIRFILE)

	http.Redirect(w, r, "/", http.StatusFound)
}
