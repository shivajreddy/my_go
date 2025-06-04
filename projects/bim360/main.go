package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("BIM 360 API TEST")

	// Load the env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("couldn't load the .env file")
	}

	// Access the environment variables
	client_id := os.Getenv("CLIENT_ID")
	client_secret := os.Getenv("CLIENT_SECRET")

	fmt.Println("client_id:", client_id)
	fmt.Println("client_secret:", client_secret)
}
