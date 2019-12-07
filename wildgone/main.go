package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dchest/uniuri"
	"github.com/miekg/dns"
	"golang.org/x/net/publicsuffix"
)

func main() {
	resolver := flag.String("r", "8.8.8.8", "Resolver to use")
	flag.Parse()
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		wildcard := uniuri.New()
		// Trailing dots cause headaches
		base := strings.TrimSuffix(sc.Text(), ".")
		// Get the ETLD+1. This will be appended to each subdomain
		etld, err := getETLDPlusOne(base)
		if err != nil {
			continue
		}
		subdomain := getSubdomain(base, etld)
		subparts := getSubdomainParts(subdomain)
		guesses := createGuesses(subparts, etld)
		for _, b := range guesses {
			n := wildcard + "." + b
			if domainResolves(n, *resolver) {
				fmt.Println(b)
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

// getSubdomain removes the ETLD+1 from a given domain along with the trailing .
func getSubdomain(domain, etld string) string {
	s := strings.Replace(domain, etld, "", 1)
	return strings.TrimSuffix(s, ".")
}

// getSubdomainParts returns a slice of the subdomains
func getSubdomainParts(subdomain string) []string {
	s := strings.Split(subdomain, ".")
	for i := range s {
		s[i] = strings.TrimSpace(s[i])
	}
	return s
}

// createGuesses starts at the root domain and builds new domains by incrementally
// adding its subdomains back.
func createGuesses(subparts []string, etld string) []string {
	var result []string
	if len(subparts) == 1 {
		return []string{etld}
	}
	base := etld
	result = append(result, base)
	for i := len(subparts); i > 0; i-- {
		// www
		s := subparts[i-1]
		// www.yahoo.com
		d := s + "." + base
		result = append(result, d)
		base = d
	}
	return result
}

// domainResolves checks whether the given name resolves
func domainResolves(name, resolver string) bool {
	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(name), dns.TypeA)
	m.RecursionDesired = true

	result, _, err := c.Exchange(m, resolver+":53")
	if err != nil {
		return false
	}
	// if we get any answer, it's a wildcard
	if len(result.Answer) != 0 {
		return true
	}
	return false
}
