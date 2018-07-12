package controller

import (
	"net/http"
	"html/template"
	"forma_shared/lib"
	"forma_shared/model"
	"strings"
	"time"
	"os"
	"strconv"
	"io"
)

func Pdf(w http.ResponseWriter, r *http.Request) {
	lib.GetSessionLogin(w, r)

	//config := new(model.Config)
	pwd, _ := os.Getwd()
	file := pwd + string(os.PathSeparator) +
		"pdf" +
		string(os.PathSeparator) +
		strings.ToLower(lib.GetSessionsValues(w, r, "firstname").(string)) +
		"_" +
		strings.ToLower(lib.GetSessionsValues(w, r, "lastname").(string)) +
		"_" +
		time.Now().Format("02-01-2006") +
		".pdf"

	Openfile, err := os.Open(file)
	defer Openfile.Close() //Close after function return
	if err != nil {
		//File not found, send 404
		http.Error(w, "File not found.", 404)
		return
	}

	//File is found, create and send the correct headers

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+
		lib.GetSessionsValues(w, r, "firstname").(string)+
		"_"+
		lib.GetSessionsValues(w, r, "lastname").(string)+
		"_" +
		time.Now().Format("02-01-2006") +
		".pdf")
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(w, Openfile) //'Copy' the file to the client
	return
}
func Stagiaire(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	lib.GetSessionLogin(w, r)

	tpl := template.Must(template.New("Stagiaire").ParseFiles("view/stagiairePdf.html", "view/layouts/header.gohtml", "view/layouts/footer.gohtml"))

	m := make(map[string]interface{})
	m["title"] = "Annuaire"
	m["email"] = lib.GetSessionsValues(w, r, "email")
	m["active"] = lib.GetSessionsValues(w, r, "active")
	m["firstname"] = strings.Title(strings.ToLower(lib.GetSessionsValues(w, r, "firstname").(string)))
	m["lastname"] = strings.Title(strings.ToLower(lib.GetSessionsValues(w, r, "lastname").(string)))
	m["prof"] = lib.GetSessionsValues(w, r, "prof")
	m["niveau"] = lib.GetSessionsValues(w, r, "niveau")
	user := &model.User{}
	user.Prof = true
	user.SearchProf()
	m["nameProf"] = strings.Title(user.Firstname + " " + user.Lastname)
	//lib.GeneratePdf()
	tpl.ExecuteTemplate(w, "layout", m)
}
func StagiairePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	lib.GetSessionLogin(w, r)

	sta := new(lib.StagiairePdf)

	header := new(lib.HeaderStruct)
	header.Stage = r.FormValue("stage")
	header.Module = r.FormValue("module")
	header.Date = datesToString(r.FormValue("date1"), r.FormValue("date2"))
	header.Formateur = r.FormValue("prof")
	header.Participant = strings.Title(r.FormValue("name"))

	sta.HeaderStruct = *header

	tab := new(lib.StagiairePdf).ColTable
	tab = append(tab,
		lib.FirstRow,
		tabResponse("Disponibilité du formateur", r.FormValue("tab1")),
		tabResponse("Structure du cours", r.FormValue("tab2")),
		tabResponse("Pertinence des excercices", r.FormValue("tab3")),
		tabResponse("Qualité des réponses apportées", r.FormValue("tab4")),
		tabResponse("Compétences techniques", r.FormValue("tab5")),
		tabResponse("Progression et rythme du stage", r.FormValue("tab6")),
		tabResponse("Adéquation à vos attentes professionnelles", r.FormValue("tab7")),
		tabResponse("Réponses à vos attentes personnelles", r.FormValue("tab8")),
		tabResponse("Matériels pédagogiques - salle", r.FormValue("tab9")),
		tabResponse("Qualité de l'accueil", r.FormValue("tab10")))

	sta.ColTable = tab

	sta.Comment = lib.Comment{Comment: r.FormValue("comment")}

	tabQ := lib.Question{
		Q1: questionTrueOrFalse(r.FormValue("q1")),
		Q2: questionTrueOrFalse(r.FormValue("q2")),
		Q3: questionTrueOrFalse(r.FormValue("q3")),
		Q4: questionTrueOrFalse(r.FormValue("q4")),
	}
	sta.Question = tabQ

	sta.NewForma = lib.NewForma{NewForma: r.FormValue("new-forma")}
	sta.Avis = lib.Avis{Avis: r.FormValue("avis")}
	sta.Indice = lib.Indice{Indice: r.FormValue("indice")}
	sta.Password = r.FormValue("password")

	sta.GeneratePdf()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func dateFormat(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

func datesToString(date1, date2 string) string {
	d1, _ := dateFormat(date1)
	d2, _ := dateFormat(date2)
	return "Du " + d1.Format("02-01-2006") + " au " + d2.Format("02-01-2006")
}

func tabResponse(title, notation string) lib.ColTable {
	n := lib.ColTable{}
	n.Title = title
	switch notation {
	case "veryGood":
		n.VeryGood = "X"
		break
	case "good":
		n.Good = "X"
		break
	case "prettyGood":
		n.PrettyGood = "X"
		break
	case "low":
		n.Low = "X"
		break
	}
	return n
}

func questionTrueOrFalse(str string) bool {
	if str == "yes" {
		return true
	}
	return false
}
