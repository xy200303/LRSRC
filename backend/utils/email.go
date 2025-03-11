package utils

import (
	"github.com/go-gomail/gomail"
	"mime"
)

func SendEmail(email string, title string, text string, smtpHost string, smtpUser string, smtpPWD string, smtpPort int, smtpUsername string) error {
	m := gomail.NewMessage()
	from := mime.QEncoding.Encode("UTF-8", smtpUsername) + " <" + smtpUser + ">"
	m.SetHeader("From", from)
	m.SetHeader("To", email)
	m.SetHeader("Subject", title)
	m.SetBody("text/html", text)
	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPWD)
	return d.DialAndSend(m)
}
