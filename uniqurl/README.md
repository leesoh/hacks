# uniqurl

Uniqurl takes a list of URLs on stdin and deduplicates the domains, returning only the HTTPS URL for a given domain pair. If no HTTPS URL exists, the HTTP URL is returned.

## Usage

```sh
$ cat urls.txt
https://www.google.com
http://www.google.com
https://www.facebook.com
http://www.facebook.com
https://www.yahoo.com
http://www.uber.com

$ cat urls.txt | uniqurl
https://www.google.com
https://www.facebook.com
https://www.yahoo.com
http://www.uber.com
```

## Installation

```sh
$ go get -u github.com/leesoh/hacks/uniqurl
```

## Thanks

@dee-see for the idea.
