package seeder

import "backend_test/models"

func (s Seed) StockSeed() {

	stock := []models.Stock{
		{
			FoodId:          1,
			CategoryStockId: 2,
			Total:           200,
		},
		{
			FoodId:          1,
			CategoryStockId: 2,
			Total:           500,
		},
		{
			FoodId:          2,
			CategoryStockId: 2,
			Total:           300,
		},
	}

	s.db.Create(&stock)
}
