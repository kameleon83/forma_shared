package lib

import (
	"forma_shared/model"
	"log"
	"net/smtp"
	"strings"
)

func auth() (smtp.Auth, string, string) {
	config := new(model.Config)
	config.Find()

	return smtp.PlainAuth("", config.MailSender, config.MailPassword, strings.Split(config.MailServer, ":")[0]), config.MailServer, config.MailSender
}

func send(recipient []string, msg []byte) {
	auth, server, sender := auth()
	err := smtp.SendMail(server, auth, sender, recipient, msg)
	if err != nil {
		log.Println(err)
	}
}

//SendEmail s
func SendEmail(email string) {
	to := []string{email}
	msg := []byte("To: " + email + "\r\n" +
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

	send(to, msg)
}

//SendEmailFormer s
func SendEmailFormer(email string) {
	config := new(model.Config)
	config.Find()
	to := []string{config.MailSender}
	msg := []byte("To: " + config.MailSender + "\r\n" +
		"Subject: Become a trainer\r\n" +
		"\r\n" +
		"Salut beau gosse.\r\n" +
		"\r\n" +
		email + " - veux devenir un formateur. Alors on le valides ou pas?\r\n" +
		"\r\n" +
		"A tout de suite. :-P.\r\n")

	send(to, msg)
}

//SendEmailRescue s
func SendEmailRescue(email string) {
	to := []string{email}
	msg := []byte("To: " + email + "\r\n" +
		"Subject: Rescue Password\r\n" +
		"\r\n" +
		"Coucou comment vas-tu?.\r\n" +
		"\r\n" +
		"Bah alors comme ça on ne se souvient plus du mot de passe de ton compte?.\r\n" +
		"\r\n" +
		"Heureusement pour toi j'ai prévu.Voici ton mot de passe provisoire\r\n" +
		"\r\n" +
		EncryptionEmailRescue(email) + "\r\n" +
		"\r\n" +
		"A tout de suite. :-P.\r\n")

	send(to, msg)
}
