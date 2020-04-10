package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
)

type host struct {
	http     bool
	https    bool
	hostname string
}

func main() {
	var hosts = make(map[string]host)
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		u, err := url.Parse(sc.Text())
		if err != nil {
			continue
		}
		var h host
		h.hostname = u.Host
		if hostInSlice(h.hostname, hosts) {
			h = hosts[h.hostname]
		}
		if u.Scheme == "https" {
			h.https = true
		}
		if u.Scheme == "http" {
			h.http = true
		}
		hosts[h.hostname] = h
	}
	for _, hh := range hosts {
		fmt.Println(buildResult(hh))
	}

}

func hostInSlice(h string, hosts map[string]host) bool {
	for _, hh := range hosts {
		if hh.hostname == h {
			return true
		}
	}
	return false
}

func buildResult(h host) string {
	var u url.URL
	u.Host = h.hostname
	if h.https == true {
		u.Scheme = "https"
	} else {
		u.Scheme = "http"
	}
	return u.String()
}
