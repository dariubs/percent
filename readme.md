# percent.go

[![Build Status](https://travis-ci.org/dariubs/percent.svg?branch=master)](https://travis-ci.org/dariubs/percent) [![GoDoc](https://godoc.org/github.com/dariubs/percent?status.svg)](https://godoc.org/github.com/dariubs/percent) [![Go Report Card](https://goreportcard.com/badge/github.com/dariubs/percent)](https://goreportcard.com/report/github.com/dariubs/percent)

Calculate percentage in Golang.

## Install

```shell
go get github.com/dariubs/percent
```

## Usage

```go
// Percent - calculate what %[number1] of [number2] is.
percent.Percent(25, 200) // return 50
percent.PercentFloat(25.0, 200.0) // return 50.0

// PercentOf - calculate what percent [number1] is of [number2].
percent.PercentOf(300, 2400) // return 12.5
percent.PercentOfFloat(300.0, 2400.0) // return 12.5

// Change - calculate the percent increase/decrease from two numbers.  
percent.Change(20, 60) // return 200.0
percent.ChangeFloat(20.0, 60.0) // return 200.0
```

## Documentation

[GoDoc](https://godoc.org/github.com/dariubs/percent)

## License

MIT

## Author

Dariush Abbasi ([@dariubs](https://github.com/dariubs) )
