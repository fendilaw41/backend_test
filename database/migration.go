package database

import "backend_test/models"

func DbMigration() {
	// TODO: Manajemen User
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.UserRole{}) // TODO: Untuk Pivot Table
	db.AutoMigrate(&models.Permission{})
	db.AutoMigrate(&models.RolePermission{}) // TODO: Untuk Pivot Table
	// TODO: Manajemen Food
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Unit{})
	db.AutoMigrate(&models.Tax{})
	db.AutoMigrate(&models.Food{})
	// TODO: Manajemen Stock
	db.AutoMigrate(&models.CategoryStock{})
	db.AutoMigrate(&models.Stock{})
	// TODO: Manajemen Order
	db.AutoMigrate(&models.Cart{})
	db.AutoMigrate(&models.StatusOrder{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.OrderDetail{})
	// TODO: Manajemen Toko
	db.AutoMigrate(&models.ShopAddress{})

}

func Drop() {
	// TODO: Manajemen User
	db.Migrator().DropTable(&models.User{})
	db.Migrator().DropTable(&models.Role{})
	db.Migrator().DropTable(&models.UserRole{}) // TODO: Untuk Pivot Table
	db.Migrator().DropTable(&models.Permission{})
	db.Migrator().DropTable(&models.RolePermission{}) // TODO: Untuk Pivot Table
	// TODO: Manajemen Stock
	db.Migrator().DropTable(&models.CategoryStock{})
	db.Migrator().DropTable(&models.Stock{})
	// TODO: Manajemen Food
	db.Migrator().DropTable(&models.Tax{})
	db.Migrator().DropTable(&models.Category{})
	db.Migrator().DropTable(&models.Unit{})
	db.Migrator().DropTable(&models.Food{})
	// TODO: Manajemen Order
	db.Migrator().DropTable(&models.Cart{})
	db.Migrator().DropTable(&models.StatusOrder{})
	db.Migrator().DropTable(&models.Order{})
	db.Migrator().DropTable(&models.OrderDetail{})
	// TODO: Manajemen Toko
	db.Migrator().DropTable(&models.ShopAddress{})
}
