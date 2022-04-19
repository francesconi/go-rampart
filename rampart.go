package rampart

import (
	"time"

	"golang.org/x/exp/constraints"
)

// Relation represents how two Intervals relate to each other.
type Relation int

const (
	RelationUnknown Relation = iota

	/*
		Interval x is before Interval y.

		    +---+
		    | x |
		    +---+
		          +---+
		          | y |
		          +---+
	*/
	RelationBefore

	/*
		Interval x meets Interval y.

		    +---+
		    | x |
		    +---+
		        +---+
		        | y |
		        +---+
	*/
	RelationMeets

	/*
		Interval x overlaps Interval y.

		    +---+
		    | x |
		    +---+
		      +---+
		      | y |
		      +---+
	*/
	RelationOverlaps

	/*
		Interval x is finished by Interval y.

		    +-----+
		    |  x  |
		    +-----+
		      +---+
		      | y |
		      +---+
	*/
	RelationFinishedBy

	/*
		Interval x contains Interval y.

		    +-------+
		    |   x   |
		    +-------+
		      +---+
		      | y |
		      +---+
	*/
	RelationContains

	/*
		Interval x starts Interval y.

		    +---+
		    | x |
		    +---+
		    +-----+
		    |  y  |
		    +-----+
	*/
	RelationStarts

	/*
		Interval x is equal to Interval y.

		    +---+
		    | x |
		    +---+
		    +---+
		    | y |
		    +---+
	*/
	RelationEqual

	/*
		Interval x is started by Interval y.

		    +-----+
		    |  x  |
		    +-----+
		    +---+
		    | y |
		    +---+
	*/
	RelationStartedBy

	/*
		Interval x is during Interval y.

		      +---+
		      | x |
		      +---+
		    +-------+
		    |   y   |
		    +-------+
	*/
	RelationDuring

	/*
		Interval x finishes Interval y.

		      +---+
		      | x |
		      +---+
		    +-----+
		    |  y  |
		    +-----+
	*/
	RelationFinishes

	/*
		Interval x is overlapped by Interval y.

		      +---+
		      | x |
		      +---+
		    +---+
		    | y |
		    +---+
	*/
	RelationOverlappedBy

	/*
		Interval x is met by Interval y.

		        +---+
		        | x |
		        +---+
		    +---+
		    | y |
		    +---+
	*/
	RelationMetBy

	/*
		Interval x is after Interval y.

		          +---+
		          | x |
		          +---+
		    +---+
		    | y |
		    +---+
	*/
	RelationAfter
)

// Interval represents two values, the lesser and the greater.
// Both must be either of the same ordered type or time type.
type Interval[T any] struct {
	x, y T
	cmp  func(T, T) int
}

// NewIntervalFunc returns an Interval out of x and y so that the Interval
// can be sorted on construction by the given comparison function.
func NewIntervalFunc[T any](x, y T, cmp func(T, T) int) Interval[T] {
	if cmp(x, y) < 0 {
		return Interval[T]{x, y, cmp}
	}
	return Interval[T]{y, x, cmp}
}

func NewOrdInterval[T constraints.Ordered](x, y T) Interval[T] {
	return NewIntervalFunc(x, y, func(t1, t2 T) int {
		if t1 < t2 {
			return -1
		}
		if t1 == t2 {
			return 0
		}
		return 1
	})
}

func NewTimeInterval(x, y time.Time) Interval[time.Time] {
	return NewIntervalFunc(x, y, func(t1, t2 time.Time) int {
		if t1.Before(t2) {
			return -1
		}
		if t1.Equal(t2) {
			return 0
		}
		return 1
	})
}

// Lesser returns the lesser value from an Interval.
func (i Interval[T]) Lesser() T {
	return i.x
}

// Greater returns the greater value from an Interval.
func (i Interval[T]) Greater() T {
	return i.y
}

// IsEmpty returns true if the given Interval is empty, false otherwise.
// An Interval is empty if its lesser equals its greater.
func (i Interval[T]) IsEmpty() bool {
	return i.cmp(i.x, i.y) == 0
}

// IsNonEmpty returns true if the given Interval is non-empty, false otherwise.
// An Interval is non-empty if its lesser is not equal to its greater.
func (i Interval[T]) IsNonEmpty() bool {
	return !i.IsEmpty()
}

// Relates tells you how Interval x relates to Interval y.
// Consult the Relation documentation for an explanation
// of all the possible results.
func (x Interval[T]) Relate(y Interval[T]) Relation {
	lxly := x.cmp(x.Lesser(), y.Lesser())
	lxgy := x.cmp(x.Lesser(), y.Greater())
	gxly := x.cmp(x.Greater(), y.Lesser())
	gxgy := x.cmp(x.Greater(), y.Greater())
	switch {
	case lxly == 0 && gxgy == 0:
		return RelationEqual
	case gxly < 0:
		return RelationBefore
	case lxly < 0 && gxly == 0 && gxgy < 0:
		return RelationMeets
	case gxly == 0:
		return RelationOverlaps
	case lxly > 0 && lxgy == 0 && gxgy > 0:
		return RelationMetBy
	case lxgy == 0:
		return RelationOverlappedBy
	case lxgy > 0:
		return RelationAfter
	case lxly < 0 && gxgy < 0:
		return RelationOverlaps
	case lxly < 0 && gxgy == 0:
		return RelationFinishedBy
	case lxly < 0 && gxgy > 0:
		return RelationContains
	case lxly == 0 && gxgy < 0:
		return RelationStarts
	case lxly == 0 && gxgy > 0:
		return RelationStartedBy
	case lxly > 0 && gxgy < 0:
		return RelationDuring
	case lxly > 0 && gxgy == 0:
		return RelationFinishes
	case lxly > 0 && gxgy > 0:
		return RelationOverlappedBy
	default:
		return RelationUnknown
	}
}

// Inverts a Relation. Every Relation has an inverse.
func (r Relation) Invert() Relation {
	switch r {
	case RelationAfter:
		return RelationBefore
	case RelationBefore:
		return RelationAfter
	case RelationContains:
		return RelationDuring
	case RelationDuring:
		return RelationContains
	case RelationEqual:
		return RelationEqual
	case RelationFinishedBy:
		return RelationFinishes
	case RelationFinishes:
		return RelationFinishedBy
	case RelationMeets:
		return RelationMetBy
	case RelationMetBy:
		return RelationMeets
	case RelationOverlappedBy:
		return RelationOverlaps
	case RelationOverlaps:
		return RelationOverlappedBy
	case RelationStartedBy:
		return RelationStarts
	case RelationStarts:
		return RelationStartedBy
	default:
		return RelationUnknown
	}
}
