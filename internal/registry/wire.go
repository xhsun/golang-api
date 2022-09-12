//go:build wireinject
// +build wireinject

package registry

import (
	"golang-api/ent"
	"golang-api/internal/config"
	"golang-api/internal/core"
	"golang-api/internal/infra"
	"golang-api/internal/web"

	"github.com/google/wire"
)

// InitializeServer - initialize the server
func InitializeServer(config config.Config, client *ent.Client) (*web.Server, error) {
	wire.Build(
		infra.NewEmployeeMapper,
		infra.NewEmployeeRepository, wire.Bind(new(core.IEmployeeRepository), new(*infra.EmployeeRepository)),
		core.NewEmployeeService, wire.Bind(new(core.IEmployeeService), new(*core.EmployeeService)),
		web.NewEmployeeHandler,
		web.NewServer,
	)
	return &web.Server{}, nil
}
