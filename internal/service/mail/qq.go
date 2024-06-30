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


func (sender *QQMailSender) proxy(
	subject string,  // 邮件主题
	content string,  // 邮件内容
	to []string,     // 收件人邮件地址
	cc []string,     // 抄送（群发可见，套瓷多尴尬）
	bcc []string,    // 抄送（群发不可见）
	attachFiles []string,  // 附件
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

func (sender *QQMailSender) SendEmail(subject, content, to string) error {
	return sender.proxy(subject, content, []string{to}, nil, nil, nil)
}