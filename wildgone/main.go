package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/leesoh/hacks/wildgone/internal/wildgone"
)

func main() {
	resolver := flag.String("r", "8.8.8.8", "Resolver to use")
	flag.Parse()
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		domain := sc.Text()
		result, err := wildgone.GetWildcard(domain, *resolver)
		if err == nil {
			fmt.Println(result)
		}
	}
}
