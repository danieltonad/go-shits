package main

import (
	"fmt"
	"gomongo/database"
)

// "github.com/joho/godotenv"

func main() {

	// Load .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	err := Load(".env")

	if err != nil {
		panic(err)
	}

	database.ConnectMongo()

	// Get a collection
	collection := database.GetCollection("users")

	fmt.Println("Using MongoDB Collection:", collection.Name())

}
