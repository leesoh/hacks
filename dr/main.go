package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	not := flag.Bool("not", false, "Only show domains that don't resolve")
	flag.Parse()

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		domain := sc.Text()
		answer, err := net.LookupHost(domain)
		// catch name resolution errors
		switch err.(type) {
		case nil:
		case *net.DNSError:
			if *not {
				fmt.Println(domain)
				return
			}
		}
		for _, aa := range answer {
			fmt.Printf("%s => %s\n", domain, aa)
		}
	}
}
