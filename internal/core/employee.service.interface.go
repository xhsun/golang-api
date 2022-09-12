package core

import "context"

// IEmployeeService - interface for employee business service
type IEmployeeService interface {
	// RetrieveAllEmployees - return a list of all existing employees
	RetrieveAllEmployees(c context.Context) ([]Employee, error)
}
