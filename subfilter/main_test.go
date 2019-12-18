package main

import (
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetSubdomains(t *testing.T) {
	var tests = []struct {
		input string
		want  []string
	}{
		{input: "www.google.com", want: []string{"www"}},
		{input: "dev.tools.yahoo.co.uk", want: []string{"dev", "tools"}},
		{input: "xasfdsa.google.com", want: []string{"xasfdsa"}},
		{input: "www.yahoo.com", want: []string{"www"}},
	}
	for _, test := range tests {
		got, err := getSubdomains(test.input)
		if err != nil {
			log.Fatal(err)
		}
		diff := cmp.Diff(got, test.want)
		if diff != "" {
			t.Errorf("got %s want %s", got, test.want)
		}
	}
}
