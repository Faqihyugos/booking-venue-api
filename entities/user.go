package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	Fullname       string `gorm:"not null" json:"fullname" form:"fullname"`
	Username       string `gorm:"not null;unique" json:"username" form:"username"`
	Email          string `gorm:"not null;unique" json:"email" form:"email"`
	Password       string `gorm:"not null" json:"password" form:"password"`
	PhoneNumber    string `gorm:"not null" json:"phone_number" form:"phone_number"`
	Role           string `gorm:"not null" json:"role" form:"role"`
	Image          string `json:"image" form:"image"`
	Business_name  string `json:"business_name" form:"business_name"`
	Business_type  string `json:"business_type" form:"business_type"`
	Business_terms string `json:"business_terms" form:"business_terms"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeleteAt       *gorm.DeletedAt
}
