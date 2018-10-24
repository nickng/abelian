# abelian [![Build Status](https://travis-ci.org/nickng/abelian.svg?branch=master)](https://travis-ci.org/nickng/abelian) [![GoDoc](https://godoc.org/github.com/nickng/abelian?status.svg)](http://godoc.org/github.com/nickng/abelian)

Implementation of integer-based
[torsion-free Abelian Group](https://en.wikipedia.org/wiki/Torsion-free_abelian_group)
in Go.

## Usage

    go get github.com/nickng/abelian

Example: creating a group

```go
// This example shows how to create an Integer-pair abelian group.
s := set.NewIntTuple(2) // tuple size set to 2
g := abelian.New(s, s.Add)

// The following g.Op is the + operation for pairs,
// hence the output is (1,2) + (2,3) = (3,5)
//
output := g.Op(s.Tuple(1, 2), s.Tuple(2, 3))
fmt.Println("Group:", g.String())
fmt.Println(output)

// prints out (3,5)
```

Example: enumerating a finite subset

```go
// This example shows how to enumerate a 2D-interval.
s := set.NewIntTuple(2)
g := abelian.New(s, s.Add)

// The given Set s implements
// - Less(x, y Elem) bool
// - Interval(lo, hi Elem) Enumerable
subset := g.Set.(prop.PartialOrdered).Interval(s.Tuple(1, 1), s.Tuple(2, 2))
for i, tuple := range subset.Slice() {
    fmt.Println(tuple)
}

// prints out (1,1) (1,2) (2,1) (2,2)
```

More examples can be found in [GoDoc](https://godoc.org/github.com/nickng/abelian).

## License

  abelian is licensed under the [Apache License](http://www.apache.org/licenses/LICENSE-2.0)
