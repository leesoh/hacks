# Jenny

Jenny generates usernames from a list of names provided on stdin.

## Usage

```sh
$ cat names.txt
John Doe
John Doe-Verhulst
John van der Doe
John-John Doe

$ cat names.txt | jenny -suffix "@example.com" -lower
johndoe@example.com
john.doe@example.com
j.doe@example.com
john.d@example.com
j.d@example.com
johndoe-verhulst@example.com
john.doe-verhulst@example.com
j.doe-verhulst@example.com
john.d@example.com
j.d@example.com
johnvanderdoe@example.com
john.vanderdoe@example.com
j.vanderdoe@example.com
john.v@example.com
j.v@example.com
john-johndoe@example.com
john-john.doe@example.com
j.doe@example.com
john-john.d@example.com
j.d@example.com
```

## Installation

```sh
go install github.com/leesoh/hacks/jenny
```
