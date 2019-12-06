package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/publicsuffix"
)

func main() {
	subs := flag.Bool("subs", false, "Return subdomains, not base domains")
	flag.Parse()
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		record := trimDot(sc.Text())
		// ensure the TLD is valid
		_, ok := publicsuffix.PublicSuffix(record)
		if ok {
			switch *subs {
			case true:
				result, _ := getSubdomain(record)
				if len(result) != 0 {
					fmt.Println(result)
				}
			case false:
				result, _ := getETLDPlusOne(record)
				if len(result) != 0 {
					fmt.Println(result)
				}
			}
		}
	}
}

// getETLDPlusOne returns the base domain.
// "www.toys.amazon.co.uk" => "amazon.co.uk"
func getETLDPlusOne(domain string) (string, error) {
	e, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return e, nil
}

// getSubdomain returns everything but the base domain
// "www.toys.amazon.co.uk" => "www.toys"
func getSubdomain(domain string) (string, error) {
	e, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		return "", err
	}
	subdomain := strings.TrimSuffix(domain, e)
	subdomain = strings.TrimSuffix(subdomain, ".")
	return subdomain, nil
}

// trimDot removes any trailing dots
func trimDot(domain string) string {
	if strings.HasSuffix(domain, ".") {
		domain = strings.TrimSuffix(domain, ".")
	}
	return domain
}
