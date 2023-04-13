package core

import "context"

// IEmployeeService - interface for employee business service
//
//go:generate mockery --name IEmployeeService
type IEmployeeService interface {
	// RetrieveAllEmployees - return a list of all existing employees
	RetrieveAllEmployees(c context.Context) ([]Employee, error)
	// AddEmployee - store employee and return the employee ID
	AddEmployee(c context.Context, gender string) (int, error)
}
