package infra

import (
	"context"
	"fmt"
	"golang-api/ent"
	"golang-api/internal/core"
)

type EmployeeRepository struct {
	client *ent.Client
	mapper *EmployeeMapper
}

// NewEmployeeRepository - creates a new EmployeeRepository
func NewEmployeeRepository(client *ent.Client, mapper *EmployeeMapper) *EmployeeRepository {
	return &EmployeeRepository{
		client: client,
		mapper: mapper,
	}
}

func (r *EmployeeRepository) GetEmployees(ctx context.Context) ([]core.Employee, error) {
	e, err := r.client.Employees.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying employees: %w", err)
	}
	return r.mapper.ToDomains(e), nil
}
