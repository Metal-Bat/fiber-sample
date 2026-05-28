package presenter

import (
	"time"
)

type UserDetail struct {
	ID           uint       `json:"id"`
	Mobile       string     `json:"mobile"`
	Email        string     `json:"email" binding:"required,email"`
	NationalCode *string    `json:"national_code"`
	BirthDate    *time.Time `json:"birthdate"`
}

type LoginInput struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginData struct {
	UserInfo   UserDetail `json:"user_info"`
	Token      string     `json:"token"`
	Expires_at int64      `json:"expires_at"`
}
