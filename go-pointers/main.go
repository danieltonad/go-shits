package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")
	myNumber := 24
	var ptr = &myNumber
	*ptr = *ptr * 10
	fmt.Println("Pointer Address: ", ptr)
	fmt.Println("Pointer Value: ", *ptr)
	fmt.Println("No: ", myNumber)
}
