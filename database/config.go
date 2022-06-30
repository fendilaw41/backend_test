package database

import (
	"backend_test/database/seeder"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConfigDB() {

	conn, err := DbSetup()
	if err != nil {
		panic(err)
	}
	db = conn

	fmt.Println("=======Database Connection Success=======")
}

func DatabaseMigration() {
	conn, err := DbSetup()
	if err != nil {
		panic(err)
	}
	db = conn
	DbMigration()
	fmt.Println("=======MIGRATION Table Success=======")
}

func DbMigrateFreshSeed() {
	DropMigration()
	DatabaseMigration()
	DatabaseSeeder()
}

func DatabaseSeeder() {
	flag.Parse()
	args := flag.Args()
	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			db, err := DbSetup()
			if err != nil {
				log.Fatalf("Error opening DB: %v", err)
			}
			seeder.Execute(db, args[1:]...)
			os.Exit(0)
			fmt.Println("=======Seeders Success=======")
		}
	}
}

func DropMigration() {
	conn, err := DbSetup()
	if err != nil {
		panic(err)
	}
	db = conn
	Drop()
	fmt.Println("=======DROP Table Success=======")
}

func DbSetup() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/backend_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true,
	})

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	return db, err
}
