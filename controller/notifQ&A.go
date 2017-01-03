package controller

import (
	"net/http"
	"time"
)

var (
	// CountQuestion c
	CountQuestion int
	// CountAnswer c
	CountAnswer int

	//QuestionName s
	QuestionName string

	//AnswerName a
	AnswerName string
)

// CheckNewQuestionAndAnswer c
func CheckNewQuestionAndAnswer(w http.ResponseWriter, r *http.Request) {
	if CountQuestion != 0 {
		w.Write([]byte(QuestionName))
		timer := time.NewTimer(time.Second * 7)
		<-timer.C
		CountQuestion = 0
		timer.Stop()

	}

	if CountAnswer != 0 {
		w.Write([]byte(AnswerName))
		timer := time.NewTimer(time.Second * 7)
		<-timer.C
		CountAnswer = 0
		timer.Stop()
	}

}
