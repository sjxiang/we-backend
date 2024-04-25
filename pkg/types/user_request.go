package types


type SignupRequest struct {
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password"  binding:"required,eqfield=ConfirmPassword"`
	ConfirmPassword string `json:"confirm_password" binding:"required,gte=8,lte=32"`
}


type LoginRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`  // e.g. validate:"required,len=6"
}

func NewLoginRequest() *LoginRequest {
	return &LoginRequest{}
}

func (req *LoginRequest) ExportEmailInString() string {
	return req.Email
}

func (req *LoginRequest) ExportPasswordInString() string {
	return req.Password
}


type (
	ProfileRequest struct {
		NickName string `json:"nickname" binding:"required,gte=8,lte=30"`
		Birthday string `json:"birthday" binding:"required,len=10"`  // e.g. 1997-09-12
		Intro    string `json:"intro" binding:"required,min=1,max=1000"`
		Avatar   string `json:"avatar" binding:"required"`
		Age      uint   `json:"age" binding:"required,numeric"`
	}

	EditInfoRequest struct {

	}
)
