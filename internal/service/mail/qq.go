package mail

import (
	"crypto/tls"
	"fmt"
	"we-backend/internal/conf"

	"gopkg.in/mail.v2"
)



type QQMailSender struct {
	name              string 
	fromEmailAddress  string 
	fromEmailPassword string 
}


func NewQQMailSender(cfg *conf.Config) EmailSender {
	return &QQMailSender{
		name:              cfg.QQmailSenderName,
		fromEmailAddress:  cfg.QQmailSenderAddress,
		fromEmailPassword: cfg.QQmailSenderPassword,
	}
}

func (sender *QQMailSender) SendEmail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string,
) error {

	m := mail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress))  
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	for _, f := range attachFiles {
		m.Attach(f)
	}

	d := mail.NewDialer("smtp.qq.com", 465, sender.fromEmailAddress, sender.fromEmailPassword)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	
	return d.DialAndSend(m)
}