package models

import "gorm.io/gorm"

type Tax struct {
	ID   int
	Name string `gorm:"size:100"`
	Qty  int

	CreatedBy int // dibuat oleh
	UpdatedBy int // diperbarui oleh
	gorm.Model
}
