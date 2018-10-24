package abelian_test

import (
	"testing"

	"github.com/nickng/abelian"
	"github.com/nickng/abelian/set"
	"github.com/nickng/abelian/set/prop"
)

/*
// Tests Sub function works.
func TestSub(t *testing.T) {
	g1 := abelian.New(1, op.AddN(1))
	t.Logf("Group(dimen=1): %v", g1.String())
	x1 := g1.Tuple(1)
	y1 := g1.Tuple(2)
	z1, _ := abelian.Sub(x1, y1)
	t.Logf("%v - %v = %v", x1, y1, z1)
	if want, got := abelian.Tuple(-1), z1; !abelian.Equal(want, got) {
		t.Errorf("Sub(%v, %v) expected to be %v but got %v", x1, y1, want, got)
	}

	g2 := abelian.New(2, op.AddN(2))
	t.Logf("Group(dimen=2): %s", g2.String())
	x2 := g2.Tuple(1, 2)
	y2 := g2.Tuple(2, 3)
	z2, _ := abelian.Sub(x2, y2)
	t.Logf("%v - %v = %v", x2, y2, z2)
	if want, got := abelian.Tuple(-1, -1), z2; !abelian.Equal(want, got) {
		t.Errorf("Sub(%v, %v) expected to be %v but got %v", x2, y2, want, got)
	}
}
*/

// This tests the documentation example.
func TestExample(t *testing.T) {
	s := set.NewIntTuple(1)
	plusOne := func(x, y set.Elem) set.Elem {
		// omitted: check x and y are int tuples of size 1.
		//return s.Tuple(x.(set.)[0] + y[0] + 1)
		return s.Tuple(1)
	}
	// This creates a new +1 abelian group.
	gAddOne := abelian.New(s, plusOne)
	_ = gAddOne
}

func TestStrictOrder(t *testing.T) {
	s := set.NewIntTuple(1)
	g := abelian.New(s, s.Add)
	if so, ok := g.Set.(prop.StrictOrdered); ok {
		t.Logf("%s (i.e. %s) implements < operator", g.String(), g.Set.Name())
		one := s.Tuple(1)
		two := s.Add(one, one)
		if !so.Less(one, two) {
			t.Errorf("unexpected: %s < %s", one, two)
		}
	} else {
		t.Errorf("Set %s is not strictly ordered", g.Set.Name())
	}
}
