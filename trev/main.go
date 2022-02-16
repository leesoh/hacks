package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	iso_format = "2006-01-02 15:04:05"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		c := sc.Text()
		// it's unix time
		if len(c) == 10 {
			utime, _ := strconv.ParseInt(c, 10, 64)
			t := time.Unix(utime, 0)
			fmt.Println(t.Format(iso_format))
		} else {
			itime, _ := time.Parse(iso_format, c)
			fmt.Println(itime.Unix())
		}

	}
}
