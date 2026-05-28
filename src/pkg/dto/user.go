package dto

import "time"

type UserInfo struct {
	ID           uint       `json:"id"`
	Mobile       string     `json:"mobile"`
	Email        string     `json:"email"`
	NationalCode *string    `json:"national_code"`
	BirthDate    *time.Time `json:"birthdate"`
	Permissions  []string   `json:"permissions"`
}

type LoginResult struct {
	UserInfo  UserInfo `json:"user_info"`
	Token     string   `json:"token"`
	ExpiresAt int64    `json:"expires_at"`
}

type UserDetail struct {
	ID           uint       `json:"id"`
	Mobile       string     `json:"mobile"`
	Email        string     `json:"email"`
	NationalCode *string    `json:"national_code"`
	BirthDate    *time.Time `json:"birthdate"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	Permissions  []string   `json:"permissions"`
}

type CreateUser struct {
	Mobile       string     `json:"mobile" validate:"required"`
	Email        string     `json:"email" validate:"required,email"`
	Password     string     `json:"password" validate:"required"`
	NationalCode *string    `json:"national_code,omitempty"`
	BirthDate    *time.Time `json:"birthdate,omitempty"`
}

type UpdateUser struct {
	Email        *string    `json:"email,omitempty"`
	Password     *string    `json:"password,omitempty"`
	NationalCode *string    `json:"national_code,omitempty"`
	BirthDate    *time.Time `json:"birthdate,omitempty"`
}
