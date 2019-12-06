package main

import (
	"bufio"
	"fmt"
	"html"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		fmt.Println(html.UnescapeString(sc.Text()))
	}
}
