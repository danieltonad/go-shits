package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
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

func fetch_url(url string, ch chan<- string, wg *sync.WaitGroup) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	ch <- string(content)
	wg.Done()
}

func concurrency_trial() {
	urls := []string{
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
	}

	channel := make(chan string, len(urls))
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetch_url(url, channel, &wg)
	}

	wg.Wait()
	close(channel)

	for content := range channel {
		fmt.Println(content)
	}
}

func main() {
	// simple_read()
	// simple_read_type_conversion()
	concurrency_trial()
}
