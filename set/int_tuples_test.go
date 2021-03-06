package set

import (
	"testing"
)

// Tests the Tuple(...) constructor is same as intTuple{...}.
func TestIntTupleTuple(t *testing.T) {
	// 0D group - unit
	s0 := NewIntTuple(0)
	x0 := s0.Tuple()
	if want, got := 0, len(x0); want != got {
		t.Errorf("expected length to be %d got %d", want, got)
		t.FailNow()
	}
	if want, got := 0, x0.Size(); want != got {
		t.Errorf("expected length to be %d got %d", want, got)
		t.FailNow()
	}
	if !s0.IsIn(x0) {
		t.Errorf("s.Tuple()=%v should be member of s=%v", x0, s0.Name())
	}

	// 1D group
	s1 := NewIntTuple(1)
	x1, y1 := s1.Tuple(1), IntTuple{1}
	if want, got := 1, len(x1); want != got {
		t.Errorf("expected length to be %d got %d", want, got)
		t.FailNow()
	}
	if want, got := 1, len(y1); want != got {
		t.Errorf("expected length to be %d got %d", want, got)
		t.FailNow()
	}
	for i := range x1 {
		if x1[i] != y1[i] {
			t.Errorf("Tuple(%v) != intTuple{%v}", x1, y1)
		}
	}
	if !s1.IsIn(x1) {
		t.Errorf("s.Tuple()=%v should be member of s=%s", x1, s1.Name())
	}

	// 2D group
	s2 := NewIntTuple(2)
	x2, y2 := s2.Tuple(1, 2), IntTuple{1, 2}
	if want, got := 2, len(x2); want != got {
		t.Errorf("expected length to be %d got %d", want, got)
		t.FailNow()
	}
	if want, got := 2, len(y2); want != got {
		t.Errorf("expected length to be %d got %d", want, got)
		t.FailNow()
	}
	for i := range x2 {
		if x2[i] != y2[i] {
			t.Errorf("Tuple(%v) != Elem{%v}", x1, y1)
		}
	}
	if !s2.IsIn(x2) {
		t.Errorf("s.Tuple()=%v should be member of s=%v", x2, s2.Name())
	}
}

// Tests Add function works.
func TestIntTupleAdd(t *testing.T) {
	s1 := NewIntTuple(1)
	x1 := s1.Tuple(1)
	y1 := s1.Tuple(2)
	z1 := s1.Add(x1, y1)
	t.Logf("%v + %v = %v", x1, y1, z1)
	if want, got := s1.Tuple(3), z1; want.Compare(got) != 0 {
		t.Errorf("Add(%v, %v) expected to be %v but got %v", x1, y1, want, got)
	}

	s2 := NewIntTuple(2)
	x2 := s2.Tuple(1, 2)
	y2 := s2.Tuple(2, 3)
	z2 := s2.Add(x2, y2)
	t.Logf("%v + %v = %v", x2, y2, z2)
	if want, got := s2.Tuple(3, 5), z2; want.Compare(got) != 0 {
		t.Errorf("Add(%v, %v) expected to be %v but got %v", x2, y2, want, got)
	}
}

func TestEnumerate(t *testing.T) {
	check := func(want []IntTuple, slice []Elem, iter Nexter) {
		for i := 0; ; i++ {
			n, more := iter.Next()
			if !more {
				if n.Compare(want[i]) != 0 {
					t.Errorf("expected %s but got %s", want[i], n)
				}
				if i != len(want)-1 {
					t.Errorf("expected %d results but got %d", len(want), i+1)
				}
				break
			}
			if n.Compare(want[i]) != 0 {
				t.Errorf("expected %s but got %s", want[i], n)
			}
		}
		for i, n := range slice {
			if n.Compare(want[i]) != 0 {
				t.Errorf("expected %s but got %s", want[i], n)
			}
		}
		if w, g := len(want), len(slice); w != g {
			t.Errorf("expected %d Elems but got %d", w, g)
		}
	}
	t.Run("Int", func(t *testing.T) {
		s := NewIntTuple(1)
		iv := s.Interval(s.Tuple(0), s.Tuple(3))
		e := iv.Enumerate()
		want := []IntTuple{IntTuple{0}, IntTuple{1}, IntTuple{2}, IntTuple{3}}
		check(want, iv.Slice(), e)
	})
	t.Run("2D", func(t *testing.T) {
		s := NewIntTuple(2)
		iv := s.Interval(s.Tuple(1, 1), s.Tuple(2, 2))
		e := iv.Enumerate()
		want := []IntTuple{
			IntTuple{1, 1},
			IntTuple{1, 2},
			IntTuple{2, 1},
			IntTuple{2, 2},
		}
		check(want, iv.Slice(), e)
	})
	t.Run("3D", func(t *testing.T) {
		s := NewIntTuple(3)
		iv := s.Interval(s.Tuple(0, 0, 0), s.Tuple(1, 1, 1))
		e := iv.Enumerate()
		want := []IntTuple{
			IntTuple{0, 0, 0},
			IntTuple{0, 0, 1},
			IntTuple{0, 1, 0},
			IntTuple{0, 1, 1},
			IntTuple{1, 0, 0},
			IntTuple{1, 0, 1},
			IntTuple{1, 1, 0},
			IntTuple{1, 1, 1},
		}
		check(want, iv.Slice(), e)
	})
	t.Run("Single-Item interval", func(t *testing.T) {
		s := NewIntTuple(1)
		iv := s.Interval(s.Tuple(1), s.Tuple(1))
		e := iv.Enumerate()
		want := []IntTuple{IntTuple{1}}
		check(want, iv.Slice(), e)
	})
	t.Run("Non-(i,i) coordinate", func(t *testing.T) {
		s := NewIntTuple(2)
		iv := s.Interval(s.Tuple(0, 1), s.Tuple(3, 2))
		e := iv.Enumerate()
		want := []IntTuple{
			IntTuple{0, 1},
			IntTuple{0, 2},
			IntTuple{1, 1},
			IntTuple{1, 2},
			IntTuple{2, 1},
			IntTuple{2, 2},
			IntTuple{3, 1},
			IntTuple{3, 2},
		}
		check(want, iv.Slice(), e)
	})
}

