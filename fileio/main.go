package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFile() {
	databyte, err := ioutil.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("File contetnt: \n", databyte)
}

func main() {
	fmt.Println("File IO")
	content := "Bull shit"

	file, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.WriteString(file, content)

	// os.WriteFile("file.txt", []byte(content), 0644)

	readFile()
}
