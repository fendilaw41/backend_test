package main

import "backend_test/database"

func main() {
	database.ConfigDB()
	database.DbMigrateFreshSeed()
}
