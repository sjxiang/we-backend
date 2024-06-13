package jwtx

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)


type Auth2Claims struct {
	// 接口要实现方法太多，还是组合吧！
	jwt.RegisteredClaims   

	// 不要放敏感数据
	ID  int64 

	// 前端采集足够多的数据，做安全校验
	UserAgent string
}

/*
   
type Claims interface {
	GetExpirationTime() (*NumericDate, error)
	GetIssuedAt() (*NumericDate, error)
	GetNotBefore() (*NumericDate, error)
	GetIssuer() (string, error)
	GetSubject() (string, error)
	GetAudience() (ClaimStrings, error)
}

v5 需要实现这些

*/


// 生成 
func GenerateAuth2Token(identity int64, userAgent string, secretKey string) (string, error) {

	claims := &Auth2Claims{
		ID:          identity,
		UserAgent:   userAgent,

		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * time.Duration(168))},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	
	return tokenStr, nil
}


// 提取
func ExtractAuth2Token(accessToken, secretKey string) (*Auth2Claims, error) {
	
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	}

	token, err := jwt.ParseWithClaims(accessToken, &Auth2Claims{}, keyFunc)
	if err != nil {
		return nil, err  // 内容被篡改（签名和密钥对不上）
	}

	claims, ok := token.Claims.(*Auth2Claims)
	if !(ok && token.Valid) {
		return nil, err  // 类型断言、过期
	}

	return claims, nil
}

