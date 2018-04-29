package lib

import (
	"log"
	"github.com/signintech/gopdf"
	"time"
	"strconv"
	"strings"
	"os"
)

var pdf gopdf.GoPdf

type pdfConfig struct {
	height float64
}

type StagiairePdf struct {
	HeaderStruct HeaderStruct
	ColTable     []ColTable
	Comment      Comment
	Question     Question
	NewForma     NewForma
	Avis         Avis
	Indice       Indice
	NamePdf      string
}

//type tabColTable struct {
//	ColTable []ColTable
//}

type ColTable struct {
	Title      string
	VeryGood   string
	Good       string
	PrettyGood string
	Low        string
}

var titleTable = []string{
	"",
	"Disponibilité du formateur",
	"Structure du cours",
	"Pertinence des excercices",
	"Qualité des réponses apportées",
	"Compétences techniques",
	"Progression et rythme du stage",
	"Adéquation à vos attentes professionnelles",
	"Réponses à vos attentes personnelles",
	"Matériels pédagogiques - salle",
	"Qualité de l'accueil",
}

var FirstRow = ColTable{
	VeryGood:   "Très bien",
	Good:       "Bien",
	PrettyGood: "Assez Bien",
	Low:        "Faible",
}

type NewForma struct {
	NewForma string
}
type Avis struct {
	Avis string
}
type Indice struct {
	Indice string
}

type allQuestion struct {
	question []Question
}

type Comment struct {
	Comment string
}

type Question struct {
	Q1 bool
	Q2 bool
	Q3 bool
	Q4 bool
}

func (s StagiairePdf) GeneratePdf() {
	//pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	info := gopdf.PdfInfo{}
	info.Author = "Samuel Michaux"
	info.CreationDate = time.Now()
	info.Creator = info.Author
	info.Subject = "Evaluation des stagiaires"
	info.Title = info.Subject
	pdf.SetInfo(info)
	pdf.AddPage()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	font("ubuntu", dir+"/static/fonts/Ubuntu-Regular.ttf")
	fontPerso("ubuntu", 25)
	w, _ := pdf.MeasureTextWidth("EVALUATION DE STAGE")
	title(w)
	text(300, 60, "EVALUATION DE STAGE")
	fontPerso("ubuntu", 10)
	s.header()

	s.generateTable()

	text(30, pdf.GetY()+50, "Commentaires sur le stage")
	text(30, pdf.GetY()+15, s.Comment.Comment)
	text(30, pdf.GetY()+40, "Les informations concernant l'organisation et le contenu de ce stage ont-elles")
	allChecked(s.Question.Q1)
	text(30, pdf.GetY()+15, "été suffisantes ?")
	text(30, pdf.GetY()+25, "Estimez-vous nécessaire une formation appronfondie sur les thèmes abordés ?")
	allChecked(s.Question.Q2)

	text(30, pdf.GetY()+25, "Conseilleriez-vous ce stage à d'autres personnes ?")
	allChecked(s.Question.Q3)

	text(30, pdf.GetY()+25, "Envisagez-vous de nouvelles formations sur d'autres sujets ?")
	allChecked(s.Question.Q4)
	text(30, pdf.GetY()+15, "Si oui, lesquelles ?")
	text(30, pdf.GetY()+15, s.NewForma.NewForma)
	text(30, pdf.GetY()+40, "A votre avis, comment pourrait-on améliorer cette formation ?")
	text(30, pdf.GetY()+15, s.Avis.Avis)
	text(30, pdf.GetY()+40, "Indice de satisfaction global : ")
	text(170, pdf.GetY(), s.Indice.Indice)
	text(30, pdf.GetY()+10, "(1 = non satisfaisant - 5 = très bien)")
	fontPerso("ubuntu", 7)
	text(30, 800, "Nous vous remercions d'avoir pris le temps de répondre à ce questionnaire.")
	text(30, pdf.GetY()+10, "Toute l'équipe M2I FORMATION vous souhaite une bonne continuation dans vos activités.")

	fontPerso("ubuntu", 10)
	signature()
	text(355, 765, "Signature")
	font("signature", dir+"/static/fonts/Arty_Signature.otf")
	fontPerso("signature", 70)
	// Signature
	text(390, 810, strings.Title(s.HeaderStruct.Participant))

	s.NamePdf = "pdf/" + strings.Replace(strings.ToLower(s.HeaderStruct.Participant), " ", "_", -1) + ".pdf"
	//pdf.WritePdf("pdf/" + strings.Replace(strings.ToLower(s.HeaderStruct.Participant), " ","_", -1) + "_" + strconv.Itoa(rand.Int()) + ".pdf")
	pdf.WritePdf(s.NamePdf)

}

