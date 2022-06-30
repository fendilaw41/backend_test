package seeder

import "backend_test/models"

func (s Seed) CartSeed() {

	cart := []models.Cart{
		{
			FoodId: 1,
			Qty:    2,
		},
		{
			FoodId: 2,
			Qty:    2,
		},
		{
			FoodId: 1,
			Qty:    2,
		},
	}

	s.db.Create(&cart)
}
