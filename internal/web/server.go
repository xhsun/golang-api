package web

import (
	"fmt"
	"golang-api/internal/config"
	"net"

	"github.com/bufbuild/protovalidate-go"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	validateMiddleware "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/protovalidate"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// Server - the REST API server
type Server struct {
	log                *zap.Logger
	config             *config.Config
	employeeHandler    *EmployeeHandler
	healthCheckHandler *HealthCheckHandler
}

// NewServer - creates a new Server
func NewServer(log *zap.Logger, config config.Config, employeeHandler *EmployeeHandler, healthCheckHandler *HealthCheckHandler) *Server {
	return &Server{
		log:                log,
		config:             &config,
		employeeHandler:    employeeHandler,
		healthCheckHandler: healthCheckHandler,
	}
}

// Start - start the gRPC API server
//
// Note: By default it serves on :5051 unless a PORT environment variable was defined
func (s *Server) Start() error {
	validator, err := protovalidate.New()
	if err != nil {
		s.log.Fatal("failed to create a validator", zap.Error(err))
	}

	logOpts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}

	// Init gRPC server with interceptors (equivalent of HTTP middlewares)
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		logging.UnaryServerInterceptor(InterceptorLogger(s.log), logOpts...),
		validateMiddleware.UnaryServerInterceptor(validator),
		// Add any other interceptor you want.
	),
		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(InterceptorLogger(s.log), logOpts...),
			validateMiddleware.StreamServerInterceptor(validator),
			// Add any other interceptor you want.
		))

	// Register server handlers
	// TODO: Register employee handler
	panic("Register employee handler")

	grpc_health_v1.RegisterHealthServer(grpcServer, s.healthCheckHandler)

	// Listen to port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.Port))
	if err != nil {
		s.log.Fatal("failed to listen to port",
			zap.Uint16("port", s.config.Port),
			zap.Error(err),
		)
	}

	// Start server
	s.log.Info("Starting Server")
	return grpcServer.Serve(lis)
}
