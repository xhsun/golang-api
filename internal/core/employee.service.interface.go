package core

import "context"

// IEmployeeService - interface for employee business service
type IEmployeeService interface {
	// RetrieveAllEmployees - return a list of all existing employees
	RetrieveAllEmployees(c context.Context) ([]Employee, error)
	// CreateEmployee - Create and save a new employee
	//
	// On success, this method will return the ID of the employee created
	CreateEmployee(ctx context.Context, gender string) (int, error)
}
