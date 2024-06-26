package types


type SentOtpRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`  // 实际是邮箱
}
