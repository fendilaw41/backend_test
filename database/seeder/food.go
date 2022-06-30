package seeder

import "backend_test/models"

func (s Seed) FoodSeed() {

	food := []models.Food{
		{
			CategoryId:    1,
			CategoryName:  "Nasi",
			UnitId:        2,
			UnitName:      "Porsi",
			TaxId:         1,
			FoodName:      "Nasi Goreng",
			Slug:          "nasi-goreng",
			Code:          "NAS0001",
			Sku:           "001-123456",
			Discount:      0,
			Weight:        10, // gram
			PricePurchase: 15000,
			PriceSelling:  30000,
			Picture:       "default.jpg",
			Tags:          "nasi, goreng",
			Content:       "Nasi Goreng Spesial Murah",
			Condition:     "Baru",
			InStock:       true,
			Note:          "Catatan",
			CreatedBy:     1,
		},
		{
			CategoryId:    1,
			CategoryName:  "Nasi",
			UnitId:        2,
			UnitName:      "Porsi",
			TaxId:         1,
			FoodName:      "Nasi Putih",
			Slug:          "nasi-putih",
			Code:          "NAS000123",
			Sku:           "001-123458",
			Discount:      0,
			Weight:        10, // gram
			PricePurchase: 10000,
			PriceSelling:  20000,
			Picture:       "default.jpg",
			Tags:          "nasi, putih",
			Content:       "Nasi Putih Spesial Murah",
			Condition:     "Baru",
			InStock:       true,
			Note:          "Catatan",
			CreatedBy:     1,
		},
		{
			CategoryId:    1,
			CategoryName:  "Nasi",
			UnitId:        2,
			UnitName:      "Porsi",
			TaxId:         1,
			FoodName:      "Nasi Goreng Spesial",
			Slug:          "nasi-goreng-spesial",
			Code:          "NAS0001",
			Sku:           "001-1234562",
			Discount:      0,
			Weight:        100, // gram
			PricePurchase: 25000,
			PriceSelling:  50000,
			Picture:       "default.jpg",
			Tags:          "nasi, goreng, spesial",
			Content:       "Nasi Goreng Murah",
			Condition:     "Baru",
			InStock:       true,
			Note:          "Catatan",
			CreatedBy:     1,
		},
	}

	s.db.Create(&food)
}
