# Boom

Boom expands CIDR ranges into usable IP addresses (broadcast and network ID removed).

## Usage

```sh
$ echo 192.168.0.0/24 | boom
192.168.0.1
192.168.0.2
...
192.168.0.253
192.168.0.254
```

## Installation

```sh
$ go get -u github.com/leesoh/hacks/boom
```

## Thanks

- This is a paper-thin wrapper around [go-cidr](https://github.com/apparentlymart/go-cidr), who have done the heavy lifting here.

