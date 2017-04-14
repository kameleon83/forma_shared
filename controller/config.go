package controller

import (
	"bufio"
	"fmt"
	"forma_shared/model"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// DIRFILE folder list files
var DIRFILE = ""

// Config see config_directory
func Config(wg *sync.WaitGroup) {

	c := &model.Config{}
	c.Find()

	if c.Directory == "" || c.MailServer == "" || c.MailSender == "" || c.MailPassword == "" {
		c.Directory = directory()
		c.MailServer = mailServer()
		c.MailSender = mailSender()
		c.MailPassword = mailPassword()

		folderIsExist(c.Directory)

		err := c.Create()

		log.Println(err)

		if err == nil {
			wg.Done()
		}

		fmt.Println("le dossier de partage se situe : ", c.Directory)
		fmt.Println("le serveur mail est : ", c.MailServer)
		fmt.Println("l'expéditeur mail est : ", c.MailSender)
		fmt.Println("le mot de passe du serveur mail est : ", c.MailPassword)

	} else {
		DIRFILE = c.Directory
		wg.Done()
	}

}

func directory() string {

	var text string
	dirActual, _ := os.Getwd()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Le répertoire de partage n'est pas configuré !")
	fmt.Println()
	fmt.Println("Veux-tu en rentrer un un répertoire ? ex => /home/user/dlna/shared")
	fmt.Println()
	fmt.Println("Si oui, rentres le chemin du dossier à partager, puis valides avec la touche entrer.")
	fmt.Println()
	fmt.Println("Si non, appuies de suite sur entrer !! Et dans ce cas, le dossier de partage sera " + dirActual + "/partage")
	text, _ = reader.ReadString('\n')

	text = strings.TrimSpace(text)
	if text == "" {
		text = dirActual + string(os.PathSeparator) + "partage"
	}

	text = filepath.Clean(text) + string(os.PathSeparator)

	return text

}

func mailServer() string {

	var text string
	reader := bufio.NewReader(os.Stdin)

	for text == "" {

		fmt.Println("Le serveur mail n'est pas configurer")
		fmt.Println()
		fmt.Println("Rentres le server mail. (ex:smtp.gmail.com:587)")
		text, _ = reader.ReadString('\n')

		text = strings.TrimSpace(text)

	}
	return text

}

func mailSender() string {

	var text string
	reader := bufio.NewReader(os.Stdin)

	for text == "" {

		fmt.Println("L'expéditeur n'est pas configurer")
		fmt.Println()
		fmt.Println("Rentres le mail de l'expéditeur. (ex:toto@gmail.com)")
		text, _ = reader.ReadString('\n')

		text = strings.TrimSpace(text)

	}
	return text

}

func mailPassword() string {

	var text string
	reader := bufio.NewReader(os.Stdin)

	for text == "" {

		fmt.Println("Le mot de passe du serveur mail n'est pas configurer")
		fmt.Println()
		fmt.Println("Rentres le mot de passe du server mail. (ex:toto123456)")
		text, _ = reader.ReadString('\n')

		text = strings.TrimSpace(text)

	}
	return text

}
