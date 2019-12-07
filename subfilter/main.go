package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"golang.org/x/net/publicsuffix"
)

func main() {
	p := flag.Int("p", 20, "Percent of the least-common subdomains to return, sorted by rarity.")
	flag.Parse()
	var domains []string
	// read domains from stdin
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		d := sc.Text()
		domains = append(domains, d)
	}
	// extract the subs
	subdomains := make(map[string]int)
	for _, dd := range domains {
		s, err := getSubdomains(dd)
		if err != nil {
			fmt.Println(err)
		}
		for _, v := range s {
			subdomains[v]++
		}
	}
	// calculate number of results, rounded to the nearest whole number
	pct := float64(*p) / float64(100)
	numResults := int(math.RoundToEven(float64(len(subdomains)) * pct))
	sortedSubs := SortMap(subdomains)
	i := 0
	// the first ones will be the most rare
	for _, ss := range sortedSubs {
		for _, dd := range domains {
			// we want the rarest subdomains. any domain that has a rare sub
			// is printed.
			if strings.Contains(dd, ss.Key) && i < numResults {
				fmt.Println(dd)
				i++
			} else {
				break
			}
		}
	}
}

// getSubdomains removes the base domain and returns a slice of subdomains.
func getSubdomains(domain string) ([]string, error) {
	domain = strings.TrimSuffix(domain, ".")
	baseDomain, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		return nil, err
	}
	s := strings.Replace(domain, "."+baseDomain, "", -1)
	subs := strings.Split(s, ".")
	return subs, nil
}

// A data structure to hold a key/value pair.
// https://groups.google.com/forum/#!msg/golang-nuts/FT7cjmcL7gw/Gj4_aEsE_IsJ
// there are better ways of doing this now, but this works
type Pair struct {
	Key   string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func SortMap(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}
