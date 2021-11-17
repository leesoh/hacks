package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	verbose := flag.Bool("verbose", false, "Verbose output")
	lower := flag.Bool("lower", false, "Lowercase generated addresses")
	suffix := flag.String("suffix", "", "Suffix to append")
	flag.Parse()
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		name := strings.TrimSpace(sc.Text())
		if name == "" {
			continue
		}
		addresses, err := createAddresses(name, *suffix)
		if err != nil {
			if *verbose {
				fmt.Println(err)
			}
			continue
		}
		for _, address := range addresses {
			if *lower {
				fmt.Println(strings.ToLower(address))
			} else {
				fmt.Println(address)
			}
		}
	}
}
