package main

import (
	"fmt"
)

type employee struct {
	firstName, lastName string
	age, sarlay         int
}

func main() {
	emp1 := employee{
		firstName: "Sam",
		age:       25,
		sarlay:    500,
		lastName:  "Andera",
	}
	emp2 := employee{"Thromas", "Json", 30, 900}

	emp3 := struct {
		firstName, lastName string
		age, slary          int
	}{
		firstName: "Mark",
		lastName:  "Tracy",
		age:       22,
		slary:     10000,
	}

	var emp4 employee

	emp5 := employee{
		age:    555,
		sarlay: 77777,
	}

	emp6 := employee{"Rose", "Keleary", 18, 3000}

	var emp7 employee
	emp7.firstName = "Jack"
	emp7.lastName = "Ricky"

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)
	fmt.Println(emp3)
	fmt.Println(emp4)
	fmt.Println(emp5)
	fmt.Println("The firstname ", emp6.firstName)
	fmt.Println("The lastname ", emp6.lastName)
	fmt.Println(emp7)
}
