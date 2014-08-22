## StructSlice
Package structslice provides an interface for sorting struct slices by their field names or field numbers

## Quick tour

```
	import ss "github.com/as/structslice"

	// Employee is a person who works for a company
	type Employee struct {
		ID, Salary int
		Name       Name
	}

	// A database of empyloyees
	internalDB := Employees{
		Employee{1, 95000, Name{"Jake", "M", "Anderson"}},
		Employee{5, 45000, Name{"Hunter", "L", "Alice"}},
		Employee{6, 345000, Name{"Steinberg", "F", "Charles"}},
		Employee{2, 108000, Name{"Williams", "L", "Bill"}},
		Employee{4, 190000, Name{"Morgan", "A", "Janice"}},
		Employee{3, 108000, Name{"Williams", "L", "Will"}},
		Employee{5, 145000, Name{"Steinberg", "L", "Alice"}},
	}

	fmt.Println("A messy database:")
	internalDB.Print()

	fmt.Println("Sort by ID")
	ss.SortByName(internalDB, "ID")
	internalDB.Print()

	fmt.Println("Sort by Name and Salary (People w/ same salary should be ordered by name)")
	ss.SortStableByName(internalDB, "Name", "Salary")
	internalDB.Print()

```

## TODO
~~* Sort functions should return an error value ~~
* Unit tests
* Recursively sort embedded structs (maybe)

## Possible GOTCHAS
* SortByName and SortByIndex are not variadic functions, running a non-stable sort for multiple keys won't work.
