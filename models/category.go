package models

import "gorm.io/gorm"

type Category struct { // Category Food
	Name string `gorm:"size:100"`

	CreatedBy int // dibuat oleh
	UpdatedBy int // diperbarui oleh
	gorm.Model
}
