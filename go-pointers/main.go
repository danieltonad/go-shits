package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")
	myNumber := 24
	var ptr = &myNumber
	fmt.Println("Pointer Address: ", ptr)
	fmt.Println("Pointer Value: ", *ptr)
}
