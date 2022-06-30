package seeder

import "backend_test/models"

func (s Seed) CategoryStockSeed() {

	stock := []models.CategoryStock{
		{
			Name: "Qty Awal",
		},
		{
			Name: "Stock In",
		},
		{
			Name: "Stock Out",
		},
		{
			Name: "Stock Opname",
		},
		{
			Name: "Return Stock",
		},
	}

	s.db.Create(&stock)
}
