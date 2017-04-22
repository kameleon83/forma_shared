package lib

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

//WriteLog s
func WriteLog(str ...string) {
	fileName := ".forma_share.log"
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		err := ioutil.WriteFile(fileName, []byte{}, 0644)
		log.Println(err)
		WriteLog(str...)

	} else {
		file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer file.Close()

		t := time.Now()

		if _, err := file.WriteString(t.Format("Mon _2 Janv 2006 - 15:04:05") + " - " + strings.Join(str, " - ") + "\n"); err != nil {
			log.Println(err)
		}
	}
}
