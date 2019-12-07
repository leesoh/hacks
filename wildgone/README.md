# Wildgone

Wildgone takes domains on stdin and prints the ones that resolve wildcards.

## Usage

```sh
$ cat test.txt 
support.com
facebook.com
amazon.com
www.bob.example.com.

$ cat test.txt | ./wildgone -r 8.8.8.8
support.com
```

## Installation

```sh
go get -u github.com/leesoh/hacks/wildgone
```
