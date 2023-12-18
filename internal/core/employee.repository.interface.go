package core

import "context"

// IEmployeeRepository - interface for employee repository
type IEmployeeRepository interface {
	// GetEmployees - get all employees from datastore
	GetEmployees(ctx context.Context) ([]Employee, error)
	// Add Employee - Add a new employee to datastore
	AddEmployee(ctx context.Context, gender string) (int, error)
}
