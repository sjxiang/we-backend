package biz

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"we-backend/internal/service/mail"
	"we-backend/internal/service/token"
	"we-backend/internal/types"
	"we-backend/pkg/faker"
	"we-backend/pkg/we"
)

type AuthUsecase struct {
	UserRepo     UserRepo
	CaptchaRepo  CaptchaRepo 
	TokenSvc     token.TokenService
	EmailSvc     mail.EmailSender
}

func NewAuthUsecase(ur UserRepo, cr CaptchaRepo, ts token.TokenService, es mail.EmailSender) *AuthUsecase {
	return &AuthUsecase{
		UserRepo:     ur, 
		CaptchaRepo:  cr,
		TokenSvc:     ts,
		EmailSvc:     es,
	}
}


func (au *AuthUsecase) Register(ctx context.Context, input types.RegisterInput) (types.RegisterResponse, error) {
	
	input.Sanitize()

	if err := input.Validate(); err != nil {
		return types.RegisterResponse{}, err
	}
	
	// check if email is already taken
	if _, err := au.UserRepo.GetByEmail(ctx, input.Email); !errors.Is(err, we.ErrNotFound) {
		return types.RegisterResponse{}, we.ErrEmailTaken
	}

	// hash the password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return types.RegisterResponse{}, fmt.Errorf("error hashing password: %v", err)
	}

	newUser := types.User{
		Email:    input.Email,
		Nickname: input.Username,
		Password: string(hashPassword),
	}

	// create the user
	user, err := au.UserRepo.Create(ctx, newUser)
	if err != nil {
		return types.RegisterResponse{}, fmt.Errorf("error creating user: %v", err)
	}
	
	rsp := types.RegisterResponse{
		User: mapUser(user),
	}

	return rsp, nil 
}

func (au *AuthUsecase) Login(ctx context.Context, input types.LoginInput) (types.LoginResponse, error) {
	input.Sanitize()

	if err := input.Validate(); err != nil {
		return types.LoginResponse{}, err
	}

	// 通过电子邮件，检索用户
	user, err := au.UserRepo.GetByEmail(ctx, input.Email)
	if err != nil {
		switch {
		case errors.Is(err, we.ErrNotFound):
			return types.LoginResponse{}, we.ErrInvalidCredentials
		default:
			return types.LoginResponse{}, err
		}
	}

	// 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return types.LoginResponse{}, we.ErrInvalidCredentials
	}

	// 签发 token
	accessToken, accessPayload, err := au.TokenSvc.CreateToken(user.ID, user.Email, time.Duration(144))
	if err != nil {
		return types.LoginResponse{}, err 
	}

	return types.LoginResponse{
		AccessToken: accessToken,
		ExpiresAt:   accessPayload.ExpiredAt,
	}, nil 
}

// 请求验证码
func (au *AuthUsecase) SentOtp(ctx context.Context, input types.SentOtpInput) error {

	code := faker.RandIntSpec()
	
	if err := au.CaptchaRepo.Store(ctx, input.Biz, input.PhoneNumber, code); err != nil {
		return err 
	}

	// 如果前面成功了

	// 如果后面失败了，怎么办？可能是超时，都不知道，发出了没有（让用户重试）

	if err := au.EmailSvc.SendEmail(input.Biz, code, input.PhoneNumber); err != nil {
		return err 
	}

	return nil 
}


// 验证码登录
func (au *AuthUsecase) VerifyOtp(ctx context.Context, input types.VerifyOtpInput) (valid bool, err error) {
	return au.CaptchaRepo.Verify(ctx, input.Biz, input.PhoneNumber, input.InputCode)
}