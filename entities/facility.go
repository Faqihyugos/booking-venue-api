package entities

import (
	"time"

	"gorm.io/gorm"
)

type Facility struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	IconName  string `gorm:"not null" json:"icon_name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  *gorm.DeletedAt
}
