package rampart

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
