# is365

is365 takes a list of email addresses on stdin and checks them against Office 365's `GetCredentialType` API.

## Usage

```sh
$ cat users.txt
jdoe@somewhere.com
realuser@nowhere.com
arya@thenorth.com

$ cat users.txt| is365
realuser@nowhere.com
```

Useful parameters include `-c` for concurrency (default: 20) and `-proxy` to send traffic via HTTP proxy.

## Installation

```sh
go install github.com/leesoh/hacks/is365@latest
```

## Thanks

This tool is entirely based off of [this](https://www.redsiege.com/blog/2020/03/user-enumeration-part-2-microsoft-office-365/) blog post, which you should certainly read.

