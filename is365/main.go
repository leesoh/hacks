package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	concurrency := flag.Int("c", 20, "Concurrency level")
	proxy := flag.String("proxy", "", "HTTP proxy URL")
	verbose := flag.Bool("verbose", false, "Verbose output")
	flag.Parse()

	users := make(chan string)
	var wg sync.WaitGroup
	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for username := range users {
				err := checkEmail(username, *proxy)
				if err != nil {
					if *verbose {
						fmt.Println(err)
					}
					continue
				}
				fmt.Println(username)
			}
		}()
	}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		u := strings.TrimSpace(sc.Text())
		if u == "" {
			continue
		}
		users <- u
	}
	close(users)
	wg.Wait()
}
