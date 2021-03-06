# Slump

Slump is a simple string template implementation for Go.

[![license](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://github.com/nanoninja/slump/blob/master/LICENSE) [![godoc](https://godoc.org/github.com/nanoninja/slump?status.svg)](https://godoc.org/github.com/nanoninja/slump) [![build status](https://travis-ci.org/nanoninja/slump.svg)](https://travis-ci.org/nanoninja/slump) [![Coverage Status](https://coveralls.io/repos/github/nanoninja/slump/badge.svg?branch=master)](https://coveralls.io/github/nanoninja/slump?branch=master) [![go report card](https://goreportcard.com/badge/github.com/nanoninja/slump)](https://goreportcard.com/report/github.com/nanoninja/slump) [![codebeat](https://codebeat.co/badges/58e89ce4-2fd8-4a93-b624-afdbbb44a6e3)](https://codebeat.co/projects/github-com-nanoninja-slump)

## Installation

```
go get github.com/nanoninja/slump
```

## Getting Started

After installing Go and setting up your [GOPATH](http://golang.org/doc/code.html#GOPATH), create your first `.go` file.

```go
package main

import "github.com/nanoninja/slump"

func main() {
    s := slump.Str("Hello, {.name}", slump.Value{"name": "Gopher"})

    fmt.Println(s)
}
```

## Usage examples

### Creating an instance

```go
s := slump.New("Hello, {.name}")
s.Value.Set("name", "Gopher")

fmt.Println(s)
```

### Using Value

```go
s := slump.New("coordinate: {.x}, {.y}, {.z}")
s.Value.Add(slump.Value{"x": 1, "y": 2, "z": 3})

fmt.Println(s.Value.IsEmpty()) // false
fmt.Println(s.Value.Count())   // 3
fmt.Println(s.Value.Get("y"))  // 2
fmt.Println(s.Value.Keys())    // [x, y, z] unsorted
```

### Getting formatted error

```go
path := "filename.txt"

err := slump.Err("no such file or directory: {.path}", slump.Value{"path": path})

fmt.Println(err.Error())
```

### Formatting floating point number

```go
s := slump.Str(`Number: {printf "%.2f" .num}`, slump.Value{"num": 0.393752})

fmt.Println(s) // Number: 0.39
```

### Using object

```go
user := struct {
    Name string
}{
    Name: "Gopher",
}

s := slump.Str("Hello, {.user.Name} ", slump.Value{"user": user})

fmt.Println(s) // Hello, Gopher
```

## Benchmarks

Run on MacBook Pro (Retina, 15-inch, Half 2014) 2,5 GHz Intel Core i7 16 Go 1600 MHz DDR3 using Go version go1.7.4 darwin/amd64.

```shell
go test -bench=. -benchmem *.go
```

```go
BenchmarkStr-8          50000         26677 ns/op       10639 B/op          99 allocs/op
BenchmarkErr-8          50000         26806 ns/op       10654 B/op         100 allocs/op
```

## License

Slump is licensed under the Creative Commons Attribution 3.0 License, and code is licensed under a [BSD license](https://github.com/nanoninja/slump/blob/master/LICENSE).
