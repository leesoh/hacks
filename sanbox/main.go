package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	concurrency := flag.Int("c", 20, "Concurrency level")
	subs := flag.Bool("subs-only", false, "Only include subdomains of the provided domain")
	flag.Parse()
	hosts := make(chan string)
	var wg sync.WaitGroup
	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for hh := range hosts {
				results, err := getSANs(hh)
				if err != nil {
					continue
				}
				for _, rr := range results {
					rr = stripWild(rr)
					// only one name on the cert
					if rr == hh {
						continue
					}
					// wildcard cert, just keep the non-wild bit
					// if we want subs only and it's a subdomain
					if *subs && isSub(rr, hh) {
						fmt.Printf("[%s] %s\n", hh, rr)
						continue
					}
					// if we want subs only and it's not a subdomain
					if *subs && !isSub(rr, hh) {
						continue
					}
					fmt.Printf("[%s] %s\n", hh, rr)
				}
			}
		}()
	}
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		host := strings.ToLower(strings.TrimSpace(sc.Text()))
		if host == "" {
			continue
		}
		hosts <- host
	}
	close(hosts)
	wg.Wait()
}

// getSANs retrieves the Subject Alternate Names from the certificates found on the host
func getSANs(host string) ([]string, error) {
	var results []string
	dialer := net.Dialer{
		Timeout: 10 * time.Second,
	}
	host = formatDomain(host)
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := tls.DialWithDialer(&dialer, "tcp", host, config)
	if err != nil {
		return results, err
	}
	defer conn.Close()
	certs := conn.ConnectionState().PeerCertificates
	for _, cc := range certs {
		names := cc.DNSNames
		for _, nn := range names {
			results = append(results, nn)
		}
	}
	return results, nil
}

// formatDomain ensures the port is specified, defaulting to 443
func formatDomain(domain string) string {
	if strings.Contains(domain, ":") {
		return domain
	}
	domain += ":443"
	return domain
}

// isSub checks whether a domain is a subdomain of another one
func isSub(candidate, domain string) bool {
	if strings.Contains(candidate, domain) {
		return true
	}
	return false
}

func stripWild(domain string) string {
	return strings.Replace(domain, "*.", "", -1)
}
