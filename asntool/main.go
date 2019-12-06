package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/leesoh/hacks/asntool/internal/bgpview"
)

func main() {
	search := flag.String("search", "", "Retrieve associated ASNs")
	flag.Parse()

	s := *search
	if s != "" {
		results, err := bgpview.Search(s)
		if err != nil {
			log.Fatal(err)
		}
		for _, rr := range results {
			fmt.Println(rr)
		}
		os.Exit(0)
	} else {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			asn := strings.TrimSpace(sc.Text())
			results, err := bgpview.LookupIP(asn)
			if err != nil {
				log.Fatal(err)
			}
			for _, rr := range results {
				fmt.Println(rr)
			}
			os.Exit(0)
		}
	}
}
