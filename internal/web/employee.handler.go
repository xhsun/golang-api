package web

import (
	"golang-api/internal/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Employee - DTO model for Employee
type Employee struct {
	Gender string `json:"gender" binding:"required,alphanum"`
}

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

// PostEmployee - POST method to add an employee
func (h *EmployeeHandler) PostEmployee(c *gin.Context) {
	var input Employee
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	result, err := h.employeeService.AddEmployee(c.Request.Context(), input.Gender)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, result)
}
