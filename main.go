package main

import (
	"backend_test/database"
)

func main() {
	database.ConfigDB()
	database.DbMigrateFreshSeed() // hidupkan jika ingin migrasi table
	// dadu.SoalPermainan() // matikan jika ingin migrasi table
}
