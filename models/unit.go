package models

import "gorm.io/gorm"

type Unit struct {
	Id   int
	Name string `gorm:"size:100"`

	CreatedBy int // dibuat oleh
	UpdatedBy int // diperbarui oleh
	gorm.Model
}
