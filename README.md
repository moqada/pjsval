# pjsval

[![Build Status][travis-image]][travis-url]
[![Doc][godoc-image]][godoc-url]
[![License][license-image]][license-url]

Generating Go Validator file from JSON Hyper Schema (via [prmd](https://github.com/interagent/prmd)).

## Demo

```
# from file
$ pjsval schema.json > validator.go
# from stdin
$ cat schema.json | pjsval > validator.go 
# output file
$ pjsval schema.json -o validator.go
# change package
$ pjsval schema.json -p validator > validator/validator.go 
```

## Installation

### CLI

```
$ go get github.com/moqada/pjsval/cmd/pjsval
```

### Internal libs

```
$ go get github.com/moqada/pjsval
```

## Usage

### CLI

```
usage: pjsval [<flags>] [<file>]

Flags:
  -h, --help            Show context-sensitive help (also try --help-long and --help-man).
  -o, --output=OUTPUT   Path of Go struct file
  -p, --package="main"  Package name for Go validator file
      --version         Show application version.

Args:
  [<file>]  Path of JSON Schema
```

### Internals

Output Example: [./pjsval_test.go](./pjsval_test.go)


[godoc-url]: https://godoc.org/github.com/moqada/pjsval
[godoc-image]: https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square
[travis-url]: https://travis-ci.org/moqada/pjsval
[travis-image]: https://img.shields.io/travis/moqada/pjsval.svg?style=flat-square
[license-url]: http://opensource.org/licenses/MIT
[license-image]: https://img.shields.io/github/license/moqada/pjsval.svg?style=flat-square
