package abelian_test

import (
	"fmt"

	"github.com/nickng/abelian"
	"github.com/nickng/abelian/set"
	"github.com/nickng/abelian/set/prop"
)

func ExampleNew_integer() {
	// This example shows how to create an Integer abelian group.
	s := set.NewIntTuple(1)
	g := abelian.New(s, s.Add)

	// IntTupleSet.Tuple(i) converts the integer i into member of IntTupleSet(1).
	//
	// The following g.Op is the + operation for integers,
	// hence the output is 1 + 2 = 3
	//
	output := g.Op(s.Tuple(1), s.Tuple(2))
	fmt.Println("Group:", g.String())
	fmt.Println(output)
	// Output:
	// Group: 〈ℤ, set.BinOp, 0〉
	// 3
}

func ExampleNew_pair() {
	// This example shows how to create an Integer-pair abelian group.
	s := set.NewIntTuple(2) // tuple size set to 2
	g := abelian.New(s, s.Add)

	// IntTupleSet.Tuple(i, j) converts the integers i, j into a member
	// of IntTupleSet(2).
	//
	// The following g.Op is the + operation for pairs,
	// hence the output is (1,2) + (2,3) = (3,5)
	//
	output := g.Op(s.Tuple(1, 2), s.Tuple(2, 3))
	fmt.Println("Group:", g.String())
	fmt.Println(output)
	// Output:
	// Group: 〈ℤxℤ, set.BinOp, (0,0)〉
	// (3,5)
}

func ExampleNew_triple() {
	// This example shows how to create an Integer-triple abelian group.
	s := set.NewIntTuple(3) // tuple size set to 3
	g := abelian.New(s, s.Add)

	// IntTupleSet.Tuple(i, j, k) converts the integers i, j, k
	// into a member of IntTupleSet(3).
	//
	// The following g.Op is the + operation for triples,
	// hence the output is (1,2,3) + (3,4,5) = (4,6,8)
	//
	output := g.Op(s.Tuple(1, 2, 3), s.Tuple(3, 4, 5))
	fmt.Println("Group:", g.String())
	fmt.Println(output)
	// Output:
	// Group: 〈ℤxℤxℤ, set.BinOp, (0,0,0)〉
	// (4,6,8)
}

func ExampleNew_enumerate() {
	// This example shows how to enumerate a 2D-interval.
	s := set.NewIntTuple(2)
	g := abelian.New(s, s.Add)
	tuples := g.Set.(prop.PartialOrdered).Interval(s.Tuple(1, 1), s.Tuple(2, 2)).Slice()
	for i, tuple := range tuples {
		fmt.Println(i, tuple)
	}
	// Output:
	// 0 (1,1)
	// 1 (1,2)
	// 2 (2,1)
	// 3 (2,2)

}
