package models

import "gorm.io/gorm"

type ShopAddress struct { // alamat toko
	ID uint `gorm:"size:50"`

	FirstNameShop string `gorm:"size:100"`
	LastNameShop  string `gorm:"size:100"`
	FullNameShop  string `gorm:"size:100"`
	Address       string `gorm:"size:100"`
	Provinces     string `gorm:"size:100"`
	City          string `gorm:"size:100"`
	District      string `gorm:"size:100"`
	Subdistrict   string `gorm:"size:100"`
	PostalCode    int
	Icon          string `gorm:"size:100"`
	Picture       string `gorm:"size:100"`

	CreatedBy int // dibuat oleh
	UpdatedBy int // diperbarui oleh
	gorm.Model
}
