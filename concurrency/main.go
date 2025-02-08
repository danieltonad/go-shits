package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
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
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
		"https://jsonplaceholder.typicode.com/posts/5",
		"https://jsonplaceholder.typicode.com/posts/6",
		"https://jsonplaceholder.typicode.com/posts/7",
		"https://jsonplaceholder.typicode.com/posts/8",
		"https://jsonplaceholder.typicode.com/posts/9",
		"https://jsonplaceholder.typicode.com/posts/10",
		"https://jsonplaceholder.typicode.com/posts/11",
		"https://jsonplaceholder.typicode.com/posts/12",
		"https://jsonplaceholder.typicode.com/posts/13",
		"https://jsonplaceholder.typicode.com/posts/14",
		"https://jsonplaceholder.typicode.com/posts/15",
		"https://jsonplaceholder.typicode.com/posts/16",
		"https://jsonplaceholder.typicode.com/posts/17",
		"https://jsonplaceholder.typicode.com/posts/18",
		"https://jsonplaceholder.typicode.com/posts/19",
		"https://jsonplaceholder.typicode.com/posts/20",
		"https://jsonplaceholder.typicode.com/posts/21",
		"https://jsonplaceholder.typicode.com/posts/22",
		"https://jsonplaceholder.typicode.com/posts/23",
		"https://jsonplaceholder.typicode.com/posts/24",
	}

	var wg sync.WaitGroup
	results := make(chan string, len(urls)) // Buffered channel to collect results

	// Launch a goroutine for each URL
	for _, url := range urls {
		wg.Add(1)
		go fetchURL(&wg, url, results)
	}

	wg.Wait()      // Wait for all goroutines to finish
	close(results) // Close the results channel

	// Print results
	for result := range results {
		fmt.Println(result)
	}
}
