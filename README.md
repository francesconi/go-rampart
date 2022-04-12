# go-rampart

[![Go Reference](https://pkg.go.dev/badge/github.com/francesconi/go-rampart.svg)](https://pkg.go.dev/github.com/francesconi/go-rampart)
![github.com/francesconi/go-rampart](https://github.com/francesconi/go-rampart/workflows/test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/francesconi/go-rampart)](https://goreportcard.com/report/github.com/francesconi/go-rampart)

Go port of the [Haskell Rampart library](https://github.com/tfausak/rampart) by [Taylor Fausak](https://taylor.fausak.me/2020/03/13/relate-intervals-with-rampart).

This package provides types and functions for defining intervals and determining how they relate to each other. This can be useful to determine if and how two ordinal types overlap.

## Install

```sh
go get github.com/francesconi/go-rampart
```

## Example

```go
a := rampart.NewInterval(2, 3)
b := rampart.NewInterval(3, 7)
rel := a.Relate(b)
// rel: RelationMeets
```

![][interval relations]

[interval relations]: ./docs/interval-relations.svg
