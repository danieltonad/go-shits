package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func simple_read() {
	welcome_msg := "Welcome Pussy"
	fmt.Println(welcome_msg)
	age_reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your age: ")
	age, _ := age_reader.ReadString('\n')
	fmt.Println("You are", age, "years old")
}

func simple_read_type_conversion() {
	welcome_msg := "Welcome To Ekpon Store!"
	fmt.Println(welcome_msg)

	rate_reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the rate: ")
	rating, _ := rate_reader.ReadString('\n')
	// fmt.Println("The rating is", rating)

	rating_no, err := strconv.ParseInt(strings.TrimSpace(rating), 10, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The rating is", rating_no)
	}
}

func main() {
	// simple_read()
	simple_read_type_conversion()
}
