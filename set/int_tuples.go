package set

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// MismatchDimErr is the type of error where an operation
// is applied to two elements with different dimensions.
type MismatchDimErr struct {
	Dim1, Dim2 int
}

func (e MismatchDimErr) Error() string {
	return fmt.Sprintf("tuple dimension mismatch: %d != %d", e.Dim1, e.Dim2)
}

// IntTupleSet is a set of Integer ℤ or tuples of Integers (ℤx...xℤ).
//
// The type represents the tuple size, e.g.
// IntTuple(1) is ℤ.
// IntTuple(2) is ℤ x ℤ.
//
type IntTupleSet int

// NewIntTuple returns a new integer tuple set with the specified tuple size.
func NewIntTuple(size int) IntTupleSet {
	return IntTupleSet(size)
}

// Tuple is a variadic function to create a tuple from v,
// a member of the set.
//
// The length of v must match tuple sizes in s, otherwise
// it throws a runtime error.
func (s IntTupleSet) Tuple(v ...int) IntTuple {
	if len(v) != s.Size() {
		log.Fatalf("cannot create tuple/%d from %v: %v", s.Size(), v, MismatchDimErr{len(v), s.Size()})
	}
	return IntTuple(v)
}

// Identity returns the identity of the set.
//
// For the set of integer, the identity is 0.
// For tuples, the identity is (0,0...).
func (s IntTupleSet) Identity() Elem {
	// Ident is the default value of []int with dimen elements (0,0...).
	return make(IntTuple, s.Size())
}

// IsIn returns true if x ∈ s.
func (s IntTupleSet) IsIn(x Elem) bool {
	xElem, ok := x.(IntTuple)
	if !ok {
		return false
	}
	return s.Size() == xElem.Size()
}

// Size returns the tuple size of the set.
func (s IntTupleSet) Size() int {
	return int(s)
}

// Name returns the formal name of the IntTuple set.
func (s IntTupleSet) Name() string {
	if s.Size() == 0 {
		return "∅"
	}
	name := make([]string, s.Size())
	for i := range name {
		name[i] = "ℤ"
	}
	return strings.Join(name, "x")
}

// Add is the + binary operation. It returns x + y.
func (s IntTupleSet) Add(x, y Elem) Elem {
	xElem, yElem := x.(IntTuple), y.(IntTuple)
	if xElem.Size() != s.Size() {
		log.Fatal(MismatchDimErr{Dim1: xElem.Size(), Dim2: s.Size()})
	}
	if yElem.Size() != s.Size() {
		log.Fatal(MismatchDimErr{Dim1: yElem.Size(), Dim2: s.Size()})
	}
	return s.add(xElem, yElem)
}

func (s IntTupleSet) add(x, y IntTuple) IntTuple {
	z := make(IntTuple, s.Size())
	for i := range z {
		z[i] = x[i] + y[i]
	}
	return z
}

// Less returns x < y.
func (s IntTupleSet) Less(x, y Elem) bool {
	return x.(IntTuple).Compare(y) < 0
}

// LessEqual returns x ≤ y.
func (s IntTupleSet) LessEqual(x, y Elem) bool {
	return x.(IntTuple).Compare(y) <= 0
}

// Interval returns a finite enumerable range.
// { a | a1 ≤ a ≤ a2 }
func (s IntTupleSet) Interval(a1, a2 Elem) Enumerable {
	return IntTupleInterval{lo: a1.(IntTuple), hi: a2.(IntTuple)}
}

// IntTupleInterval is a finite subset of IntTuple
// that can be enumerated.
type IntTupleInterval struct {
	Set
	lo, hi IntTuple
}

// IsIn returns true if x ∈ s.
func (r IntTupleInterval) IsIn(x Elem) bool {
	return r.lo.Compare(x) <= 0 && r.hi.Compare(x) >= 0
}

// Name returns the description of the subset.
func (r IntTupleInterval) Name() string {
	return fmt.Sprintf("%s≤..≤%s", r.lo, r.hi)
}

// Enumerate creates an iterator for looping over the IntTuple in the range.
func (r IntTupleInterval) Enumerate() Nexter {
	return &IntTupleIter{IntTupleInterval: r, curr: r.lo}
}

// Slice returns ordered Elem in the range as a slice.
func (r IntTupleInterval) Slice() []Elem {
	var s []Elem
	e := r.Enumerate()
	for {
		next, more := e.Next()
		if !more {
			s = append(s, next)
			break
		}
		s = append(s, next)
	}
	return s
}

// IntTupleIter is a IntTuple iterator.
type IntTupleIter struct {
	IntTupleInterval
	curr IntTuple
}

func (n *IntTupleIter) next(curr IntTuple) IntTuple {
	next := make(IntTuple, curr.Size())
	copy(next, curr)
	carry := 1
	for i := curr.Size() - 1; i >= 0; i-- {
		if next[i]+carry > n.hi[i] {
			next[i] = n.lo[i]
		} else {
			next[i] += carry
			carry = 0
		}
	}
	// If overflow, use max intTuple in range.
	if carry == 1 {
		copy(next, n.hi)
	}
	return next
}

// Next returns the next Elem in the range, and indicates
// if there are more elements in the range with more.
func (n *IntTupleIter) Next() (next Elem, more bool) {
	next = n.curr
	n.curr = n.next(n.curr)
	more = next.Compare(n.curr) != 0
	return next, more
}

// IntTuple is an Elem in a IntTupleSet.
type IntTuple []int

// Size returns the tuple size of e.
func (e IntTuple) Size() int {
	return len(e)
}

// String returns a numeric/tuple representation set element e.
func (e IntTuple) String() string {
	if e.Size() == 1 {
		return strconv.Itoa(e[0]) // integer
	}

	// e.Size() == 0: ()
	// e.Size() >= 2: (e[0],e[1],...)
	var buf bytes.Buffer
	buf.WriteRune('(')
	for i := range e {
		if i != 0 {
			buf.WriteRune(',')
		}
		buf.WriteString(fmt.Sprintf("%d", e[i]))
	}
	buf.WriteRune(')')
	return buf.String()
}

// Compare returns 0 if e == x, -ve int if e < x, +ve int if e > x.
func (e IntTuple) Compare(x Elem) int {
	tuple := x.(IntTuple)
	for i := 0; i < e.Size(); i++ {
		if e[i] < tuple[i] {
			return -1
		} else if e[i] > tuple[i] {
			return 1
		}
	}
	return 0
}
