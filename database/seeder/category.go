package seeder

import "backend_test/models"

func (s Seed) CategorySeed() {

	category := []models.Category{
		{
			Name:      "Nasi",
			CreatedBy: 1,
			UpdatedBy: 1,
		},
		{
			Name:      "Cemilan",
			CreatedBy: 1,
			UpdatedBy: 1,
		},
		{
			Name:      "Makanan",
			CreatedBy: 1,
			UpdatedBy: 1,
		},
		{
			Name:      "Makanan Khas",
			CreatedBy: 1,
			UpdatedBy: 1,
		},
	}

	s.db.Create(&category)
}
