package structslice

import (
	"fmt"
	"reflect"
)

// Comparer is an interface for types that can compare themselves to each other.
type Comparer interface {
	Less(Comparer) bool
}

// Stringer is an interface for types that have a string representation
type Stringer interface {
	String() string
}

// SortByName sorts the slice of structs by the field name given by 'n'
func SortByName(v interface{}, n string) error {
	s := attach(v)
	fmt.Println("s.Value type is", reflect.TypeOf(s.Value.Interface()))
	f, ok := s.Value.Index(0).Type().FieldByName(n)

	if !ok {
		return fmt.Errorf(errFieldFmt, n, s.Value.Index(0).Type())
	}

	sortByIndex(v, f.Index[0])
	return nil
}

// SortStableByName is like SortByName, except it performs a stable sort.
// Because it performs a stable sort, it accepts a variadic number of sort
// keys. Sorting is done for every key in the order that the key is passed in
// to the function.
func SortStableByName(v interface{}, n ...string) error {
	if len(n) == 0 {
		return nil
	}
	s := attach(v)

	keys := make([]int, len(n))
	for i, v := range n {
		f, ok := s.Value.Index(0).Type().FieldByName(v)
		if !ok {
			return fmt.Errorf(errFieldFmt, v, s.Value.Index(0).Type())
		}
		keys[i] = f.Index[0]
	}

	sortStableByIndex(v, keys...)

	return nil
}
