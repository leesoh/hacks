package wildgone

import (
	"errors"
	"strings"

	"github.com/dchest/uniuri"
	"github.com/miekg/dns"
	"golang.org/x/net/publicsuffix"
)

func getETLDPlusOne(domain string) (string, error) {
	etld, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		return "", err
	}
	return etld, nil
}

func stripDots(s string) string {
	s = strings.TrimSuffix(s, ".")
	s = strings.TrimPrefix(s, ".")
	return s
}

func getSubdomain(domain, etld string) string {
	domain = stripDots(domain)
	etld = stripDots(etld)
	baseDomain := "." + etld
	subdomain := strings.Replace(domain, baseDomain, "", -1)
	return subdomain
}

func getSubdomainParts(subdomain string) []string {
	subdomain = stripDots(subdomain)
	parts := strings.Split(subdomain, ".")
	return parts
}

func createRandomSubdomain() string {
	return uniuri.New()
}

type wildcard struct {
	parts  []string
	random string
}

func (w wildcard) createGuesses(baseDomain string, guesses []string) []string {
	baseDomain = stripDots(baseDomain)
	// we're out of parts to guess, return
	if len(w.parts) == 0 {
		return guesses
	}
	baseDomain = w.parts[len(w.parts)-1] + "." + baseDomain
	guess := w.random + "." + baseDomain
	// part is now in basedomain, remove it
	w.parts = w.parts[:len(w.parts)-1]
	guesses = append(guesses, guess)
	guesses = w.createGuesses(baseDomain, guesses)
	return guesses
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

// ErrWildErrWildcardNotFound is returned when no wildcard domain can be found
var ErrWildcardNotFound = errors.New("wildcard record not found")

// GetWildcard handles the actual checking for wildcards
func GetWildcard(domain, resolver string) (string, error) {
	var w wildcard
	var result string
	domain = stripDots(domain)
	etld, err := getETLDPlusOne(domain)
	if err != nil {
		return result, err
	}
	etld = stripDots(etld)
	subdomain := getSubdomain(domain, etld)
	w.parts = getSubdomainParts(subdomain)
	w.random = createRandomSubdomain()

	randomPrefix := w.random + "."
	guesses := []string{randomPrefix + etld}
	guesses = w.createGuesses(etld, guesses)

	for _, gg := range guesses {
		if domainResolves(gg, resolver) {
			result = strings.Replace(gg, randomPrefix, "", -1)
			return result, nil
		}
	}
	return result, ErrWildcardNotFound
}
