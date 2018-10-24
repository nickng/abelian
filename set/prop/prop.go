// Package prop are properties of sets.
package prop

import "github.com/nickng/abelian/set"

// PartialOrdered is the property where the
// elements is partially ordered (≤ defined).
type PartialOrdered interface {
	LessEqual(x, y set.Elem) bool

	// If a set is partial ordered an interval
	// can be specified for enumeration.
	Interval(lo, hi set.Elem) set.Enumerable
}

// StrictOrdered is the property where the
// elements is strictly ordered (< defined).
type StrictOrdered interface {
	Less(x, y set.Elem) bool
}
