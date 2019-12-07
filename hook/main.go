package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	mPrefix := flag.String("m", "", "Message prefix")
	flag.Parse()
	var c config
	c.Load()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		notification := fmt.Sprintf("%s %s", *mPrefix, text)
		message := Message{
			Text: notification,
		}
		fmt.Println(text)
		err := message.Send(c.HookURL)
		if err != nil {
			log.Fatal(err)
		}
	}
}
