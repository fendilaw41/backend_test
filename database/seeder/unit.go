package seeder

import (
	"backend_test/models"
)

func (s Seed) UnitSeed() {

	units := []models.Unit{
		{
			Name:      "Pcs",
			CreatedBy: 1,
		},
		{
			Name:      "Porsi",
			CreatedBy: 1,
		},
	}

	s.db.Create(&units)
}
