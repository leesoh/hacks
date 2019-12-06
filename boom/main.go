package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/apparentlymart/go-cidr/cidr"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		c := sc.Text()
		_, net, err := net.ParseCIDR(c)
		if err != nil {
			continue
		}
		start, _ := cidr.AddressRange(net)
		// remove network ID and broadcast
		count := uint64(cidr.AddressCount(net) - 2)
		// we will increment this as we go
		currentIP := cidr.Inc(start)
		for i := uint64(0); i < count; i++ {
			fmt.Println(currentIP)
			currentIP = cidr.Inc(currentIP)
		}
	}
}
