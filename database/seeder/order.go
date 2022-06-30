package seeder

import (
	"backend_test/models"
	"time"

	"gorm.io/datatypes"
)

func (s Seed) OrderSeed() {

	order := []models.Order{
		{
			StatusOrderId:    1,
			DateTransaction:  datatypes.Date(time.Now()),
			TimeTransaction:  datatypes.NewTime(10, 00, 00, 00),
			MetodePembayaran: "Cash",
			Invoice:          "INV-2022-001",
			Weight:           200,
			GrandTotal:       100000,
			CreatedBy:        1,
		},
		{
			StatusOrderId:    3,
			DateTransaction:  datatypes.Date(time.Now()),
			TimeTransaction:  datatypes.NewTime(10, 00, 00, 00),
			MetodePembayaran: "Cash",
			Invoice:          "INV-2022-002",
			Weight:           200,
			GrandTotal:       200000,
			CreatedBy:        1,
		},
		{
			StatusOrderId:    4,
			DateTransaction:  datatypes.Date(time.Now()),
			TimeTransaction:  datatypes.NewTime(10, 00, 00, 00),
			MetodePembayaran: "Cash",
			Invoice:          "INV-2022-003",
			Weight:           200,
			GrandTotal:       500000,
			CreatedBy:        1,
		},
	}

	s.db.Create(&order)
}
