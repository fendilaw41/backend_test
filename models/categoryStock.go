package models

import "gorm.io/gorm"

type CategoryStock struct {
	Name string `gorm:"size: 150"`

	CreatedBy int // dibuat oleh
	UpdatedBy int // diperbarui oleh
	gorm.Model
}
