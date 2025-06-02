package main

import (
	"fmt"

	"server/database"
)

func main() {
	fmt.Println("SERVER")

	// Connect the database
	database.Setup()
}
