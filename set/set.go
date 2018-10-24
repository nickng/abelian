// Package set implements a Set data structure.
package set

// Set is a generic set.
type Set interface {
	// IsIn tests if the Elem x is a member of the set.
	IsIn(x Elem) bool

	// Name returns a name for the set.
	Name() string

	// Identity returns the identity Elem of the set.
	Identity() Elem
}

// BinOp is a binary operation between two set Elems.
type BinOp func(Elem, Elem) Elem

// Elem represents an element of a Set.
type Elem interface {
	String() string
	Compare(x Elem) int
}

// Enumerater allows ranging over a set of Elems as an iterator.
type Enumerater interface {
	Enumerate() Nexter
}

// Slicer allows ranging over a set of Elems as a slice.
type Slicer interface {
	Slice() []Elem
}

// Enumerable is an interface for ranging over a set of Elem.
type Enumerable interface {
	Enumerater
	Slicer
}

// Nexter is implemented by object or data structure
// which may have a Next method.
type Nexter interface {
	Next() (next Elem, more bool)
}
