package web

import (
	"golang-api/internal/core"
)

type EmployeeHandler struct {
	employeeService core.IEmployeeService
}

// NewEmployeeHandler - creates a new EmployeeHandler
func NewEmployeeHandler(employeeService core.IEmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		employeeService: employeeService,
	}
}

// TODO: Get All Employees RPC

// TODO: Add Employee RPC
