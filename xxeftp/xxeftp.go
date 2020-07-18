// hacky thing that I cobbled together from https://github.com/ONsec-Lab/scripts/blob/master/xxe-ftp-server.rb
package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	// We grab an int then stringify to ensure we don't get "a" as the port or something
	listenPort := flag.Int("p", 21, "Port to listen on")
	flag.Parse()
	p := ":" + strconv.Itoa(*listenPort)

	// Listen on 21/tcp for a connection
	l, err := net.Listen("tcp", p)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// Wait for a connection, then handle it
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go catchFTP(conn)
	}
}

func catchFTP(c net.Conn) {
	// Sent FTP banner
	c.Write([]byte("220 xxeftp\n"))
	for {
		data := make([]byte, 4096)
		n, _ := c.Read(data)
		buf := new(bytes.Buffer)
		buf.Write(data[:n])
		// Grab the verb
		if buf.Len() > 4 {
			cmd := string(buf.Bytes()[:4])
			switch cmd {
			// Plz give us creds
			case "USER":
				c.Write([]byte("331 password please - version check\n"))
			// Password is fine, send away
			case "PASS":
				c.Write([]byte("230 logged in\n"))
			// I don't think we'll ever see this
			case "QUIT":
				c.Write([]byte("221 Goodbye.\n"))
			// Send more data using the existing connection
			case "RETR":
				c.Write([]byte("230 more data please!\n"))
			default:
				c.Write([]byte("230 more data please!\n"))
			}
		}
		fmt.Print(strings.Trim(string(buf.Bytes()[4:]), " "))
	}
	c.Close()
}
