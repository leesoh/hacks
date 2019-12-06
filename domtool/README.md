# Domtool

Domtool is a simple tool for validating and parsing domains.

## Usage

```sh
$ cat domains.txt
www.example.com
www.facebook.com
m.facebook.com
www.amazon.co.uk
this is a sentence

$ cat domains.txt | domtool
example.com
facebook.com
facebook.com
amazon.co.uk

$ cat domains.txt | domtool -subs
www
www
m
www
```

## Intstallation

```sh
$ go get -u github.com/leesoh/hacks/domtool`
```

## Thanks

- [publicsuffix](https://godoc.org/golang.org/x/net/publicsuffix)
