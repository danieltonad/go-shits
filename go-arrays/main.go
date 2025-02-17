package main

import "fmt"

func main() {

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
	fmt.Println(len(numbers))

}
