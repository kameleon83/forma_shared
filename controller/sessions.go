package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("samestunbeaugosse"))

// GetSessions g
func GetSessions(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "formation_PHP")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if session.Values["admin"] == true {
		log.Println("admin ok")
	} else {
		log.Println("admin not ok")
	}
	// Save it before we write to the response/return from the handler.
	session.Save(r, w)
}
