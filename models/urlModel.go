package models

import (
	"time"

	"gorm.io/gorm"
)

type Url struct {
	ID string `gorm:"primaryKey`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	RidrectTo string 
	Visit uint 
}