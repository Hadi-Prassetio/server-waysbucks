package models

import "time"

type User struct {
	ID          int             `json:"id"`
	Fullname    string          `json:"fullname" gorm:"type: varchar(255)"`
	Email       string          `json:"email" gorm:"type: varchar(255)"`
	Password    string          `json:"-" gorm:"type: varchar(255)"`
	Status      string          `json:"status" gorm:"default:customer"`
	Profile     ProfileResponse `json:"profile"`
	Transaction []Transaction
	Cart        []Cart
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type UsersProfileResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
