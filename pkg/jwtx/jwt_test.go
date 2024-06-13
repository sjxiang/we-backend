package jwtx

import "testing"


const SecretKey string = "nciwgh9wvt2"

func Test_generate_auth_to_token(t *testing.T) {
	
	var (
		id int64  = 3
		userAgent = "android"
	)

	token, err := GenerateAuth2Token(id, userAgent, SecretKey)
	if err != nil {
		t.Log(err)
	}

	claims, err := ExtractAuth2Token(token, SecretKey)
	if err != nil {
		t.Log(err)
	}

	t.Log(claims.ID)
	t.Log(claims.UserAgent)
}