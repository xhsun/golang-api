package core

import "context"

// IEmployeeRepository - interface for employee repository
//
//go:generate mockery --name IEmployeeRepository
type IEmployeeRepository interface {
	// GetEmployees - get all employees from datastore
	GetEmployees(ctx context.Context) ([]Employee, error)
	// AddEmployee - add an employee to datastore and return datastore ID of the employee
	AddEmployee(ctx context.Context, gender string) (int, error)
}
