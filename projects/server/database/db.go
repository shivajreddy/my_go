package database

import (
	"fmt"
	"log"

	. "server/database/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connect() *gorm.DB {
	// Connect to database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Connect to database")
	}
	return db
}

func Setup() {
	db := connect()

	// Migrate the schema
	db.AutoMigrate(Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	table := db.Table("products")
	if p, err := table.Rows(); err == nil {
		fmt.Println("p", p)

	}

}
