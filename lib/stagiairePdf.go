package lib

import (
	"log"
	"github.com/signintech/gopdf"
	"time"
	"strconv"
	"strings"
	"os"
	wordwrap "github.com/mitchellh/go-wordwrap"
)

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
	Password     string
	pdf          *gopdf.GoPdf
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
	s.pdf = new(gopdf.GoPdf)
	s.pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	info := gopdf.PdfInfo{}
	info.Author = "Samuel Michaux"
	info.CreationDate = time.Now()
	info.Creator = info.Author
	info.Subject = "Evaluation des stagiaires"
	info.Title = info.Subject
	s.pdf.SetInfo(info)
	s.pdf.AddPage()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	s.font("ubuntu", dir+"/static/fonts/Ubuntu-Regular.ttf")
	s.fontPerso("ubuntu", 25)
	w, _ := s.pdf.MeasureTextWidth("EVALUATION DE STAGE")
	s.title(w)
	s.text(300, 60, "EVALUATION DE STAGE")
	s.fontPerso("ubuntu", 10)
	s.header()

	s.generateTable()

	s.text(30, s.pdf.GetY()+50, "Commentaires sur le stage")

	s.text(30, s.pdf.GetY()+10, "")
	s.writeParagraph(s.Comment.Comment, 110)
	s.text(30, s.pdf.GetY()+20, "Les informations concernant l'organisation et le contenu de ce stage ont-elles")
	s.allChecked(s.Question.Q1)
	s.text(30, s.pdf.GetY()+15, "été suffisantes ?")
	s.text(30, s.pdf.GetY()+25, "Estimez-vous nécessaire une formation appronfondie sur les thèmes abordés ?")
	s.allChecked(s.Question.Q2)

	s.text(30, s.pdf.GetY()+25, "Conseilleriez-vous ce stage à d'autres personnes ?")
	s.allChecked(s.Question.Q3)

	s.text(30, s.pdf.GetY()+25, "Envisagez-vous de nouvelles formations sur d'autres sujets ?")
	s.allChecked(s.Question.Q4)
	s.text(30, s.pdf.GetY()+15, "Si oui, lesquelles ?")
	s.text(30, s.pdf.GetY()+10, "")
	s.writeParagraph(s.NewForma.NewForma, 110)
	s.text(30, s.pdf.GetY()+20, "A votre avis, comment pourrait-on améliorer cette formation ?")
	s.text(30, s.pdf.GetY()+10, "")
	s.writeParagraph(s.Avis.Avis, 110)
	s.text(30, s.pdf.GetY()+20, "Indice de satisfaction global : ")
	s.text(170, s.pdf.GetY(), s.Indice.Indice)
	s.text(30, s.pdf.GetY()+10, "(1 = non satisfaisant - 5 = très bien)")
	s.fontPerso("ubuntu", 7)
	s.text(30, 800, "Nous vous remercions d'avoir pris le temps de répondre à ce questionnaire.")
	s.text(30, s.pdf.GetY()+10, "Toute l'équipe M2I FORMATION vous souhaite une bonne continuation dans vos activités.")

	s.fontPerso("ubuntu", 10)
	s.signature()
	s.text(355, 765, "Signature")
	s.font("signature", dir+"/static/fonts/Arty_Signature.otf")
	s.fontPerso("signature", 70)
	//// Signature
	//s.text(390, 810, strings.Title(s.HeaderStruct.Participant))

	s.NamePdf = "pdf/" + strings.Replace(strings.ToLower(s.HeaderStruct.Participant), " ", "_", -1) + "_" + time.Now().Format("02-01-2006")  + ".pdf"
	//pdf.WritePdf("pdf/" + strings.Replace(strings.ToLower(s.HeaderStruct.Participant), " ","_", -1) + "_" + strconv.Itoa(rand.Int()) + ".pdf")
	s.secutity()
	s.pdf.WritePdf(s.NamePdf)

}

func (s StagiairePdf) secutity() {
	protection := gopdf.PDFProtection{}
	protection.SetProtection(gopdf.PermissionsModify|gopdf.PermissionsCopy|gopdf.PermissionsPrint, []byte(s.Password), []byte("72683564"))
}

