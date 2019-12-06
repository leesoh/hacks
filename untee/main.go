package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <target>\n", os.Args[0])
		os.Exit(1)
	}
	filename := os.Args[1]
	var data Data
	data.Read(filename)
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		if data.Contains(line) {
			data.Remove(line)
			fmt.Println(line)
		}
	}
	data.Write(filename)
}

// Data contains the... data to be deduplicated
type Data struct {
	Lines map[string]struct{}
}

// Read reads the target file into memory
func (d *Data) Read(path string) {
	d.Lines = make(map[string]struct{})
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		k := sc.Text()
		d.Lines[k] = struct{}{}
	}
}

// Contains checks whether the given value exists in the data
func (d Data) Contains(value string) bool {
	if _, ok := d.Lines[value]; ok {
		return true
	}
	return false
}

// Remove removes the given value from the data
func (d *Data) Remove(value string) {
	delete(d.Lines, value)
}

// Write writes the results to the target file
func (d Data) Write(path string) {
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	for k := range d.Lines {
		f.Write([]byte(k + "\n"))
	}
}
