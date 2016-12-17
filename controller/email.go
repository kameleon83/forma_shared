package controller

import (
	"log"
	"net/smtp"
)

//SendEmail s
func SendEmail(email string) {
	hostname := "smtp.gmail.com"
	// Set up authentication information.
	auth := smtp.PlainAuth("", "samuel.michaux@gmail.com", "Mich_Sam_83600", hostname)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{email}
	msg := []byte("To: kameleon836@gmail.com\r\n" +
		"Subject: Valid your account\r\n" +
		"\r\n" +
		"Coucou comment vas-tu?.\r\n" +
		"\r\n" +
		"Alors tu es prêt à valider ton compte?.\r\n" +
		"\r\n" +
		"Dans ce cas il faut que tu copies ce code ci-dessous dans la page qui vient de s'ouvrir.\r\n" +
		"\r\n" +
		EncryptionEmail(email) + "\r\n" +
		"\r\n" +
		"A tout de suite. Biz :-P.\r\n")

	err := smtp.SendMail(hostname+":587", auth, "samuel.michaux@gmail.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

//SendEmailRescue s
func SendEmailRescue(email string) {
	hostname := "smtp.gmail.com"
	// Set up authentication information.
	auth := smtp.PlainAuth("", "samuel.michaux@gmail.com", "Mich_Sam_83600", hostname)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{email}
	msg := []byte("To: kameleon836@gmail.com\r\n" +
		"Subject: Rescue Password\r\n" +
		"\r\n" +
		"Coucou comment vas-tu?.\r\n" +
		"\r\n" +
		"Bah alors comme ça on ne se souvient plus du mot de passe de ton compte?.\r\n" +
		"\r\n" +
		"Heureusement pour toi j'ai prévu et tu n'as qu'à copier ce code ci-dessous dans la page qui vient de s'ouvrir.\r\n" +
		"\r\n" +
		EncryptionEmailRescue(email) + "\r\n" +
		"\r\n" +
		"A tout de suite. Biz :-P.\r\n")

	err := smtp.SendMail(hostname+":587", auth, "samuel.michaux@gmail.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
