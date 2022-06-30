package seeder

import "backend_test/models"

func (s Seed) TaxSeed() {

	tax := []models.Tax{
		{
			Name: "PPn 10% (Include)",
			Qty:  10,
		},
		{
			Name: "PPn 11% (Include)",
			Qty:  11,
		},
		{
			Name: "PPn 20% (Include)",
			Qty:  20,
		},
	}

	s.db.Create(&tax)
}
