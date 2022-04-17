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
type Interval[T constraints.Ordered | time.Time] struct {
	x, y T
}

// NewInterval returns an Interval out of x and y so that the Interval
// can be sorted on construction.
func NewInterval[T constraints.Ordered | time.Time](x, y T) Interval[T] {
	if compare(x, y) == lt {
		return Interval[T]{x, y}
	}
	return Interval[T]{y, x}
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
	return compare(i.x, i.y) == eq
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
	lxly := compare(x.Lesser(), y.Lesser())
	lxgy := compare(x.Lesser(), y.Greater())
	gxly := compare(x.Greater(), y.Lesser())
	gxgy := compare(x.Greater(), y.Greater())
	switch {
	case lxly == eq && gxgy == eq:
		return RelationEqual
	case gxly == lt:
		return RelationBefore
	case lxly == lt && gxly == eq && gxgy == lt:
		return RelationMeets
	case gxly == eq:
		return RelationOverlaps
	case lxly == gt && lxgy == eq && gxgy == gt:
		return RelationMetBy
	case lxgy == eq:
		return RelationOverlappedBy
	case lxgy == gt:
		return RelationAfter
	case lxly == lt && gxgy == lt:
		return RelationOverlaps
	case lxly == lt && gxgy == eq:
		return RelationFinishedBy
	case lxly == lt && gxgy == gt:
		return RelationContains
	case lxly == eq && gxgy == lt:
		return RelationStarts
	case lxly == eq && gxgy == gt:
		return RelationStartedBy
	case lxly == gt && gxgy == lt:
		return RelationDuring
	case lxly == gt && gxgy == eq:
		return RelationFinishes
	case lxly == gt && gxgy == gt:
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
