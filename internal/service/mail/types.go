package mail


type EmailSender interface {
	SendEmail(subject, content, to string) error
}

