package models

import "gorm.io/gorm"

type Cart struct {
	ID        uint `gorm:"size:50"`
	FoodId    int
	Food      Food `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; association_foreignkey:FoodId"`
	Qty       int
	CreatedBy int // dibuat oleh
	UpdatedBy int // diperbarui oleh
	gorm.Model
}
