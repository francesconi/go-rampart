package rampart

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewInterval(t *testing.T) {
	t.Run("sorts the tuple", func(t *testing.T) {
		require.Equal(t, NewInterval("a", "b"), NewInterval("b", "a"))
	})
}

func TestLesser(t *testing.T) {
	t.Run("returns the lesser element", func(t *testing.T) {
		require.Equal(t, "a", NewInterval("a", "b").Lesser())
	})
}

func TestGreater(t *testing.T) {
	t.Run("returns the greater element", func(t *testing.T) {
		require.Equal(t, "b", NewInterval("a", "b").Greater())
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("returns true when the interval is empty", func(t *testing.T) {
		require.Equal(t, true, NewInterval("a", "a").IsEmpty())
	})
	t.Run("returns false when the interval is non-empty", func(t *testing.T) {
		require.Equal(t, false, NewInterval("a", "b").IsEmpty())
	})
}

func TestIsNonEmpty(t *testing.T) {
	t.Run("returns false when the interval is empty", func(t *testing.T) {
		require.Equal(t, false, NewInterval("a", "a").IsNonEmpty())
	})
	t.Run("returns true when the interval is non-empty", func(t *testing.T) {
		require.Equal(t, true, NewInterval("a", "b").IsNonEmpty())
	})
}

func relate(x1, y1 int) func(int, int) Relation {
	return func(x2, y2 int) Relation {
		return NewInterval(x1, y1).Relate(NewInterval(x2, y2))
	}
}

func TestRelate(t *testing.T) {
	t.Run("identifies the before relation", func(t *testing.T) {
		require.Equal(t, relate(1, 2)(3, 7), RelationBefore)
	})
	t.Run("identifies the before relation", func(t *testing.T) {
		require.Equal(t, relate(2, 3)(3, 7), RelationMeets)
	})
	t.Run("identifies the overlaps relation", func(t *testing.T) {
		require.Equal(t, relate(2, 4)(3, 7), RelationOverlaps)
	})
	t.Run("identifies the finished by relation", func(t *testing.T) {
		require.Equal(t, relate(2, 7)(3, 7), RelationFinishedBy)
	})
	t.Run("identifies the contains relation", func(t *testing.T) {
		require.Equal(t, relate(2, 8)(3, 7), RelationContains)
	})
	t.Run("identifies the starts relation", func(t *testing.T) {
		require.Equal(t, relate(3, 4)(3, 7), RelationStarts)
	})
	t.Run("identifies the equal relation", func(t *testing.T) {
		require.Equal(t, relate(3, 7)(3, 7), RelationEqual)
	})
	t.Run("identifies the started by relation", func(t *testing.T) {
		require.Equal(t, relate(3, 8)(3, 7), RelationStartedBy)
	})
	t.Run("identifies the during relation", func(t *testing.T) {
		require.Equal(t, relate(4, 6)(3, 7), RelationDuring)
	})
	t.Run("identifies the finishes relation", func(t *testing.T) {
		require.Equal(t, relate(6, 7)(3, 7), RelationFinishes)
	})
	t.Run("identifies the overlapped by relation", func(t *testing.T) {
		require.Equal(t, relate(6, 8)(3, 7), RelationOverlappedBy)
	})
	t.Run("identifies the met by relation", func(t *testing.T) {
		require.Equal(t, relate(7, 8)(3, 7), RelationMetBy)
	})
	t.Run("identifies the after relation", func(t *testing.T) {
		require.Equal(t, relate(8, 9)(3, 7), RelationAfter)
	})

	t.Run("empty left interval", func(t *testing.T) {
		t.Run("before", func(t *testing.T) {
			t.Run("before", func(t *testing.T) {
				require.Equal(t, relate(2, 2)(3, 7), RelationBefore)
			})
			t.Run("at lesser", func(t *testing.T) {
				require.Equal(t, relate(3, 3)(3, 7), RelationOverlaps)
			})
			t.Run("during", func(t *testing.T) {
				require.Equal(t, relate(5, 5)(3, 7), RelationDuring)
			})
			t.Run("at greater", func(t *testing.T) {
				require.Equal(t, relate(7, 7)(3, 7), RelationOverlappedBy)
			})
			t.Run("after", func(t *testing.T) {
				require.Equal(t, relate(8, 8)(3, 7), RelationAfter)
			})
		})
	})

	t.Run("empty right interval", func(t *testing.T) {
		t.Run("before", func(t *testing.T) {
			require.Equal(t, relate(3, 7)(2, 2), RelationAfter)
		})
		t.Run("at lesser", func(t *testing.T) {
			require.Equal(t, relate(3, 7)(3, 3), RelationOverlappedBy)
		})
		t.Run("during", func(t *testing.T) {
			require.Equal(t, relate(3, 7)(5, 5), RelationContains)
		})
		t.Run("at greater", func(t *testing.T) {
			require.Equal(t, relate(3, 7)(7, 7), RelationOverlaps)
		})
		t.Run("after", func(t *testing.T) {
			require.Equal(t, relate(3, 7)(8, 8), RelationBefore)
		})
	})

	t.Run("both empty intervals", func(t *testing.T) {
		t.Run("before", func(t *testing.T) {
			require.Equal(t, relate(4, 4)(5, 5), RelationBefore)
		})
		t.Run("equal", func(t *testing.T) {
			require.Equal(t, relate(5, 5)(5, 5), RelationEqual)
		})
		t.Run("after", func(t *testing.T) {
			require.Equal(t, relate(6, 6)(5, 5), RelationAfter)
		})
	})
}

func TestInvert(t *testing.T) {
	t.Run("inverts the after relation", func(t *testing.T) {
		require.Equal(t, RelationAfter.Invert(), RelationBefore)
	})
	t.Run("inverts the before relation", func(t *testing.T) {
		require.Equal(t, RelationBefore.Invert(), RelationAfter)
	})
	t.Run("inverts the contains relation", func(t *testing.T) {
		require.Equal(t, RelationContains.Invert(), RelationDuring)
	})
	t.Run("inverts the during relation", func(t *testing.T) {
		require.Equal(t, RelationDuring.Invert(), RelationContains)
	})
	t.Run("inverts the equal relation", func(t *testing.T) {
		require.Equal(t, RelationEqual.Invert(), RelationEqual)
	})
	t.Run("inverts the finished by relation", func(t *testing.T) {
		require.Equal(t, RelationFinishedBy.Invert(), RelationFinishes)
	})
	t.Run("inverts the finishes relation", func(t *testing.T) {
		require.Equal(t, RelationFinishes.Invert(), RelationFinishedBy)
	})
	t.Run("inverts the meets relation", func(t *testing.T) {
		require.Equal(t, RelationMeets.Invert(), RelationMetBy)
	})
	t.Run("inverts the met by relation", func(t *testing.T) {
		require.Equal(t, RelationMetBy.Invert(), RelationMeets)
	})
	t.Run("inverts the overlapped by relation", func(t *testing.T) {
		require.Equal(t, RelationOverlappedBy.Invert(), RelationOverlaps)
	})
	t.Run("inverts the overlaps relation", func(t *testing.T) {
		require.Equal(t, RelationOverlaps.Invert(), RelationOverlappedBy)
	})
	t.Run("inverts the started by relation", func(t *testing.T) {
		require.Equal(t, RelationStartedBy.Invert(), RelationStarts)
	})
	t.Run("inverts the starts relation", func(t *testing.T) {
		require.Equal(t, RelationStarts.Invert(), RelationStartedBy)
	})
	t.Run("inverts an invalid relation", func(t *testing.T) {
		require.Equal(t, Relation(42).Invert(), RelationUnknown)
	})
}
