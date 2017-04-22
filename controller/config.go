package controller

import (
	"forma_shared/helper"
	"forma_shared/lib"
	"forma_shared/model"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var flashes []interface{}

func Config(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		tpl := template.Must(template.New("Config").ParseFiles("view/config.gohtml", "view/layouts/header.gohtml", "view/layouts/footer.gohtml"))

		config := new(model.Config)
		if config.Count() > 0 {
			http.Redirect(w, r, "/", http.StatusFound)
		}

		// user := &model.User{}
		// CheckpointUserChange(w, r)
		m := make(map[string]interface{})
		m["title"] = "Config"
		m["pwd"] = directory()
		m["token"] = lib.TokenCreate()
		if len(flashes) > 0 {
			m["errors"] = flashes[0]
		}

		tpl.ExecuteTemplate(w, "layout", m)
	}
}

func ConfigPost(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		token := r.FormValue("token")
		pwd := r.FormValue("pwd")
		mailSender := r.FormValue("mailSender")
		mailServer := r.FormValue("mailServer")
		mailPassword := r.FormValue("password")

		var inputEmpty = func() bool {
			if token == "" || pwd == "" || mailSender == "" || mailServer == "" || mailPassword == "" {
				flashes = lib.SetSessionsFlashes(w, r, "Un ou plusieurs champs ne sont pas rempli ! :-(")
				return false
			}
			return true
		}

		if !inputEmpty() {
			http.Redirect(w, r, "/config", http.StatusFound)
			return
		}

		var smtpValidation = func() bool {
			if !helper.ValSmtp(mailServer) {
				flashes = lib.SetSessionsFlashes(w, r, "Vérifier bien que le serveur soit écrit de cette manière : smtp.gmail.com:587")
				return false
			}
			return true
		}

		if !smtpValidation() {
			http.Redirect(w, r, "/config", http.StatusFound)
			return
		}

		var emailValidation = func() bool {
			if !helper.ValEmail(mailSender) {
				flashes = lib.SetSessionsFlashes(w, r, "L'email est appremment éronné")
				return false
			}
			return true
		}

		if !emailValidation() {
			http.Redirect(w, r, "/config", http.StatusFound)
			return
		}

		if inputEmpty() && emailValidation() && smtpValidation() {
			err := configSendDb(token, pwd, mailSender, mailServer, mailPassword)

			if err == nil {
				go lib.SendEmailTest(mailSender)
				http.Redirect(w, r, "/register", http.StatusFound)
			}

		}
	}
}

func configSendDb(token, pwd, mailSender, mailServer, mailPassword string) error {
	c := &model.Config{}
	c.Find()

	var err error

	if c.Token == "" || c.Directory == "" || c.MailServer == "" || c.MailSender == "" || c.MailPassword == "" {
		c.Token = token
		c.Directory = pwd
		c.MailServer = mailServer
		c.MailSender = mailSender
		c.MailPassword = mailPassword

		lib.FolderIsExist(c.Directory)

		err = c.Create()

		log.Println(err)
		return err

	}
	return err
}

func directory() string {

	dirActual, _ := os.Getwd()

	text := dirActual + string(os.PathSeparator) + "partage"

	text = filepath.Clean(text) + string(os.PathSeparator)

	return text

}
