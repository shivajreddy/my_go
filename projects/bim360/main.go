package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var CLIENT_ID string
var CLIENT_SECRET string

func main() {
	fmt.Println("BIM 360 API TEST")

	// Step0: CLIENT_ID & CLIENT_SECRET
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ERROR: Couldn't load the .env file")
	}

	// Access the environment variable
	CLIENT_ID = os.Getenv("CLIENT_ID")
	CLIENT_SECRET = os.Getenv("CLIENT_SECRET")

	// fmt.Println("CLIENT_ID:", CLIENT_ID)
	// fmt.Println("CLIENT_SECRET:", CLIENT_SECRET)

	// Step 1: Combine <CLIENT_ID>:<CLIENT_SECRET> and encode it in base64
	id_secret := CLIENT_ID + ":" + CLIENT_SECRET
	encoded_string := base64.StdEncoding.EncodeToString([]byte(id_secret))

	fmt.Println("")
	fmt.Println(encoded_string)

	// Step2 : use encoded string to get access token
}
