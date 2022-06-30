package seeder

import "backend_test/models"

func (s Seed) StatusOrderSeed() {

	status := []models.StatusOrder{
		{
			Name: "Processed",
		},
		{
			Name: "Pending",
		},
		{
			Name: "Success",
		},
		{
			Name: "Take Home",
		},
		{
			Name: "Cancel",
		},
	}

	s.db.Create(&status)
}
