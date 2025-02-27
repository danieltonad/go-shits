package main

import (
	"fmt"
	"time"
)

func delay_me() {
	time.Sleep(5 * time.Second)
	fmt.Printf("I am delayed %v", time.Second)
}

func main() {
	// today := time.Now()
	// fmt.Println("Today is", today.Format("Monday, January 2, 2006 -> 15:04:05 PM | 01-02-2006 -> 03:04:05 PM"))

	start := time.Now()
	delay_me()
	fmt.Println("Time taken to delay (s)", time.Since(start))
}
