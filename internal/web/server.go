package web

import (
	"golang-api/internal/config"

	"github.com/gin-gonic/gin"
)

// Server - the REST API server
type Server struct {
	config          config.Config
	employeeHandler *EmployeeHandler
}

// NewServer - creates a new Server
func NewServer(config config.Config, employeeHandler *EmployeeHandler) *Server {
	return &Server{
		config:          config,
		employeeHandler: employeeHandler,
	}
}

// Start - start the REST API server
//
// Note: By default it serves on :8080 unless a PORT environment variable was defined
func (s *Server) Start() error {
	router := gin.Default()
	router.GET("/employees", s.employeeHandler.GetEmployees)
	return router.Run()
}
