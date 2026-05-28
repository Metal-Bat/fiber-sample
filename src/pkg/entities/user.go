package entities

import "database/sql"

type Permission struct {
	BaseModel
	Name        string         `gorm:"index" json:"name"`
	Description sql.NullString `json:"description"`
	Users       []User         `gorm:"many2many:user_permissions;" json:"users"`
}

type User struct {
	BaseModel
	Mobile       string         `gorm:"unique,index" json:"mobile"`
	Email        string         `json:"email"`
	Password     string         `json:"password"`
	NationalCode sql.NullString `json:"national_code"`
	BirthDate    sql.NullTime   `json:"birthdate"`
	Permissions  []Permission   `gorm:"many2many:user_permissions;" json:"permissions"`
}
