# Subfilter

Subfilter takes a list of domains on stdin, removes the base domain (e.g. google.com), and then returns a percentage of the domains containing the least-common subdomains.

## Usage

```sh
$ cat domains.txt
www.google.com
xxy-pancakes.google.com
mail.google.com
www.yahoo.com
mail.yahoo.com
mail-dev.yahoo.com
www.facebook.com
mail.facebook.com
www.uber.com
mail.uber.com

$ cat domains.txt | subfilter -p 20
xxy-pancakes.google.com
mail-dev.yahoo.com
```

## Install

```sh
$ go get -u github.com/leesoh/hacks/subfilter
```
