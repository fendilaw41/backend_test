package models

import "gorm.io/gorm"

type Stock struct {
	ID              uint `gorm:"size:50"`
	FoodId          int
	Food            Food `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CategoryStockId int
	CategoryStock   CategoryStock `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Total           int

	CreatedBy int // dibuat oleh
	UpdatedBy int // diperbarui oleh
	gorm.Model
}
