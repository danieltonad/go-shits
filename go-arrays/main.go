package main

import (
	"fmt"
	"sort"
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

func slices() {
	var fruits = []string{"Apple", "Banana", "Mango"}
	fruits = append(fruits, "Orange")
	fmt.Println(fruits)

	scores := make([]int, 5)
	scores[0] = 1
	scores[1] = 264
	scores[2] = 394
	scores[3] = 4
	scores[4] = 5
	scores = append(scores, 856)
	scores = append(scores, 7)
	sort.Ints(scores)
	fmt.Println(scores)
}

func main() {
	slices()
}
