# go-rampart

[![Go Reference](https://pkg.go.dev/badge/github.com/francesconi/go-rampart.svg)](https://pkg.go.dev/github.com/francesconi/go-rampart)
![github.com/francesconi/go-rampart](https://github.com/francesconi/go-rampart/workflows/test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/francesconi/go-rampart)](https://goreportcard.com/report/github.com/francesconi/go-rampart)

Go port of the [Haskell Rampart library](https://github.com/tfausak/rampart) by [Taylor Fausak](https://taylor.fausak.me/2020/03/13/relate-intervals-with-rampart).

This package provides types and functions for defining intervals and determining how they relate to each other. This can be useful to determine if an event happened during a certain time frame, or if two time frames overlap (and if so, how exactly they overlap).

## Install

```sh
go get github.com/francesconi/go-rampart
```

## Examples

```go
a := rampart.NewInterval(2, 3)
b := rampart.NewInterval(3, 7)
rel := a.Relate(b)
// rel: RelationMeets
```

```go
func compareTimes(t1, t2 time.Time) int {
    return int(t1.Sub(t2))
}

a := rampart.NewIntervalFunc(
    time.Date(2022, time.April, 1, 0, 0, 0, 0, time.UTC),
    time.Date(2022, time.April, 8, 0, 0, 0, 0, time.UTC),
    compareTimes,
)
b := rampart.NewIntervalFunc(
    time.Date(2022, time.April, 6, 0, 0, 0, 0, time.UTC),
    time.Date(2022, time.April, 15, 0, 0, 0, 0, time.UTC),
    compareTimes,
)
rel := a.Relate(b)
// rel: RelationOverlaps
```

![][interval relations]

[interval relations]: ./docs/interval-relations.svg