func signature() {
	pdf.SetStrokeColor(0, 0, 255)
	pdf.RectFromLowerLeftWithStyle(350, 820, 200, 70, "D")
}

func checked(x, y float64, str string, ok bool) {
	text(x, y, str)
	pdf.SetStrokeColor(0, 0, 255)
	if ok {
		pdf.RectFromLowerLeftWithStyle(x+25, y, 10, 10, "F")
	} else {
		pdf.RectFromLowerLeftWithStyle(x+25, y, 10, 10, "D")
	}

}

func allChecked(question bool) {
	checked(460, pdf.GetY(), "Oui", question)
	checked(515, pdf.GetY(), "Non", !question)
}

func (s StagiairePdf) generateTable() {
	pdf.SetY(pdf.GetY() + 5)
	for i := range s.ColTable {
		pdf.SetX(30)
		pdf.SetY(pdf.GetY() + 20)
		if i+1 == len(titleTable) {
			s.table(i, true)
		} else {
			s.table(i, false)
		}
	}
}

func (s StagiairePdf) table(i int, end bool) {
	col1 := gopdf.Rect{W: 250, H: 20}
	colOther := gopdf.Rect{W: 70, H: 20}

	cell1 := gopdf.CellOption{Border: 14, Align: 40}
	cellOther := gopdf.CellOption{Border: 14, Align: 48}
	if end {
		cell1.Border = 15
		cellOther.Border = 15
	}
	pdf.CellWithOption(&col1, " "+s.ColTable[i].Title, cell1)
	pdf.CellWithOption(&colOther, s.ColTable[i].VeryGood, cellOther)
	pdf.CellWithOption(&colOther, s.ColTable[i].Good, cellOther)
	pdf.CellWithOption(&colOther, s.ColTable[i].PrettyGood, cellOther)
	pdf.CellWithOption(&colOther, s.ColTable[i].Low, cellOther)
}

type HeaderStruct struct {
	Participant string
	Stage       string
	Module      string
	Date        string
	Formateur   string
}

func (h StagiairePdf) header() {
	c := &pdfConfig{}
	c.height = 130.00
	// A changer
	text(300, 100, "Lille, le "+timeFormat())

	text(30, c.height, "Participant : ")
	text(130, c.height, h.HeaderStruct.Participant)
	c.height = c.heightSpace()
	text(30, c.height, "Stage : ")
	text(130, c.height, h.HeaderStruct.Stage)
	c.height = c.heightSpace()
	text(30, c.height, "Module : ")
	text(130, c.height, h.HeaderStruct.Module)
	c.height = c.heightSpace()
	text(30, c.height, "Dates : ")
	text(130, c.height, h.HeaderStruct.Date)
	c.height = c.heightSpace()
	text(30, c.height, "Formateur : ")
	text(130, c.height, h.HeaderStruct.Formateur)
}

func (h *pdfConfig) heightSpace() float64 {
	return h.height + 20
}

func text(x float64, y float64, text string) {
	pdf.SetX(x)
	pdf.SetY(y)
	pdf.Text(text)
}
func fontPerso(font string, size int) {
	err := pdf.SetFont(font, "", size)
	logError(err)
}

var months = [...]string{
	"Janvier",
	"Février",
	"Mars",
	"Avril",
	"Mai",
	"Juin",
	"Juillet",
	"Août",
	"Septembre",
	"Octobre",
	"Novembre",
	"Décembre",
}

func monthFrench() string {
	i, _ := checkMonthActual()
	return months[i-1]
}

func checkMonthActual() (int, error) {
	return strconv.Atoi(time.Now().Format("1"))
}

func timeFormat() string {
	t := time.Now().Format("02-01-2006")
	month := monthFrench()
	str := strings.Split(t, "-")
	return str[0] + " " + month + " " + str[2]
}

func strConv(i int) string {
	return strconv.Itoa(i)
}

func title(w float64) {
	pdf.SetFillColor(191, 191, 191)
	pdf.RectFromLowerLeftWithStyle(290, 75, w+20, 50, "F")
	pdf.SetFillColor(0, 0, 0)
}

func font(name, path string) {
	err := pdf.AddTTFFont(name, path)
	logError(err)
}

func logError(err error) {
	if err != nil {
		log.Print(err.Error())
		return
	}
}
