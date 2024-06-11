package types

type RegisterRequest struct {
	Email           string   `json:"email"            validate:"required,email"         binding:"required,email"`
	Password        string   `json:"password"         validate:"required,min=8,max=32"  binding:"required,min=8,max=48"`
	PasswordConfirm string   `json:"password_confirm" validate:"eqfield=Password"       binding:"eqfield=Password"`
}

type RegisterResponse struct {
	UserID int64 `json:"user_id"`
}
