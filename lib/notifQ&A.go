package lib

import (
	"net/http"
	"time"
)

var (
	CountQuestion int
	CountAnswer   int
	QuestionName  string
	AnswerName    string
)

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
