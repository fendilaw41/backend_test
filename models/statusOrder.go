package models

import "gorm.io/gorm"

type StatusOrder struct {
	ID          uint   `gorm:"size:50"`
	Name        string `gorm:"size:50"`
	Description string `gorm:"size:200"`

	CreatedBy int // dibuat oleh
	UpdatedBy int // diperbarui oleh
	gorm.Model
}
