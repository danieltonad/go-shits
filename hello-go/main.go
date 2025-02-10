package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome_msg := "Welcome Pussy"
	fmt.Println(welcome_msg)
	age_reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your age: ")
	age, _ := age_reader.ReadString('\n')
	fmt.Println("You are", age, "years old")
}
