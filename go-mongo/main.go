package main

import (
	"fmt"
	"os"
	// "github.com/joho/godotenv"
)

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
	fmt.Println(os.Getenv("MONGO_URL"))
}
