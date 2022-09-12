package web

import (
	"golang-api/internal/core"
	"net/http"

	"github.com/gin-gonic/gin"
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

// GetEmployees - GET method to get all employees
func (h *EmployeeHandler) GetEmployees(c *gin.Context) {
	result, err := h.employeeService.RetrieveAllEmployees(c.Request.Context())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, result)
}
