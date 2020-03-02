package wildgone

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetETLDPlusOne(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{input: "www.amazon.co.uk", want: "amazon.co.uk"},
		{input: "www.google.com", want: "google.com"},
		{input: "www.yahoo.com", want: "yahoo.com"},
	}
	for _, test := range tests {
		got, _ := getETLDPlusOne(test.input)
		if got != test.want {
			t.Errorf("getETLDPlusOne(%q): %v != %v", test.input, test.want, got)
		}
	}
}

func TestStripDots(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{input: ".leading.only", want: "leading.only"},
		{input: "trailing.only.", want: "trailing.only"},
		{input: ".leading.trailing.", want: "leading.trailing"},
	}
	for _, test := range tests {
		got := stripDots(test.input)
		if got != test.want {
			t.Errorf("stripDots(%v) != %v", test.input, got)
		}
	}
}

func TestGetSubdomain(t *testing.T) {
	var tests = []struct {
		domain string
		etld   string
		want   string
	}{
		{domain: "foo.bar.baz.com", etld: "baz.com", want: "foo.bar"},
		{domain: "foo.bar.baz.com.", etld: "baz.com", want: "foo.bar"},
	}
	for _, test := range tests {
		got := getSubdomain(test.domain, test.etld)
		if got != test.want {
			t.Errorf("getSubdomain(%q): %v != %v", test.domain, test.etld, got)
		}
	}
}

func TestGetSubdomainParts(t *testing.T) {
	var tests = []struct {
		input string
		want  []string
	}{
		{input: "foo.bar.baz", want: []string{"foo", "bar", "baz"}},
		{input: ".foo.bar.baz.", want: []string{"foo", "bar", "baz"}},
	}
	for _, test := range tests {
		got := getSubdomainParts(test.input)
		if !cmp.Equal(got, test.want) {
			t.Errorf("getSubdomainParts(%v) != %v", test.input, got)
		}
	}
}

func TestCreateGuesses(t *testing.T) {
	var tests = []struct {
		parts  []string
		base   string
		random string
		want   []string
	}{
		{parts: []string{"foo", "bar"}, base: "baz.com", random: "randomString", want: []string{"randomString.baz.com", "randomString.bar.baz.com", "randomString.foo.bar.baz.com"}},
	}
	for _, test := range tests {
		var w wildcard
		w.parts = test.parts
		w.random = "randomString"
		got := w.createGuesses(test.base, []string{"randomString.baz.com"})
		if !cmp.Equal(got, test.want) {
			t.Errorf("%v != %v", got, test.want)
		}
	}
}
func TestGetWildcard(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{input: "foo.bar.platform.tripadvisor.com", want: "platform.tripadvisor.com"},
		{input: "foo.bar.ext.tripadvisor.com.", want: "ext.tripadvisor.com"},
	}
	for _, test := range tests {
		got, _ := GetWildcard(test.input, "8.8.8.8")
		if got != test.want {
			t.Errorf("GetWildcard(%v) != %v", test.input, got)
		}
	}
}
