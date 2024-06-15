package types


type EditRequest struct {
	Nickname string `json:"nickname" validate:"required,min=8,max=30"`
	Birthday string `json:"birthday" validate:"required,len=10"`  // 1997-09-12
	Intro    string `json:"intro"    validate:"required,min=1,max=1000"`
	Avatar   string `json:"avatar"   validate:"required"`
}


type EditParam struct {
	UserID   int64
	Nickname string 
	Birthday int64 
	Intro    string 
	Avatar   string 
}


// tips 小贴士 修改密码、邮箱、手机，都要二次验证
type ResetPasswordRequest struct {
	Password        string   `json:"password"         validate:"required,min=8,max=32"`
	PasswordConfirm string   `json:"password_confirm" validate:"eqfield=Password"`
}

