package structslice

import (
	"fmt"
	"reflect"
	"sort"
)

// structSlice is a control structure for implementing abstract operations
// on a slice of structs.
type structSlice struct {
	Value          reflect.Value
	SortFieldIndex int
}

// SortByName sorts the slice of structs by the field name given by 'n'
func SortByName(v interface{}, n string) {
	s := attach(v)
	f, ok := s.Value.Index(0).Type().FieldByName(n)
	if !ok {
		return
	}

	SortByIndex(v, f.Index[0])
}

// SortByIndex sorts the slice of structs by the field index 'i'
func SortByIndex(v interface{}, i int) {
	s := attach(v)
	s.SortFieldIndex = i
	sort.Sort(s)
}

// SortStableByName is like SortByName, except it performs a stable sort.
// Because it performs a stable sort, it accepts a variadic number of sort
// keys. Sorting is done for every key in the order that the key is passed in
// to the function.
func SortStableByName(v interface{}, n ...string) {
	if len(n) == 0 {
		return
	}
	s := attach(v)

	keys := make([]int, len(n))
	for i, v := range n {
		f, ok := s.Value.Index(0).Type().FieldByName(v)
		if !ok {
			return
		}
		keys[i] = f.Index[0]
	}

	SortStableByIndex(v, keys...)
}

// SortStableByIndex is like SortByIndex, except it performs a stable sort.
// Because it performs a stable sort, it accepts a variadic number of sort
// keys. Sorting is done for every key in the order that the key is passed in
// to the function.
func SortStableByIndex(v interface{}, i ...int) {
	if len(i) == 0 {
		return
	}

	s := attach(v)
	for _, v := range i {
		s.SortFieldIndex = v
		sort.Stable(s)
	}
}

// attach binds to the slice of structs and returns a structSlice object
// for executing sorting operations on the slice elements. Attach panics
// if the underlying interface, v, is not a slice of structs.
func attach(v interface{}) *structSlice {
	//panicf panics with a pre-formatted error string
	panicf := func(f string, s ...interface{}) {
		panic(fmt.Sprintf("structslice: invalid input on 'v' interface{}, 'v' must be a slice of structs. %s", fmt.Sprintf(f, s)))
	}

	s := new(structSlice)
	s.Value = reflect.ValueOf(v)

	// Test one: Panics if the v interface isn't a slice
	if t := reflect.TypeOf(s); t.Kind() != reflect.Ptr && (t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Slice) {
		panicf("v's Kind is %v, expected reflect.Slice\n", t.Kind())
	}

	// Test two: Panics if the elements of v are not structs
	if reflect.TypeOf(s).Elem().Kind() != reflect.Struct {
		panicf("v is a slice, but it contains elements of Kind %s, expected reflect.Struct\n", reflect.TypeOf(s).Kind())
	}

	return s
}

// Comparer is an interface for types that can compare themselves to each other.
type Comparer interface {
	Less(Comparer) bool
}

type Stringer interface {
	String() string
}

// Less satisfies the sort.Interface type in the go standard library
func (s structSlice) Less(i, j int) bool {
	it := s.Value.Index(i).Type().Field(s.SortFieldIndex)
	jt := s.Value.Index(j).Type().Field(s.SortFieldIndex)

	if it.Type.Kind() != jt.Type.Kind() {
		panic(fmt.Sprintf("structSlice.Less(): Type mismatch %s != %s", it.Type.Name(), jt.Type.Name()))
	}

	iv := s.Value.Index(i).Field(s.SortFieldIndex).Interface()
	jv := s.Value.Index(j).Field(s.SortFieldIndex).Interface()

	switch t := iv.(type) {
	case string:
		return t < jv.(string)
	case bool:
		return t && !jv.(bool)
	case int:
		return t < jv.(int)
	case int64:
		return t < jv.(int64)
	case float64:
		return t < jv.(float64)
	case float32:
		return t < jv.(float32)
	case Stringer:
		return t.String() < jv.(Stringer).String()
	case Comparer:
		return t.Less(jv.(Comparer))
	}

	return false
}

// Len satisfies the sort.Interface type in the go standard library
func (s structSlice) Len() int {
	return s.Value.Len()
}

// Swap satisfies the sort.Interface type in the go standard library
func (s structSlice) Swap(i, j int) {
	v := s.Value
	tmp := v.Index(i).Interface()
	v.Index(i).Set(v.Index(j))
	v.Index(j).Set(reflect.ValueOf(tmp))
}

// Index returns the value given by the index of the struct slice
func (s structSlice) Index(i int) reflect.Value {
	return s.Value.Index(i)
}
