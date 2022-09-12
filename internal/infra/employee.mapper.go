package infra

import (
	"golang-api/ent"
	"golang-api/internal/core"
)

// EmployeeMapper - model mapper for mapper DAO to domain and vice versa
type EmployeeMapper struct{}

// NewEmployeeMapper - creates a new EmployeeMapper
func NewEmployeeMapper() *EmployeeMapper {
	return &EmployeeMapper{}
}

// ToDomains - map DAOs to domain object
func (r *EmployeeMapper) ToDomains(daos []*ent.Employees) []core.Employee {
	result := []core.Employee{}
	for _, e := range daos {
		result = append(result, core.Employee{
			ID:     e.ID,
			Gender: e.Gender,
		})
	}

	return result
}
