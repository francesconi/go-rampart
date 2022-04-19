package rampart

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewInterval(t *testing.T) {
	t.Run("sorts the tuple", func(t *testing.T) {
		a := NewOrdInterval("a", "b")
		b := NewOrdInterval("b", "a")
		require.Equal(t, a.x, b.x)
		require.Equal(t, a.y, b.y)
	})
}

func TestLesser(t *testing.T) {
	t.Run("returns the lesser element", func(t *testing.T) {
		require.Equal(t, "a", NewOrdInterval("a", "b").Lesser())
	})
}

func TestGreater(t *testing.T) {
	t.Run("returns the greater element", func(t *testing.T) {
		require.Equal(t, "b", NewOrdInterval("a", "b").Greater())
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("returns true when the interval is empty", func(t *testing.T) {
		require.Equal(t, true, NewOrdInterval("a", "a").IsEmpty())
	})
	t.Run("returns false when the interval is non-empty", func(t *testing.T) {
		require.Equal(t, false, NewOrdInterval("a", "b").IsEmpty())
	})
}

func TestIsNonEmpty(t *testing.T) {
	t.Run("returns false when the interval is empty", func(t *testing.T) {
		require.Equal(t, false, NewOrdInterval("a", "a").IsNonEmpty())
	})
	t.Run("returns true when the interval is non-empty", func(t *testing.T) {
		require.Equal(t, true, NewOrdInterval("a", "b").IsNonEmpty())
	})
}

func relate(x1, y1 int) func(int, int) Relation {
	return func(x2, y2 int) Relation {
		return NewOrdInterval(x1, y1).Relate(NewOrdInterval(x2, y2))
	}
}

func TestRelate(t *testing.T) {
	t.Run("identifies the before relation", func(t *testing.T) {
		require.Equal(t, RelationBefore, relate(1, 2)(3, 7))
	})
	t.Run("identifies the before relation", func(t *testing.T) {
		require.Equal(t, RelationMeets, relate(2, 3)(3, 7))
	})
	t.Run("identifies the overlaps relation", func(t *testing.T) {
		require.Equal(t, RelationOverlaps, relate(2, 4)(3, 7))
	})
	t.Run("identifies the finished by relation", func(t *testing.T) {
		require.Equal(t, RelationFinishedBy, relate(2, 7)(3, 7))
	})
	t.Run("identifies the contains relation", func(t *testing.T) {
		require.Equal(t, RelationContains, relate(2, 8)(3, 7))
	})
	t.Run("identifies the starts relation", func(t *testing.T) {
		require.Equal(t, RelationStarts, relate(3, 4)(3, 7))
	})
	t.Run("identifies the equal relation", func(t *testing.T) {
		require.Equal(t, RelationEqual, relate(3, 7)(3, 7))
	})
	t.Run("identifies the started by relation", func(t *testing.T) {
		require.Equal(t, RelationStartedBy, relate(3, 8)(3, 7))
	})
	t.Run("identifies the during relation", func(t *testing.T) {
		require.Equal(t, RelationDuring, relate(4, 6)(3, 7))
	})
	t.Run("identifies the finishes relation", func(t *testing.T) {
		require.Equal(t, RelationFinishes, relate(6, 7)(3, 7))
	})
	t.Run("identifies the overlapped by relation", func(t *testing.T) {
		require.Equal(t, RelationOverlappedBy, relate(6, 8)(3, 7))
	})
	t.Run("identifies the met by relation", func(t *testing.T) {
		require.Equal(t, RelationMetBy, relate(7, 8)(3, 7))
	})
	t.Run("identifies the after relation", func(t *testing.T) {
		require.Equal(t, RelationAfter, relate(8, 9)(3, 7))
	})

	t.Run("empty left interval", func(t *testing.T) {
		t.Run("before", func(t *testing.T) {
			t.Run("before", func(t *testing.T) {
				require.Equal(t, RelationBefore, relate(2, 2)(3, 7))
			})
			t.Run("at lesser", func(t *testing.T) {
				require.Equal(t, RelationOverlaps, relate(3, 3)(3, 7))
			})
			t.Run("during", func(t *testing.T) {
				require.Equal(t, RelationDuring, relate(5, 5)(3, 7))
			})
			t.Run("at greater", func(t *testing.T) {
				require.Equal(t, RelationOverlappedBy, relate(7, 7)(3, 7))
			})
			t.Run("after", func(t *testing.T) {
				require.Equal(t, RelationAfter, relate(8, 8)(3, 7))
			})
		})
	})

	t.Run("empty right interval", func(t *testing.T) {
		t.Run("before", func(t *testing.T) {
			require.Equal(t, RelationAfter, relate(3, 7)(2, 2))
		})
		t.Run("at lesser", func(t *testing.T) {
			require.Equal(t, RelationOverlappedBy, relate(3, 7)(3, 3))
		})
		t.Run("during", func(t *testing.T) {
			require.Equal(t, RelationContains, relate(3, 7)(5, 5))
		})
		t.Run("at greater", func(t *testing.T) {
			require.Equal(t, RelationOverlaps, relate(3, 7)(7, 7))
		})
		t.Run("after", func(t *testing.T) {
			require.Equal(t, RelationBefore, relate(3, 7)(8, 8))
		})
	})

	t.Run("both empty intervals", func(t *testing.T) {
		t.Run("before", func(t *testing.T) {
			require.Equal(t, RelationBefore, relate(4, 4)(5, 5))
		})
		t.Run("equal", func(t *testing.T) {
			require.Equal(t, RelationEqual, relate(5, 5)(5, 5))
		})
		t.Run("after", func(t *testing.T) {
			require.Equal(t, RelationAfter, relate(6, 6)(5, 5))
		})
	})
}

func TestInvert(t *testing.T) {
	t.Run("inverts the after relation", func(t *testing.T) {
		require.Equal(t, RelationBefore, RelationAfter.Invert())
	})
	t.Run("inverts the before relation", func(t *testing.T) {
		require.Equal(t, RelationAfter, RelationBefore.Invert())
	})
	t.Run("inverts the contains relation", func(t *testing.T) {
		require.Equal(t, RelationDuring, RelationContains.Invert())
	})
	t.Run("inverts the during relation", func(t *testing.T) {
		require.Equal(t, RelationContains, RelationDuring.Invert())
	})
	t.Run("inverts the equal relation", func(t *testing.T) {
		require.Equal(t, RelationEqual, RelationEqual.Invert())
	})
	t.Run("inverts the finished by relation", func(t *testing.T) {
		require.Equal(t, RelationFinishes, RelationFinishedBy.Invert())
	})
	t.Run("inverts the finishes relation", func(t *testing.T) {
		require.Equal(t, RelationFinishedBy, RelationFinishes.Invert())
	})
	t.Run("inverts the meets relation", func(t *testing.T) {
		require.Equal(t, RelationMetBy, RelationMeets.Invert())
	})
	t.Run("inverts the met by relation", func(t *testing.T) {
		require.Equal(t, RelationMeets, RelationMetBy.Invert())
	})
	t.Run("inverts the overlapped by relation", func(t *testing.T) {
		require.Equal(t, RelationOverlaps, RelationOverlappedBy.Invert())
	})
	t.Run("inverts the overlaps relation", func(t *testing.T) {
		require.Equal(t, RelationOverlappedBy, RelationOverlaps.Invert())
	})
	t.Run("inverts the started by relation", func(t *testing.T) {
		require.Equal(t, RelationStarts, RelationStartedBy.Invert())
	})
	t.Run("inverts the starts relation", func(t *testing.T) {
		require.Equal(t, RelationStartedBy, RelationStarts.Invert())
	})
	t.Run("inverts an invalid relation", func(t *testing.T) {
		require.Equal(t, RelationUnknown, Relation(42).Invert())
	})
}
