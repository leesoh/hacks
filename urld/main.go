package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		in := sc.Text()
		out, err := url.QueryUnescape(in)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(out))
	}
}
