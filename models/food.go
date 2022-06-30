package models

import (
	"gorm.io/gorm"
)

type Food struct {
	ID            uint `gorm:"size:50"`
	CategoryId    int
	CategoryName  string   `gorm:"size: 150"`
	Category      Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UnitId        int
	UnitName      string `gorm:"size: 150"`
	Unit          Unit   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TaxId         int
	Tax           Tax    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FoodName      string `gorm:"size: 200"`
	Slug          string `gorm:"size: 120"`
	Code          string `gorm:"size: 100"`
	Sku           string `gorm:"size: 100"`
	Discount      int
	Weight        int
	PricePurchase float64
	PriceSelling  float64
	Picture       string
	Tags          string
	Content       string
	Condition     string // baru or lama
	InStock       bool   // ready or not ready
	Note          string
	Status        int
	CreatedBy     int
	UpdatedBy     int
	gorm.Model
}
