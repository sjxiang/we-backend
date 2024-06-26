package biz

import (
	"context"
	"we-backend/internal/service/mail"
	"we-backend/pkg/utils"
)

type OtpUsecase struct {
	otpRepo     OtpRepo 
	emailSvc    mail.EmailSender
}

func NewOtpUsecase(otpRepo OtpRepo, emailSvc mail.EmailSender) *OtpUsecase {
	return &OtpUsecase{
		otpRepo:  otpRepo,
		emailSvc: emailSvc,
	}
}

func (impl *OtpUsecase) SendOtp(ctx context.Context, biz string, phoneNumber string) error {
	code := utils.GenerateNum()

	if err := impl.otpRepo.Insert(ctx, biz, phoneNumber,code); err != nil {
		return err 
	}

	// 这前面成功了

	if err := impl.emailSvc.SendEmail(biz, code, []string{phoneNumber}, nil, nil, nil); err != nil {
		// 这里失败怎么办 
		// err 可能是超时的 err 你都不知道 发出了没
		// 让用户重试
		return err 
	}

	return nil 
}

func (impl *OtpUsecase) VerifyOtp(ctx context.Context, biz string, phoneNumber string, inputCode string) (valid bool, err error) {
	return impl.otpRepo.FindOne(ctx, biz, phoneNumber, inputCode)
}