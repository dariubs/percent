# percent.go [![Build Status](https://travis-ci.org/negah/percent.svg?branch=master)](https://travis-ci.org/negah/percent) [![GoDoc](https://godoc.org/github.com/negah/percent?status.svg)](https://godoc.org/github.com/negah/percent)

Calculate percentage in Golang.

## Install

```shell
go get github.com/negah/percent
```

## Usage

```go
/*
* Calculate what is [percent]% of [number]
*/
percent.Percent(25, 2000) // return 500

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

[GoDoc](https://godoc.org/github.com/negah/percent)

## License

MIT

## Author

Dariush Abbasi ([@dariubs](https://github.com/dariubs) )
