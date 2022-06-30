package models

type OrderDetail struct {
	ID            uint `gorm:"size:50"`
	OrderId       int
	Order         Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FoodId        int
	Food          Food   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FoodName      string `gorm:"size:200"`
	CategoryName  string `gorm:"size:200"`
	UnitName      string `gorm:"size:200"`
	Qty           int    `gorm:"default:null"`
	Price         int    `gorm:"default:null"`
	TotalDiscount int    `gorm:"default:null"`
	TotalWeight   int    `gorm:"default:null"`
	Total         int    `gorm:"default:null"`
}
