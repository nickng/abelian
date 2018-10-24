package set_test

import (
	"fmt"

	"github.com/nickng/abelian/set"
)

func ExampleIntTupleSet_Interval() {
	// To iterate over a specified range
	// use Interval() to create a finite subset.
	s := set.NewIntTuple(2)
	iv := s.Interval(s.Tuple(1, 1), s.Tuple(2, 2))

	// Enumerate returns an iterator.
	iter := iv.Enumerate()
	for {
		n, more := iter.Next()
		fmt.Println(n)
		if !more {
			break
		}
	}
	// Output:
	// (1,1)
	// (1,2)
	// (2,1)
	// (2,2)
}

func ExampleIntTupleSet_Interval_alternative() {
	// To iterate over a specified range
	// use Range() to create a finite subset.
	s := set.NewIntTuple(2)
	iv := s.Interval(s.Tuple(1, 1), s.Tuple(2, 2))

	// Slice returns a finite instantiated slice.
	tuples := iv.Slice()
	for _, tuple := range tuples {
		fmt.Println(tuple)
	}
	// Output:
	// (1,1)
	// (1,2)
	// (2,1)
	// (2,2)
}
