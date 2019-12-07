package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var wg sync.WaitGroup
	for sc.Scan() {
		wg.Add(1)
		domain := sc.Text()
		go func() {
			defer wg.Done()
			answer, err := net.LookupHost(domain)
			if err != nil {
				return
			}
			for _, aa := range answer {
				fmt.Printf("%s => %s\n", domain, aa)
			}
		}()
	}
	wg.Wait()
}
