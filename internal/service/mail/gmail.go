package mail

import (
	"fmt"
	"net/smtp"
	"we-backend/internal/conf"

	"github.com/jordan-wright/email"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"      // 只负责认证，不传送邮件
	smtpServerAddress = "smtp.gmail.com:587"  // 负责实际传送邮件的SMTP服务器地址及端口
)


type GmailSender struct {
	name              string  // 发件人（泛指）
	fromEmailAddress  string  // 发件人邮箱地址
	fromEmailPassword string  // 发件人邮箱密码
}

func NewGmailSender(cfg *conf.Config) EmailSender {
	return &GmailSender{
		name:              cfg.GmailSenderName,
		fromEmailAddress:  cfg.GmailSenderAddress,
		fromEmailPassword: cfg.GmailSenderPassword,
	}
}

func (sender *GmailSender) SendEmail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string,
) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", 
					sender.name, 
					sender.fromEmailAddress)  // 发件人 we community <no-reply@gmail.com>
	e.Subject = subject  			          // 主题  
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	for _, f := range attachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("failed to attach file %s: %w", f, err)  // 附加文件失败
		}
	}

	smtpAuth := smtp.PlainAuth("", sender.fromEmailAddress, sender.fromEmailPassword, smtpAuthAddress)
	return e.Send(smtpServerAddress, smtpAuth)
}
