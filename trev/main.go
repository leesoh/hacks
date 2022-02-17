package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	iso_format = "2006-01-02 15:04:05 -0700"
)

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			c := sc.Text()
			printTime(c)
		}
	} else {
		printTime(os.Args[1])
	}
}

func printTime(tm string) {
	if len(tm) == 10 {
		utime, _ := strconv.ParseInt(tm, 10, 64)
		t := time.Unix(utime, 0)
		fmt.Println(t.Format(iso_format))
	} else {
		itime, _ := time.Parse(iso_format, tm)
		fmt.Println(itime.Unix())
	}

}
