// Package abelian provides data structures and functions abelian group in Go.
//
// Abelian Group
//
// An abelian group is a mathematical group 〈S,·〉 in which the binary
// operation · can be applied to two group elements x,y ∈ S commutatively,
// i.e. x·y = y·x.
//
// Usage
//
// This package provides a minimal representation for such group, where the
// set S is any types that implements the set.Set interface. Most of the
// functionalities of the Group will be provided by the set.Set implementation.
// A number of properties interfaces (e.g. StrictOrdered, Enumerable, ...) are
// also provided for convenience.
//
//   s := set.NewIntTuple(2)
//   g := abelian.New(s, s.Add) // creates 〈ℤxℤ, +〉 group
//
//   // The < operator is defined for this group,
//   // so this group is strictly ordered.
//   if so, ok := g.Set.(prop.StrictOrdered); ok {
//   	x, y := s.Tuple(1, 2), s.Tuple(2, 3)
//   	so.Less(x, y)
//   }
//
// Example
//
// To use this package to represent custom abelian groups, simply
// provide an implementation for the Set and binary function Op.
//
//   s := set.NewIntTuple(1) // s implements set.Set
//   plusOne := func(x, y set.Elem) set.Elem { ... }
//   // This creates a new plusOne abelian group.
//   gAddOne := abelian.New(s, plusOne)
//
package abelian

import (
	"fmt"

	"github.com/nickng/abelian/set"
)

// Group is a generic abelian group: 〈S, op〉.
// S is the (possibly infinite) set and op is the binary
// operation that can be applied to elements of S to obtain
// another element of S.
type Group struct {
	// Set is an abstract representation of the set S
	// that forms the group.
	//
	// Set can be infinite and may not be concretely populated.
	set.Set

	// Op is a binary operation on elements of the group.
	Op set.BinOp
}

// String returns a formal string representation of the group.
func (g Group) String() string {
	return fmt.Sprintf("〈%s, %T, %s〉", g.Set.Name(), g.Op, g.Set.Identity())
}

// New return a new instance of abelian group.
func New(s set.Set, op set.BinOp) Group {
	return Group{Set: s, Op: op}
}