func (s StagiairePdf) signature() {
	s.pdf.SetStrokeColor(0, 0, 255)
	s.pdf.RectFromLowerLeftWithStyle(350, 820, 200, 70, "D")
}

func (s StagiairePdf) checked(x, y float64, str string, ok bool) {
	s.text(x, y, str)
	s.pdf.SetStrokeColor(0, 0, 255)
	if ok {
		s.pdf.RectFromLowerLeftWithStyle(x+25, y, 10, 10, "F")
	} else {
		s.pdf.RectFromLowerLeftWithStyle(x+25, y, 10, 10, "D")
	}

}

func (s StagiairePdf) allChecked(question bool) {
	s.checked(460, s.pdf.GetY(), "Oui", question)
	s.checked(515, s.pdf.GetY(), "Non", !question)
}

func (s StagiairePdf) generateTable() {
	s.pdf.SetY(s.pdf.GetY() + 5)
	for i := range s.ColTable {
		s.pdf.SetX(30)
		s.pdf.SetY(s.pdf.GetY() + 20)
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
	s.pdf.CellWithOption(&col1, " "+s.ColTable[i].Title, cell1)
	s.pdf.CellWithOption(&colOther, s.ColTable[i].VeryGood, cellOther)
	s.pdf.CellWithOption(&colOther, s.ColTable[i].Good, cellOther)
	s.pdf.CellWithOption(&colOther, s.ColTable[i].PrettyGood, cellOther)
	s.pdf.CellWithOption(&colOther, s.ColTable[i].Low, cellOther)
}

type HeaderStruct struct {
	Participant string
	Stage       string
	Module      string
	Date        string
	Formateur   string
}

func (s StagiairePdf) header() {
	c := &pdfConfig{}
	c.height = 130.00
	// A changer
	s.text(300, 100, "Lille, le "+timeFormat())

	s.text(30, c.height, "Participant : ")
	s.text(130, c.height, s.HeaderStruct.Participant)
	c.height = c.heightSpace()
	s.text(30, c.height, "Stage : ")
	s.text(130, c.height, s.HeaderStruct.Stage)
	c.height = c.heightSpace()
	if s.HeaderStruct.Module != ""{
		s.text(30, c.height, "Module : ")
		s.text(130, c.height, s.HeaderStruct.Module)
	}
	c.height = c.heightSpace()
	s.text(30, c.height, "Dates : ")
	s.text(130, c.height, s.HeaderStruct.Date)
	c.height = c.heightSpace()
	s.text(30, c.height, "Formateur : ")
	s.text(130, c.height, s.HeaderStruct.Formateur)
}

func (h *pdfConfig) heightSpace() float64 {
	return h.height + 20
}

func (s StagiairePdf) text(x float64, y float64, text string) {
	s.pdf.SetX(x)
	s.pdf.SetY(y)
	s.pdf.Text(text)
}
func (s StagiairePdf) fontPerso(font string, size int) {
	err := s.pdf.SetFont(font, "", size)
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

func (s StagiairePdf) title(w float64) {
	s.pdf.SetFillColor(191, 191, 191)
	s.pdf.RectFromLowerLeftWithStyle(290, 75, w+20, 50, "F")
	s.pdf.SetFillColor(0, 0, 0)
}

func (s StagiairePdf) font(name, path string) {
	err := s.pdf.AddTTFFont(name, path)
	logError(err)
}

func logError(err error) {
	if err != nil {
		log.Print(err.Error())
		return
	}
}

func word_wrap(text string, lineWidth int) string {
	words := strings.Fields(strings.TrimSpace(text))
	if len(words) == 0 {
		return text
	}
	wrapped := words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}

	return wrapped

}
const fontHeight = 11
const lineHeight = fontHeight * 1.2

func (s StagiairePdf) writeParagraph(text string, width uint) error {
	wrappedtext := wordwrap.WrapString(text, width)
	lines := strings.Split(wrappedtext, "\n")

	for _, line := range lines {
		s.pdf.SetX(40)
		err := (*s.pdf).Cell(nil, line)
		if err != nil {
			return err
		}

		(*s.pdf).Br(lineHeight)
	}

	return nil
}
