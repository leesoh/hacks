# Wildgone

Wildgone takes domains on stdin and prints the ones that resolve wildcards. It starts at the base domain and works backwards through each of the subdomains.

## Usage

In the following example, any domain under `bar.baz.com` will resolve:

```sh
$ cat test.txt 
facebook.com
amazon.com
foo.bar.baz.com
bar.baz.com
baz.com

$ cat test.txt | ./wildgone -r 8.8.8.8
bar.baz.com
bar.baz.com
```

## Installation

```sh
go get -u github.com/leesoh/hacks/wildgone
```
