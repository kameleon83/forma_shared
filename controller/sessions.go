package controller

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("samestunbeaugosse"))

// SetSessionsFlashes g
func SetSessionsFlashes(w http.ResponseWriter, r *http.Request, message string) interface{} {
	// Get a session.
	session, err := store.Get(r, "formation_PHP")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return session.Flashes()
	}
	flashes := session.Flashes()

	if len(flashes) == 0 {
		session.AddFlash(message)
		session.Save(r, w)
	}
	return flashes
}

// GetSessionsValues g
func GetSessionsValues(w http.ResponseWriter, r *http.Request, value string) interface{} {
	// Get a session.
	session, _ := store.Get(r, "formation_PHP")

	if session.Values[value] != nil {
		return session.Values[value]
	}
	return nil
}

// SetSessionsValues g
func SetSessionsValues(w http.ResponseWriter, r *http.Request, name string, value interface{}) {
	session, _ := store.Get(r, "formation_PHP")

	session.Values[name] = value
	session.Save(r, w)
}

//GetSessionLogin g
func GetSessionLogin(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "formation_PHP")
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session.Values["email"] == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

	if session.Values["active"] == nil {
		http.Redirect(w, r, "/valid", http.StatusFound)
	}
	// Save it before we write to the response/return from the handler.
}

//GetSessionLogout g
func GetSessionLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "formation_PHP")
	session.Options = &sessions.Options{
		MaxAge: -1,
	}
	session.Save(r, w)
}
