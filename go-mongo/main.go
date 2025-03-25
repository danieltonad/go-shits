package main

import (
	"fmt"
	"os"
)

func main() {
	err := LoadEnv(".env")

	if err != nil {
		panic(err)
	}
	fmt.Println(os.Getenv("MONGO_URL"))
}
