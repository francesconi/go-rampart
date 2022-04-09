package rampart

import (
	"golang.org/x/exp/constraints"
)

type Relation int

const (
	RelationUnknown Relation = iota
	RelationBefore
	RelationMeets
	RelationOverlaps
	RelationFinishedBy
	RelationContains
	RelationStarts
	RelationEqual
	RelationStartedBy
	RelationDuring
	RelationFinishes
	RelationOverlappedBy
	RelationMetBy
	RelationAfter
)

type Interval[T constraints.Ordered] struct {
	x, y T
}

func NewInterval[T constraints.Ordered](x, y T) Interval[T] {
	if x < y {
		return Interval[T]{x, y}
	}
	return Interval[T]{y, x}
}

func (i Interval[T]) Lesser() T {
	return i.x
}

func (i Interval[T]) Greater() T {
	return i.y
}

func (i Interval[T]) IsEmpty() bool {
	return i.x == i.y
}

func (i Interval[T]) IsNonEmpty() bool {
	return !i.IsEmpty()
}

type comparisonResult int

const (
	LT comparisonResult = iota
	EQ
	GT
)

func compare[T constraints.Ordered](x, y T) comparisonResult {
	if x < y {
		return LT
	} else if x == y {
		return EQ
	} else {
		return GT
	}
}

func (x Interval[T]) Relate(y Interval[T]) Relation {
	lxly := compare(x.Lesser(), y.Lesser())
	lxgy := compare(x.Lesser(), y.Greater())
	gxly := compare(x.Greater(), y.Lesser())
	gxgy := compare(x.Greater(), y.Greater())
	switch {
	case lxly == EQ && gxgy == EQ:
		return RelationEqual
	case gxly == LT:
		return RelationBefore
	case lxly == LT && gxly == EQ && gxgy == LT:
		return RelationMeets
	case gxly == EQ:
		return RelationOverlaps
	case lxly == GT && lxgy == EQ && gxgy == GT:
		return RelationMetBy
	case lxgy == EQ:
		return RelationOverlappedBy
	case lxgy == GT:
		return RelationAfter
	case lxly == LT && gxgy == LT:
		return RelationOverlaps
	case lxly == LT && gxgy == EQ:
		return RelationFinishedBy
	case lxly == LT && gxgy == GT:
		return RelationContains
	case lxly == EQ && gxgy == LT:
		return RelationStarts
	case lxly == EQ && gxgy == GT:
		return RelationStartedBy
	case lxly == GT && gxgy == LT:
		return RelationDuring
	case lxly == GT && gxgy == EQ:
		return RelationFinishes
	case lxly == GT && gxgy == GT:
		return RelationOverlappedBy
	default:
		return RelationUnknown
	}
}

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
