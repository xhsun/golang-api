package main

import (
	"flag"
	"fmt"
)

func main() {
	add := flag.Bool("add", false, "Set this to true to add an employee")
	gender := flag.String("g", "undefined", "For Add mode only: Set gender of the employee")
	flag.Parse()

	if add != nil && *add {
		// TODO: Add a new employee with the given gender
		fmt.Println(*gender)
	} else {
		// TODO: Get all employees
	}

	panic("Implement the client")
}
