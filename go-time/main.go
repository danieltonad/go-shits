package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()
	fmt.Println("Today is", today.Format("Monday, January 2, 2006 -> 15:04:05 PM | 01-02-2006 -> 03:04:05 PM"))
}
