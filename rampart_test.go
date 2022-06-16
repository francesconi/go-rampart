package rampart

import (
	"runtime"
	"testing"
	"time"
)

func TestTimeInterval(t *testing.T) {
	now := time.Now()
	a := NewIntervalFunc(now.Add(1000), now, func(t1, t2 time.Time) int { return int(t1.Sub(t2)) })
	requireEqual(t, now, a.Lesser())
}

func TestNewInterval(t *testing.T) {
	t.Run("sorts the tuple", func(t *testing.T) {
		a := NewInterval("a", "b")
		b := NewInterval("b", "a")
		requireEqual(t, a.x, b.x)
		requireEqual(t, a.y, b.y)
	})
}

func TestLesser(t *testing.T) {
	t.Run("returns the lesser element", func(t *testing.T) {
		requireEqual(t, "a", NewInterval("a", "b").Lesser())
	})
}

func TestGreater(t *testing.T) {
	t.Run("returns the greater element", func(t *testing.T) {
		requireEqual(t, "b", NewInterval("a", "b").Greater())
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("returns true when the interval is empty", func(t *testing.T) {
		requireEqual(t, true, NewInterval("a", "a").IsEmpty())
	})
	t.Run("returns false when the interval is non-empty", func(t *testing.T) {
		requireEqual(t, false, NewInterval("a", "b").IsEmpty())
	})
}

func TestIsNonEmpty(t *testing.T) {
	t.Run("returns false when the interval is empty", func(t *testing.T) {
		requireEqual(t, false, NewInterval("a", "a").IsNonEmpty())
	})
	t.Run("returns true when the interval is non-empty", func(t *testing.T) {
		requireEqual(t, true, NewInterval("a", "b").IsNonEmpty())
	})
}

func relate(x1, y1 int) func(int, int) Relation {
	return func(x2, y2 int) Relation {
		return NewInterval(x1, y1).Relate(NewInterval(x2, y2))
	}
}

func TestRelate(t *testing.T) {
	t.Run("identifies the before relation", func(t *testing.T) {
		requireEqual(t, RelationBefore, relate(1, 2)(3, 7))
	})
	t.Run("identifies the meets relation", func(t *testing.T) {
		requireEqual(t, RelationMeets, relate(2, 3)(3, 7))
	})
	t.Run("identifies the overlaps relation", func(t *testing.T) {
		requireEqual(t, RelationOverlaps, relate(2, 4)(3, 7))
	})
	t.Run("identifies the finished by relation", func(t *testing.T) {
		requireEqual(t, RelationFinishedBy, relate(2, 7)(3, 7))
	})
	t.Run("identifies the contains relation", func(t *testing.T) {
		requireEqual(t, RelationContains, relate(2, 8)(3, 7))
	})
	t.Run("identifies the starts relation", func(t *testing.T) {
		requireEqual(t, RelationStarts, relate(3, 4)(3, 7))
	})
	t.Run("identifies the equal relation", func(t *testing.T) {
		requireEqual(t, RelationEqual, relate(3, 7)(3, 7))
	})
	t.Run("identifies the started by relation", func(t *testing.T) {
		requireEqual(t, RelationStartedBy, relate(3, 8)(3, 7))
	})
	t.Run("identifies the during relation", func(t *testing.T) {
		requireEqual(t, RelationDuring, relate(4, 6)(3, 7))
	})
	t.Run("identifies the finishes relation", func(t *testing.T) {
		requireEqual(t, RelationFinishes, relate(6, 7)(3, 7))
	})
	t.Run("identifies the overlapped by relation", func(t *testing.T) {
		requireEqual(t, RelationOverlappedBy, relate(6, 8)(3, 7))
	})
	t.Run("identifies the met by relation", func(t *testing.T) {
		requireEqual(t, RelationMetBy, relate(7, 8)(3, 7))
	})
	t.Run("identifies the after relation", func(t *testing.T) {
		requireEqual(t, RelationAfter, relate(8, 9)(3, 7))
	})

	t.Run("empty left interval", func(t *testing.T) {
		t.Run("before", func(t *testing.T) {
			requireEqual(t, RelationBefore, relate(2, 2)(3, 7))
		})
		t.Run("at lesser", func(t *testing.T) {
			requireEqual(t, RelationOverlaps, relate(3, 3)(3, 7))
		})
		t.Run("during", func(t *testing.T) {
			requireEqual(t, RelationDuring, relate(5, 5)(3, 7))
		})
		t.Run("at greater", func(t *testing.T) {
			requireEqual(t, RelationOverlappedBy, relate(7, 7)(3, 7))
		})
		t.Run("after", func(t *testing.T) {
			requireEqual(t, RelationAfter, relate(8, 8)(3, 7))
		})
	})

	t.Run("empty right interval", func(t *testing.T) {
		t.Run("before", func(t *testing.T) {
			requireEqual(t, RelationAfter, relate(3, 7)(2, 2))
		})
		t.Run("at lesser", func(t *testing.T) {
			requireEqual(t, RelationOverlappedBy, relate(3, 7)(3, 3))
		})
		t.Run("during", func(t *testing.T) {
			requireEqual(t, RelationContains, relate(3, 7)(5, 5))
		})
		t.Run("at greater", func(t *testing.T) {
			requireEqual(t, RelationOverlaps, relate(3, 7)(7, 7))
		})
		t.Run("after", func(t *testing.T) {
			requireEqual(t, RelationBefore, relate(3, 7)(8, 8))
		})
	})

	t.Run("both empty intervals", func(t *testing.T) {
		t.Run("before", func(t *testing.T) {
			requireEqual(t, RelationBefore, relate(4, 4)(5, 5))
		})
		t.Run("equal", func(t *testing.T) {
			requireEqual(t, RelationEqual, relate(5, 5)(5, 5))
		})
		t.Run("after", func(t *testing.T) {
			requireEqual(t, RelationAfter, relate(6, 6)(5, 5))
		})
	})
}

func TestInvert(t *testing.T) {
	t.Run("inverts the after relation", func(t *testing.T) {
		requireEqual(t, RelationBefore, RelationAfter.Invert())
	})
	t.Run("inverts the before relation", func(t *testing.T) {
		requireEqual(t, RelationAfter, RelationBefore.Invert())
	})
	t.Run("inverts the contains relation", func(t *testing.T) {
		requireEqual(t, RelationDuring, RelationContains.Invert())
	})
	t.Run("inverts the during relation", func(t *testing.T) {
		requireEqual(t, RelationContains, RelationDuring.Invert())
	})
	t.Run("inverts the equal relation", func(t *testing.T) {
		requireEqual(t, RelationEqual, RelationEqual.Invert())
	})
	t.Run("inverts the finished by relation", func(t *testing.T) {
		requireEqual(t, RelationFinishes, RelationFinishedBy.Invert())
	})
	t.Run("inverts the finishes relation", func(t *testing.T) {
		requireEqual(t, RelationFinishedBy, RelationFinishes.Invert())
	})
	t.Run("inverts the meets relation", func(t *testing.T) {
		requireEqual(t, RelationMetBy, RelationMeets.Invert())
	})
	t.Run("inverts the met by relation", func(t *testing.T) {
		requireEqual(t, RelationMeets, RelationMetBy.Invert())
	})
	t.Run("inverts the overlapped by relation", func(t *testing.T) {
		requireEqual(t, RelationOverlaps, RelationOverlappedBy.Invert())
	})
	t.Run("inverts the overlaps relation", func(t *testing.T) {
		requireEqual(t, RelationOverlappedBy, RelationOverlaps.Invert())
	})
	t.Run("inverts the started by relation", func(t *testing.T) {
		requireEqual(t, RelationStarts, RelationStartedBy.Invert())
	})
	t.Run("inverts the starts relation", func(t *testing.T) {
		requireEqual(t, RelationStartedBy, RelationStarts.Invert())
	})
	t.Run("inverts an invalid relation", func(t *testing.T) {
		requireEqual(t, RelationUnknown, Relation(42).Invert())
	})
}

func requireEqual(t *testing.T, expected, actual any) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return
	}

	if actual != expected {
		t.Errorf("\nTrace:\t%s:%d"+
			"\nError:\texpected %v, got %v"+
			"\nTest:\t%s", file, line, actual, expected, t.Name())
	}
}
