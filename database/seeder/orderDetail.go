package seeder

import (
	"backend_test/models"
)

func (s Seed) OrderDetailSeed() {

	order := []models.OrderDetail{
		{
			OrderId:       1,
			FoodId:        1,
			FoodName:      "Nasi Goreng",
			CategoryName:  "Nasi",
			UnitName:      "Porsi",
			Qty:           1,
			Price:         20000,
			TotalDiscount: 0,
			TotalWeight:   1000,
			Total:         1000,
		},
		{
			OrderId:       1,
			FoodId:        2,
			FoodName:      "Nasi Putih",
			CategoryName:  "Nasi",
			UnitName:      "Porsi",
			Qty:           1,
			Price:         20000,
			TotalDiscount: 0,
			TotalWeight:   1000,
			Total:         1000,
		},
		{
			OrderId:       1,
			FoodId:        1,
			FoodName:      "Nasi Goreng Spesial",
			CategoryName:  "Nasi",
			UnitName:      "Porsi",
			Qty:           1,
			Price:         20000,
			TotalDiscount: 0,
			TotalWeight:   1000,
			Total:         1000,
		},
	}

	s.db.Create(&order)
}
