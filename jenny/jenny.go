package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

type Address struct {
	First  string
	Last   string
	Suffix string
}

func createAddresses(name, suffix string) ([]string, error) {
	nameParts := strings.Split(name, " ")
	a := Address{
		First:  nameParts[0],
		Last:   strings.Join(nameParts[1:], ""),
		Suffix: suffix,
	}
	emails, err := processTemplates(a)
	if err != nil {
		return []string{}, err
	}
	return emails, nil
}

func processTemplates(a Address) ([]string, error) {
	t := `{{ .First }}{{ .Last }}{{ .Suffix }}
{{ .First }}.{{ .Last }}{{ .Suffix }}
{{ slice .First 0 1 }}.{{ .Last }}{{ .Suffix }}
{{ .First }}.{{ slice .Last 0 1 }}{{ .Suffix }}
{{ slice .First 0 1 }}.{{ slice .Last 0 1 }}{{ .Suffix -}}`

	var emails []string
	buf := new(bytes.Buffer)
	tmpl, err := template.New("t").Parse(t)
	if err != nil {
		return emails, fmt.Errorf("error parsing template: %v", err)
	}
	err = tmpl.Execute(buf, a)
	emails = append(emails, string(buf.Bytes()))
	return emails, nil
}
