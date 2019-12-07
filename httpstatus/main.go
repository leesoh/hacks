package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	concurrency := flag.Int("c", 20, "Number of threads")
	timeout := flag.Int("t", 10000, "Timeout in milliseconds")
	status := flag.Int("s", 0, "Status code to check for")
	flag.Parse()

	// Convert to time.Milliseconds
	t := time.Duration(*timeout * 1000000)

	// Implement timeout and don't follow redirects
	client := &http.Client{
		Timeout: t,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	urls := make(chan string)

	// Spin up our worker pool
	var wg sync.WaitGroup
	for i := 0; i < *concurrency; i++ {
		wg.Add(1)

		go func() {
			for url := range urls {
				if StatusCode(client, url, *status) {
					fmt.Println(url)
				}
			}
			wg.Done()
		}()
	}

	// Read stdin and add each URL to our worker pool
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		urls <- scanner.Text()
	}
	close(urls)
	wg.Wait()
}

func StatusCode(client *http.Client, url string, status int) bool {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false
	}
	req.Header.Add("Connection", "close")
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp != nil {
		return status == resp.StatusCode
		resp.Body.Close()
	}
	return false
}
