# percent.go

[![Build Status](https://travis-ci.org/dariubs/percent.svg?branch=master)](https://travis-ci.org/dariubs/percent) [![GoDoc](https://godoc.org/github.com/dariubs/percent?status.svg)](https://godoc.org/github.com/dariubs/percent) [![Go Report Card](https://goreportcard.com/badge/github.com/dariubs/percent)](https://goreportcard.com/report/github.com/dariubs/percent)

Calculate percentage in Golang.

## Install

```shell
go get github.com/dariubs/percent
```

## Usage

```go
/*
* Calculate what is [percent]% of [number]
*/
percent.Percent(25, 2000) // return 500
percent.PercentFloat(25.0, 2000) // return 500.0

/*
* Calculate [number1] is what percent of [number2]
*/
percent.PercentOf(300, 2400) // return 12.5

/*
* Calculate what is the percentage increase/decrease from [number1] to [number2]
*/
percent.Change(20, 60) // return 200
```

## Documentation

[GoDoc](https://godoc.org/github.com/dariubs/percent)

## License

MIT

## Author

Dariush Abbasi ([@dariubs](https://github.com/dariubs) )
