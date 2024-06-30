package mail



import (
	"fmt"
	"we-backend/internal/conf"
)



type LocalMailSender struct {
}


func NewLocalMailSender(cfg *conf.Config) EmailSender {
	return &LocalMailSender{}
}

func (sender *LocalMailSender) SendEmail(subject, content, to string) error {
	fmt.Println(content)
	return nil 
}