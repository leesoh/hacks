# HTTPStatus

HTTPStatus takes a list of URLs on stdin and only returns URLs that respond with the specified code.

## Usage

```sh
$ cat urls.txt 
https://httpstat.us/200
https://httpstat.us/302
https://httpstat.us/500

$ cat urls.txt | ./httpstatus -s 200
https://httpstat.us/200
```

## Installation

```sh
$ go get -u github.com/leesoh/hacks/httpstatus
```
