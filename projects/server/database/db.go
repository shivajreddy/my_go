package database

import (
	"fmt"
	"log"

	// . "server/database/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connect() *gorm.DB {
	// Connect to database
	db, err := gorm.Open(sqlite.Open("/mnt/c/Users/sreddy/Desktop/test.db"), &gorm.Config{})
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Connect to database")
	}
	return db
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func (p Product) String() string {
	return fmt.Sprintf("Product[ID=%d, Code=%s, Price=%d]", p.ID, p.Code, p.Price)
}

func Setup() {
	db := connect()

	// Drop the entire table
	db.Migrator().DropTable(&Product{})

	// Migrate the schema
	db.AutoMigrate(Product{})

	// /* Add bunch of rows
	{
		products := []Product{
			{Code: "A01", Price: 100},
			{Code: "A01", Price: 150},
			{Code: "B02", Price: 200},
			{Code: "C03", Price: 300},
			{Code: "C03", Price: 320},
			{Code: "D04", Price: 370},
			{Code: "E05", Price: 280},
			{Code: "F06", Price: 210},
			{Code: "G07", Price: 111},
			{Code: "H08", Price: 135},
		}
		db.Create(&products)
	}
	// */

	/* Delete a target product
	{
		var products []Product
		db.Find(&products)
		count := 0
		for _, p := range products {
			if p.Code == "C03" {
				fmt.Printf("%d : Deleting: %s \n", count, p)
				// db.Delete(&p)	// soft delete
				db.Unscoped().Delete(&p) // hard delete
				count++
			}
		}
	}
	// */
}
