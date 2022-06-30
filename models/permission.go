package models

import "gorm.io/gorm"

type Permission struct {
	ID          uint   `gorm:"size:50"`
	Name        string `gorm:"size:50"`
	Description string `gorm:"size:200"`
	gorm.Model
}
