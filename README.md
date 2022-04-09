# go-rampart

![github.com/francesconi/go-rampart](https://github.com/francesconi/go-rampart/workflows/test/badge.svg)

Go port of the [Haskell Rampart library](https://github.com/tfausak/rampart) by [Taylor Fausak](https://taylor.fausak.me/2020/03/13/relate-intervals-with-rampart).

![][interval relations]

## Install

```sh
go get github.com/francesconi/go-rampart
```

## Examples

```go
a := rampart.NewInterval(2, 3)
b := rampart.NewInterval(3, 7)
rel := a.Relate(b) // RelationMeets
```

[interval relations]: ./docs/interval-relations.svg
