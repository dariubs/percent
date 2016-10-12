percent.go [![Build Status](https://travis-ci.org/dariubs/percent.go.svg?branch=master)](https://travis-ci.org/dariubs/percent.go)
=======
Calculate percentage in Golang.

Install
-------
```shell
go get github.com/dariubs/percent.go
```

Usage
-----
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
* What is the percentage increase/decrease from [number1] to [number2]
*/
percent.Change(20, 60) // return 200
```

License
-------
MIT

Author
------
Dariush Abbasi ([@dariubs](https://github.com/dariubs) )
