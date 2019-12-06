package bgpview

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/leesoh/hacks/asntool/internal/http"
)

const endpoint = "https://api.bgpview.io/"

// Search retrieves the ASNs associated with the provided search term
func Search(term string) ([]string, error) {
	var results []string
	v := url.Values{}
	v.Add("query_term", term)

	u, err := url.Parse(endpoint)
	if err != nil {
		return results, nil
	}
	u.Path = "search"
	u.RawQuery = v.Encode()

	resp, err := http.GetResponse(u.String())
	if err != nil {
		return results, err
	}
	results, err = getASNs(resp)
	if err != nil {
		return results, nil
	}
	return results, nil
}

// LookupIP retrieves the IPs associated with the provided ASN
func LookupIP(asn string) ([]string, error) {
	var results []string
	qs := fmt.Sprintf("asn/" + asn + "/prefixes")
	u, err := url.Parse(endpoint)
	if err != nil {
		return results, nil
	}
	u.Path = qs

	resp, err := http.GetResponse(u.String())
	if err != nil {
		return results, nil
	}
	results, err = getIPs(resp)
	if err != nil {
		return results, nil
	}
	return results, nil
}

// getASNs retrieves the ASNs from a blob of JSON
func getASNs(resp []byte) ([]string, error) {
	var results []string
	var s SearchResults
	err := json.Unmarshal(resp, &s)
	if err != nil {
		return results, err
	}
	for _, aa := range s.Data.Asns {
		r := fmt.Sprintf("[%s] %d", aa.Name, aa.Asn)
		results = append(results, r)
	}
	return results, nil
}

// getIPs retrieves the IPs associated with a given ASN
func getIPs(resp []byte) ([]string, error) {
	var results []string
	var p ASNPrefixResults
	err := json.Unmarshal(resp, &p)
	if err != nil {
		return results, err
	}
	for _, pp := range p.Data.IPv4Prefixes {
		r := fmt.Sprintf("[%s] %s", pp.Name, pp.Prefix)
		results = append(results, r)
	}
	return results, nil
}
