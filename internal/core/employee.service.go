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

type EmployeeService struct {
	employeeRepository IEmployeeRepository
}

// NewEmployeeService - creates a new EmployeeService
func NewEmployeeService(employeeRepository IEmployeeRepository) *EmployeeService {
	return &EmployeeService{
		employeeRepository: employeeRepository,
	}
}

func (s *EmployeeService) RetrieveAllEmployees(c context.Context) ([]Employee, error) {
	return s.employeeRepository.GetEmployees(c)
}

func (s *EmployeeService) AddEmployee(c context.Context, gender string) (int, error) {
	return s.employeeRepository.AddEmployee(c, gender)
}
