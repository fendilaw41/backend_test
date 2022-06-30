package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Order struct {
	Id               int
	StatusOrderId    int
	StatusOrder      StatusOrder `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DateTransaction  datatypes.Date
	TimeTransaction  datatypes.Time
	OrderDetails     []OrderDetail `gorm:"foreignKey:OrderId"`
	MetodePembayaran string        `gorm:"size:200"`
	Invoice          string
	Weight           int     `gorm:"default:null"`
	GrandTotal       float64 `gorm:"default:null"`
	CreatedBy        int     // dibuat oleh
	UpdatedBy        int     // diperbarui oleh
	gorm.Model
}
