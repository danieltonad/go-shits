package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func array() {

	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[3] = "Mango"
	fmt.Println(fruits)

	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[4] = 5

}

type Book struct {
	Id     int
	Name   string
	Author string
}

func slices() {
	var fruits = []string{"Apple", "Banana", "Mango"}
	fruits = append(fruits, "Orange")
	fmt.Println(fruits)

	//using make to create a slice
	scores := make([]int, 5)
	scores[0] = 1
	scores[1] = 264
	scores[2] = 394
	scores[3] = 4
	scores[4] = 5
	scores = append(scores, 856)
	scores = append(scores, 7)
	//playing arround sort with slices
	sort.Ints(scores)
	//remove from slice
	fmt.Println(scores)
	scores = append(scores[:2], scores[3:]...)
	fmt.Println(scores)
}

func maps() {
	cities := make(map[string]string)
	cities["Kenya"] = "Nairobi"
	cities["Uganda"] = "Kampala"
	cities["Tanzania"] = "Dodoma"
	cities["Rwanda"] = "Kigali"
	for index, value := range cities {
		fmt.Println(index, value)
	}
	//delete from map

}

func RandomB() {
	rand.Seed(time.Now().UnixNano())
	dice := rand.Intn(6)

ekpon:
	fmt.Println("I am here")

	switch dice {
	case 1:
		fmt.Println("You rolled a 1")
	case 2:
		fmt.Println("You rolled a 2")
		fallthrough
	case 3:
		fmt.Println("You rolled a 3")
	case 4:
		fmt.Println("You rolled a 4")
		goto ekpon
	case 5:
		fmt.Println("You rolled a 5")
	case 6:
		fmt.Println("You rolled a 6")
	default:
		fmt.Println("Invalid Roll")
	}
}

func add_integers(ints ...int) int {
	total := 0
	for _, value := range ints {
		total += value
	}
	return total
}

func main() {
	// slices()
	// maps()
	// RandomB()
	// fmt.Printf("Total: %v", add_integers(1, 2, -3, 4, 5, 6, 7, 8, 9, 10))
	// book := Book{Id: 1, Name: "The Alchemist", Author: "Paulo Coelho"}
	// if book.Id = 10; book.Id == 1 {
	// 	fmt.Printf("Id: %v, Name: %v and Author: %v", book.Id, book.Name, book.Author)
	// }

	start := time.Now()
	for i := 0; i < 1e11; i++ {
	}
	fmt.Printf("Time taken: %v", time.Since(start))
}
