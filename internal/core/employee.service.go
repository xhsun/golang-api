package core

import "context"

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

func (s *EmployeeService) CreateEmployee(ctx context.Context, gender string) (int, error) {
	return s.employeeRepository.AddEmployee(ctx, gender)
}
