package main

import (
	"fmt"
)

import (
	ss "github.com/as/structslice"
)

// Name holds a person's name
type Name struct {
	last, middle, first string
}

// String returns a string representation of the person's name
func (n Name) String() string {
	return fmt.Sprintf("%-10s %-2s %-12s", n.last, n.middle, n.first)
}

// Implements structslice.Comparer so that the Name can be sorted as a field
func (n Name) Less(c ss.Comparer) bool {
	n2 := c.(Name)
	return (n.last + n.middle + n.first) < (n2.last + n2.middle + n2.first)
}

// Employee is a person who works for a company
type Employee struct {
	ID, Salary int
	Name       Name
}

// Print prints employee information to stdout
func (e Employee) Print() {
	fmt.Printf("%-2d %s %9d\n", e.ID, e.Name, e.Salary)
}

// Employees is a slice of Employee records
type Employees []Employee

// Print prints the Employee info to stdout
func (e Employees) Print() {
	fmt.Println("===================== Employee Roster =====================")
	fmt.Printf("%-2s %-26s %9s \n", "ID", "Name", "Salary")
	for _, v := range e {
		v.Print()
	}
	fmt.Println("===========================================================")
	fmt.Println()

}

func main() {
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
}
