# DR

DR is a simple tool that attempts to resolve a given domain and returns the result.

## Usage

```sh
$ cat domains.txt
www.facebook.com
www.google.com
safdsafsafsadf.dsafsafsadfa.com

$ cat domains.txt | dr
www.google.com => 172.217.3.164
www.google.com => 2607:f8b0:400a:808::2004
www.facebook.com => 157.240.3.35
www.facebook.com => 2a03:2880:f101:83:face:b00c:0:25de

$ cat domains.txt | dr -not
safdsafsafsadf.dsafsafsadfa.com
```

## Installation

```sh
$ go get -u github.com/leesoh/hacks/dr
```
