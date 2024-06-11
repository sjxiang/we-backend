package otp


// sms
type OtpService interface {

}

type wrapper struct {

}

func NewOtpService() OtpService {
	return &wrapper{}
}

func (w *wrapper) Send() error {
	return nil 
}