package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func fetchURL(wg *sync.WaitGroup, url string, results chan<- string) {
	defer wg.Done() // Notify when the goroutine completes
	resp, err := http.Get(url)
	if err != nil {
		results <- fmt.Sprintf("Error fetching %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		results <- fmt.Sprintf("Error reading response from %s: %v", url, err)
		return
	}

	results <- fmt.Sprintf("Response from %s: %s", url, string(body)[:100]) // Limit output
}

func main() {
	// List of URLs to fetch
	urls := []string{
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
		"https://ipinfo.io/json",
	}

	var wg sync.WaitGroup
	results := make(chan string, len(urls)) // Buffered channel to collect results
	start := time.Now()
	// Launch a goroutine for each URL
	for _, url := range urls {
		wg.Add(1)
		go fetchURL(&wg, url, results)
	}

	wg.Wait()      // Wait for all goroutines to finish
	close(results) // Close the results channel

	// Print results
	i := 0
	for result := range results {
		i++
		fmt.Printf("No: %v \n", i)
		fmt.Println(result)
	}
	fmt.Printf("Total time taken: %v\n", time.Since(start))
}
