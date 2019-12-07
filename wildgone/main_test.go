package main

import (
	"reflect"
	"testing"
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
			t.Errorf("GetETLDPlusOne(%q): %v != %v", test.input, test.want, got)
		}
	}
}

func TestGetSubdomain(t *testing.T) {
	var tests = []struct {
		base string
		etld string
		want string
	}{
		{base: "www.amazon.co.uk", etld: "amazon.co.uk", want: "www"},
		{base: "www.dev.prod.google.com", etld: "google.com", want: "www.dev.prod"},
		{base: "www.test.yahoo.com", etld: "yahoo.com", want: "www.test"},
	}
	for _, test := range tests {
		got := getSubdomain(test.base, test.etld)
		if got != test.want {
			t.Errorf("GetETLDPlusOne(%q): %v != %v", test.base, test.want, got)
		}
	}
}

func TestGetSubdomainParts(t *testing.T) {
	var tests = []struct {
		input string
		want  []string
	}{
		{input: "www", want: []string{"www"}},
		{input: "www.test", want: []string{"www", "test"}},
	}
	for _, test := range tests {
		got := getSubdomainParts(test.input)
		if reflect.DeepEqual(got, test.want) != true {
			t.Errorf("GetSubdomainParts(%q): %v != %v", test.input, test.want, got)
		}
	}
}

func TestCreateGuesses(t *testing.T) {
	var tests = []struct {
		subs []string
		etld string
		want []string
	}{
		{subs: []string{"www", "test"}, etld: "amazon.co.uk", want: []string{"amazon.co.uk", "test.amazon.co.uk", "www.test.amazon.co.uk"}},
		{subs: []string{""}, etld: "google.com", want: []string{"google.com"}},
		{subs: []string{"dev", "ca"}, etld: "yahoo.com.", want: []string{"yahoo.com.", "ca.yahoo.com.", "dev.ca.yahoo.com."}},
	}
	for _, test := range tests {
		got := createGuesses(test.subs, test.etld)
		if reflect.DeepEqual(got, test.want) != true {
			t.Errorf("CreateGuesses(%q): %v != %v", test.subs, test.want, got)
		}
	}
}
