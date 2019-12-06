# ASNtool

ASNtool retrieves ASNs associated with a given search term, and IPs associated with ASNs.



## Usage

```sh
$ asntool -search dropbox
[DROPBOX] 19679
[DROPBOXC] 393874
[DROPBOX-CORP] 54372
[Dropbox-Corp-IE] 62190

$ asntool -search dropbox | awk '{print $2}' | asntool
[DROPBOX] 45.58.64.0/20
[DROPBOX] 45.58.66.0/23
[DROPBOX] 45.58.68.0/24
[DROPBOX] 45.58.70.0/24
[DROPBOX] 45.58.71.0/24
[DROPBOX] 45.58.72.0/24
[DROPBOX] 45.58.73.0/24
[DROPBOX] 45.58.74.0/24
[DROPBOX] 45.58.75.0/24
[DROPBOX] 45.58.76.0/23
[DROPBOX] 45.58.78.0/24
```

## Installation

```sh
$ go get -u github.com/leesoh/hacks/asntool
```

## Thanks

- Another thin wrapper project, this time around [BGPview's](https://bgpview.io/) excellent [API](https://bgpview.docs.apiary.io/).
