package rampart

import (
	"fmt"
	"time"

	"golang.org/x/exp/constraints"
)

type comparisonResult int

const (
	lt comparisonResult = iota + 1
	eq
	gt
)

func compare[T constraints.Ordered | time.Time](x, y T) comparisonResult {
	switch x := any(x).(type) {
	case time.Time:
		return compareTime(x, any(y).(time.Time))
	case int:
		return compareOrd(x, any(y).(int))
	case int8:
		return compareOrd(x, any(y).(int8))
	case int16:
		return compareOrd(x, any(y).(int16))
	case int32:
		return compareOrd(x, any(y).(int32))
	case int64:
		return compareOrd(x, any(y).(int64))
	case uint:
		return compareOrd(x, any(y).(uint))
	case uint8:
		return compareOrd(x, any(y).(uint8))
	case uint16:
		return compareOrd(x, any(y).(uint16))
	case uint32:
		return compareOrd(x, any(y).(uint32))
	case uint64:
		return compareOrd(x, any(y).(uint64))
	case uintptr:
		return compareOrd(x, any(y).(uintptr))
	case float32:
		return compareOrd(x, any(y).(float32))
	case float64:
		return compareOrd(x, any(y).(float64))
	case string:
		return compareOrd(x, any(y).(string))
	default:
		panic(fmt.Sprintf("unhandled type %T", x))
	}
}

func compareOrd[T constraints.Ordered](x, y T) comparisonResult {
	if x < y {
		return lt
	}
	if x == y {
		return eq
	}
	return gt
}

func compareTime(x, y time.Time) comparisonResult {
	if x.Before(y) {
		return lt
	}
	if x.Equal(y) {
		return eq
	}
	return gt
}
