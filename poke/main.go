package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	flag.Usage = usage
	targetFlag := flag.String("t", "", "Target to check")
	portsFlag := flag.String("p", "", "Ports to check")
	intervalFlag := flag.Int("i", 1, "Tick interval in seconds")
	durationFlag := flag.Int("d", 3600, "Duration in seconds")
	flag.Parse()

	var t target
	err := t.parseArguments(*targetFlag, *portsFlag)
	if err != nil {
		log.Fatal(err)
	}
	t.run(*intervalFlag, *durationFlag)
}

func (t *target) run(interval, duration int) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(time.Duration(duration) * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case <-ticker.C:
			for _, pp := range t.ports {
				err := t.connectTCP(pp)
				if err == nil {
					fmt.Printf("%v:%v is up!\n", t.ipv4, pp)
				}
			}
		}
	}

}

func usage() {
	fmt.Printf("Usage: %s -t 192.168.1.1 -p 80,443\n", os.Args[0])
	flag.PrintDefaults()
}

type target struct {
	ipv4  net.IP
	ports []int
}

func (t *target) parseArguments(target, ports string) error {
	t.ipv4 = net.ParseIP(target)
	p, err := parsePorts(ports)
	if err != nil {
		return err
	}
	t.ports = p
	return nil
}

func parsePorts(ports string) ([]int, error) {
	var result []int
	s := strings.Split(ports, ",")
	for _, ss := range s {
		i, err := strconv.Atoi(ss)
		if err != nil {
			return result, err
		}
		result = append(result, i)
	}
	return result, nil
}

func (t *target) connectTCP(port int) error {
	//address := fmt.Sprintf("%v:%v", t.ipv4, port)
	address, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("%v:%v", t.ipv4, port))
	//conn, err := d.Dial("tcp", address)
	conn, err := net.DialTCP("tcp4", nil, address)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}
