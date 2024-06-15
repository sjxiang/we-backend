package types

import "time"



type ProfileRequest struct {
	UserID int64 `json:"user_id"`
}

type Profile struct {
	Nickname          string    `json:"nickname"`
	Mobile            string    `json:"mobile"`
	Email             string    `json:"email"`
	Intro             string    `json:"intro"`
	Avatar            string    `json:"avatar"`
	CreatedAt         time.Time `json:"created_at"`
}

type ProfileResponse struct {
	Profile Profile
}

func ExportUserForFeedback(u *User) Profile {
	return Profile{
		Nickname:  u.Nickname,
		Mobile:    u.Mobile,
		Email:     u.Email,
		Intro:     u.Intro,
		Avatar:    u.Avatar,
		CreatedAt: u.CreatedAt,
	}
}
