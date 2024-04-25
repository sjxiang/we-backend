package request

type Signup struct {
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password"  binding:"required,eqfield=ConfirmPassword"`
	ConfirmPassword string `json:"confirm_password" binding:"required,gte=8,lte=32"`
}

type Login struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`  // e.g. validate:"required,len=6"
}

func NewLoginRequest() *Login {
	return &Login{}
}

func (login *Login) ExportEmailInString() string {
	return login.Email
}

func (login *Login) ExportPasswordInString() string {
	return login.Password
}


// 注意，其它字段，尤其是密码、邮箱和手机
// 修改都要通过别的手段
// 邮箱和手机都要验证
// 密码更加不用多说了
type Edit struct {
	NickName string `json:"nickname" binding:"required,gte=8,lte=30"`
	Birthday string `json:"birthday" binding:"required,len=10"`  // 1997-09-12
	Intro    string `json:"intro" binding:"required,min=1,max=1000"`
	Avatar   string `json:"avatar" binding:"required"`
	Age      uint   `json:"age" binding:"required,numeric"`
}