func TestComparision(t *testing.T) {
	s := NewIntTuple(1)
	one, two := s.Tuple(1), s.Tuple(2)
	if !s.Less(one, two) {
		t.Errorf("expecting Less(%v,%v) = true", one, two)
	}
	if s.Less(two, one) {
		t.Errorf("expecting Less(%v,%v) = false", two, one)
	}
	if !s.Lt(one, two) {
		t.Errorf("expecting Lt(%v,%v) = true", one, two)
	}
	if s.Lt(two, one) {
		t.Errorf("expecting Lt(%v,%v) = false", two, one)
	}
	if !s.Gt(two, one) {
		t.Errorf("expecting Gt(%v,%v) = false", two, one)
	}
	if s.Gt(one, two) {
		t.Errorf("expecting Gt(%v,%v) = true", one, two)
	}
	if !s.LessEqual(one, two) {
		t.Errorf("expecting LessEqual(%v,%v) = true", one, two)
	}
	if !s.LessEqual(one, one) {
		t.Errorf("expecting LessEqual(%v,%v) = true", one, one)
	}
	if !s.Le(one, two) {
		t.Errorf("expecting Le(%v,%v) = true", one, two)
	}
	if !s.Le(one, one) {
		t.Errorf("expecting Le(%v,%v) = true", one, one)
	}
	if s.Ge(one, two) {
		t.Errorf("expecting Ge(%v,%v) = true", one, two)
	}
	if !s.Ge(one, one) {
		t.Errorf("expecting Ge(%v,%v) = true", one, one)
	}
	if !s.Equal(one, one) {
		t.Errorf("expecting Equal(%v,%v) = true", one, one)
	}
	if s.Equal(one, two) {
		t.Errorf("expecting Equal(%v,%v) = false", one, two)
	}
	if !s.Eq(one, one) {
		t.Errorf("expecting Eq(%v,%v) = true", one, one)
	}
	if s.Eq(one, two) {
		t.Errorf("expecting Eq(%v,%v) = false", one, two)
	}
}

func TestIntervalSubset(t *testing.T) {
	s := NewIntTuple(2)
	iv := s.Interval(s.Tuple(1, 2), s.Tuple(3, 2))
	subset, ok := iv.(Set)
	if !ok {
		t.Errorf("Interval should be a set too")
		t.FailNow()
	}
	t.Logf("subset %s: Identity=%s", subset.Name(), subset.Identity().String())
	if low := s.Tuple(1, 1); subset.IsIn(low) {
		t.Errorf("%s should not be in the interval %s", low.String(), subset.Name())
	}
	if v := s.Tuple(1, 2); !subset.IsIn(v) {
		t.Errorf("%s should be in the interval %s", v.String(), subset.Name())
	}
	if v := s.Tuple(2, 2); !subset.IsIn(v) {
		t.Errorf("%s should be in the interval %s", v.String(), subset.Name())
	}
	if v := s.Tuple(3, 2); !subset.IsIn(v) {
		t.Errorf("%s should be in the interval %s", v.String(), subset.Name())
	}
	if v := s.Tuple(4, 2); subset.IsIn(v) {
		t.Errorf("%s should not be in the interval %s", v.String(), subset.Name())
	}
	if v := s.Tuple(3, 3); subset.IsIn(v) {
		t.Errorf("%s should not be in the interval %s", v.String(), subset.Name())
	}
}
