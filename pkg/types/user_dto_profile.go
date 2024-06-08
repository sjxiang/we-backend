package types


// type RegisterRequest struct {
// 	Email           string   `json:"email"            validate:"required,email"         binding:"required,email"`
// 	Password        string   `json:"password"         validate:"required,min=8,max=32"  binding:"required,min=8,max=48"`
// 	PasswordConfirm string   `json:"password_confirm" validate:"eqfield=Password"       binding:"eqfield=Password"`
// }

// func (x *RegisterRequest) GetEmail() string {
// 	if x != nil {
// 		return x.Email
// 	}
// 	return ""
// }

// func (x *RegisterRequest) GetPassword() string {
// 	if x != nil {
// 		return x.Password
// 	}
// 	return ""
// }

// func (x *RegisterRequest) GetPasswordConfirm() string {
// 	if x != nil {
// 		return x.PasswordConfirm
// 	}
// 	return ""
// }


// type RegisterResponse struct {
// 	UserID int64 `json:"user_id"`
// }

// func (x *RegisterResponse) GetUserID() int64 {
// 	if x != nil {
// 		return x.UserID
// 	}
// 	return 0
// }

// func (x *RegisterResponse) ExportForFeedback() *RegisterResponse {
// 	return x
// }


// type LoginRequest struct {
// 	Email    string `json:"email"    validate:"required,email"`
// 	Password string `json:"password" validate:"required,min=8,max=32"`  
// }

// func (req *LoginRequest) Validate() []string {
// 	errs := make([]string, 0)
	
// 	if err := ValidateString(req.Password, 8, 32); err != nil {
// 		errs = append(errs, err.Error())
// 	}

// 	if _, err := mail.ParseAddress(req.Email); err != nil {
// 		errs = append(errs, "不是有效的 email")
// 	}

// 	if err := validator.New().Var(req.Email, "required,email"); err != nil {
// 		errs = append(errs, "不是有效的 email")
// 	}

// 	allowed := false
// 	for _, domain := range emailDomainWhitelist {
// 		if strings.HasSuffix(req.Email, "@"+domain) {
// 			allowed = true
// 			break
// 		}
// 	}
// 	if !allowed {
// 		errs = append(errs, "邮箱地址的域名不在白名单中")
// 	}

// 	return errs
// }

// func (req *LoginRequest) ExportEmailInString() string {
// 	return req.Email
// }

// func (req *LoginRequest) ExportPasswordInString() string {
// 	return req.Password
// }


// func ValidateString(value string, minLength int, maxLength int) error {
// 	// n := len(value)  // 返回字节长度
// 	n := utf8.RuneCountInString(value)  // 返回字符长度

// 	if n < minLength || n > maxLength {
// 		return fmt.Errorf("字符长度必须在 %d-%d 之间", minLength, maxLength)
// 	}
// 	return nil
// }


// // 白名单
// var emailDomainWhitelist = []string{
// 	"gmail.com",
// 	"163.com",
// 	"126.com",
// 	"qq.com",
// 	"outlook.com",
// 	"hotmail.com",
// 	"yahoo.com",
// 	"foxmail.com",
// }

// type LoginResponse struct {

// }



// type (
// 	EditInfoRequest struct {
// 		NickName string `json:"nickname" binding:"required,gte=8,lte=30"`
// 		Birthday string `json:"birthday" binding:"required,len=10"`  // 1997-09-12
// 		Intro    string `json:"intro" binding:"required,min=1,max=1000"`
// 		Avatar   string `json:"avatar" binding:"required"`
// 	}

	
// 	// 注意，其它字段，尤其是密码、邮箱和手机
// 	// 修改都要通过别的手段
// 	// 邮箱和手机都要验证
// 	// 密码更加不用多说了
// 	ProfileRequest struct {
// 		NickName string `json:"nickname" binding:"required,gte=8,lte=30"`
// 		Birthday string `json:"birthday" binding:"required,len=10"`  // e.g. 1997-09-12
// 		Intro    string `json:"intro" binding:"required,min=1,max=1000"`
// 		Avatar   string `json:"avatar" binding:"required"`
// 		Age      uint   `json:"age" binding:"required,numeric"`
// 	}


// )
