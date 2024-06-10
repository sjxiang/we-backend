package utils

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	CTX_KEY_USERID = "user_id"
	CTX_KEY_EMAIL  = "user_email"
)


func GetUserIDFromContext(c *gin.Context) (int64, error) {
	uidRaw, exists := c.Get(CTX_KEY_USERID)
	if !exists {
		return 0, errors.New("kid missing from header")
	}

	uid, ok := uidRaw.(int64)
	if !ok {
		return 0, errors.New("kid malformed")
	}

	return uid, nil
}


func GetUserIDFromAuth(c *gin.Context) (uint, error) {
	// get request param
	userID, ok := c.Get(CTX_KEY_USERID)
	if !ok {
		// "auth token invalied, can not fetch user ID in it.")
		return 0, errors.New("input missing user_id field.")
	}

	userIDInt, okAssert := userID.(uint)
	if !okAssert {
		// "auth token invalied, user ID is not int type in it.")
		return 0, errors.New("input user_id in wrong format.")
	}

	return userIDInt, nil
}


func GetEmailFromAuth(c *gin.Context) (string, error) {
	// get request param
	email, ok := c.Get(CTX_KEY_EMAIL)
	if !ok {
		// "auth token invalied, can not fetch username in it.")
		return "", errors.New("input missing email field.")
	}

	emailString, okAssert := email.(string)
	if !okAssert {
		// "auth token invalied, username is not string type in it.")
		return "", errors.New("input email in wrong format.")
	}

	return emailString, nil
}

const (
	PARAM_AUTHORIZATION    = "Authorization"
	PARAM_USER_ID          = "userID"
)

func GetUserAuthTokenFromHeader(c *gin.Context) (string, error) {
	// fetch token
	rawToken := c.Request.Header[PARAM_AUTHORIZATION]
	if len(rawToken) != 1 {
		return "", errors.New("HTTP request header missing request token.")
	}
	var token string 
	token = rawToken[0]
	return token, nil
}


func GetUserAuthTokenFromHeaderX(c *gin.Context) (string, error) {
	
	authHeader := c.Request.Header.Get(PARAM_AUTHORIZATION)
	if authHeader == "" {
		return "", errors.New("HTTP request header missing request token.")
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("expected authorization header format: Bearer <token>")
	}

	var token string
	token = parts[1]

	return token, nil
}
