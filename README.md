# Slump

Slump is a simple string template implementation for Go.

[![license](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://github.com/nanoninja/slump/blob/master/LICENSE) [![godoc](https://godoc.org/github.com/nanoninja/slump?status.svg)](https://godoc.org/github.com/nanoninja/slump)
[![build status](https://travis-ci.org/nanoninja/slump.svg)](https://travis-ci.org/nanoninja/slump)
[![Coverage Status](https://coveralls.io/repos/github/nanoninja/slump/badge.svg?branch=master)](https://coveralls.io/github/nanoninja/slump?branch=master)
[![go report card](https://goreportcard.com/badge/github.com/nanoninja/slump)](https://goreportcard.com/report/github.com/nanoninja/slump) [![codebeat](https://codebeat.co/badges/58e89ce4-2fd8-4a93-b624-afdbbb44a6e3)](https://codebeat.co/projects/github-com-nanoninja-slump)

## Installation

    go get github.com/nanoninja/slump

## Getting Started

After installing Go and setting up your
[GOPATH](http://golang.org/doc/code.html#GOPATH), create your first `.go` file.

``` go
package main

import "github.com/nanoninja/slump"

func main() {
    s := slump.Str("Hello, {.name}", slump.Values{"name": "Gopher"})

    println(s)
}
```

## Usage examples

### Creating an instance

``` go
s := slump.New("Hello, {.name}")
s.Set("name", "Gopher")

println(s.String())
```

### Getting keys of the values

``` go
s := slump.New("Hello, {.name}")
s.Add(slump.Values{"a": 1, "b": 2, "c": 3})

println(strings.Join(s.Keys(), ", ")) // a, b, c
```


### Getting formatted error

``` go
path := "filename.txt"

err := slump.Err("no such file or directory: {.path}", slump.Values{"path": path})

println(err.Error())
```

### Formatting floating point number

```go
s := slump.Str(`Number: {printf "%.2f" .num}`, slump.Values{"num": 0.393752})

println(s) // Number: 0.39
```

### Using object

```go
user := struct {
    Name string
}{
    Name: "Gopher",
}

s := slump.Str(`Hello, {.user.Name} `, slump.Values{"user": user})

println(s) // Hello, Gopher
```


## License

Slump is licensed under the Creative Commons Attribution 3.0 License, and code is licensed under a [BSD license](https://github.com/nanoninja/slump/blob/master/LICENSE).
