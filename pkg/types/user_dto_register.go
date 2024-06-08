package types


type RegisterRequest struct {
	Email           string   `json:"email"            validate:"required,email"         binding:"required,email"`
	Password        string   `json:"password"         validate:"required,min=8,max=32"  binding:"required,min=8,max=48"`
	PasswordConfirm string   `json:"password_confirm" validate:"eqfield=Password"       binding:"eqfield=Password"`
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetPasswordConfirm() string {
	if x != nil {
		return x.PasswordConfirm
	}
	return ""
}


type RegisterResponse struct {
	UserID int64 `json:"user_id"`
}

func (x *RegisterResponse) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *RegisterResponse) ExportForFeedback() *RegisterResponse {
	return x
}
