# SANbox

SANbox takes a list of domains on stdin and returns any SANs associated with any certificate found on the host. It will strip wildcards (e.g. `*.example.com` becomes `example.com`) and will not print any domain where the input is identical to the output (i.e. a certificate with a single name associated with it). 

## Usage

```sh
$ cat domains.txt
example.com
foo.example.com
1.2.3.4

$ cat domains.txt | sanbox
[example.com] example.com
[example.com] secret.example.com
[example.com] really-secret.example.com
[foo.example.com] foo.example.com
[foo.example.com] something.otherdomain.com
[1.2.3.4] ips.work.too.com

$ cat domains.txt | sanbox -subs-only
[example.com] example.com
[example.com] secret.example.com
[example.com] really-secret.example.com
[foo.example.com] foo.example.com
```

## Installation

```sh
$ go get -u github.com/leesoh/hacks/sanbox
```

## Thanks

- @Michael1026 for the idea :)

## References

- [x509](https://golang.org/pkg/crypto/x509/#Certificate)
