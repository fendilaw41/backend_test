package seeder

import (
	"log"
	"reflect"

	"gorm.io/gorm"
)

type Seed struct {
	db *gorm.DB
}

func Execute(db *gorm.DB, seedMethodNames ...string) {
	s := Seed{db}
	seed(s, "UserSeed")
	seed(s, "UnitSeed")
	seed(s, "TaxSeed")
	seed(s, "CategorySeed")
	seed(s, "FoodSeed")
	seed(s, "RoleSeed")
	seed(s, "UserRoleSeed") // untuk pivot
	seed(s, "PermissionSeed")
	seed(s, "RolePermissionSeed") // untuk pivot
	seed(s, "StatusOrderSeed")
	seed(s, "CartSeed")
	seed(s, "CategoryStockSeed")
	seed(s, "StockSeed")
	seed(s, "OrderSeed")
	seed(s, "OrderDetailSeed")

}

func seed(s Seed, seedMethodName string) {
	m := reflect.ValueOf(s).MethodByName(seedMethodName)
	if !m.IsValid() {
		log.Fatal("No method called ", seedMethodName)
	}
	log.Println("Seeding", seedMethodName, "...")
	m.Call(nil)
	log.Println("Seed", seedMethodName, "success")
}
