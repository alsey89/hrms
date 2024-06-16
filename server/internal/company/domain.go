package company

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/alsey89/gogetter/pkg/jwt_manager"
	"github.com/alsey89/gogetter/pkg/mailer"
	"github.com/alsey89/gogetter/pkg/pg_connector"
	"github.com/alsey89/gogetter/pkg/server"
	"github.com/alsey89/people-matter/internal/auth"
	"github.com/alsey89/people-matter/internal/middleware"
)

type Domain struct {
	params Params
	scope  string
	logger *zap.Logger
}

type Params struct {
	fx.In
	Lifecycle fx.Lifecycle

	//Modules
	Logger   *zap.Logger
	Server   *server.Module
	Database *pg_connector.Module
	Mailer   *mailer.Module
	JWT      *jwt_manager.Module

	//Internal Domains
	Auth *auth.Domain
}

func InitiateDomain(scope string) fx.Option {
	var d *Domain

	return fx.Options(
		fx.Provide(func(p Params) *Domain {
			d := &Domain{
				params: p,
				scope:  scope,
				logger: p.Logger.Named("[" + scope + "]"),
			}
			return d
		}),
		fx.Populate(&d),
		fx.Invoke(func(p Params) {
			p.Lifecycle.Append(
				fx.Hook{
					OnStart: d.onStart,
					OnStop:  d.onStop,
				},
			)
		}),
	)
}

func (d *Domain) onStart(ctx context.Context) error {
	d.logger.Info("Starting APIs")
	d.registerRoutes()
	return nil
}

func (d *Domain) onStop(ctx context.Context) error {
	d.logger.Info("Stopped APIs")
	return nil
}

func (d *Domain) registerRoutes() {
	// Router
	e := d.params.Server.GetServer()

	// *System
	// for creating and deleting companies
	companyGroup := e.Group(
		"api/v1/system/company",
		d.params.JWT.GetJWTMiddleware("jwt_auth"),
	)
	companyGroup.POST("", d.CreateCompanyHandler)
	companyGroup.DELETE("", d.DeleteCompanyHandler, middleware.MustBeAdmin)

	// *EMPLOYEE
	// can view company information
	employeeGroup := e.Group(
		"api/v1/company",
		d.params.JWT.GetJWTMiddleware("jwt_auth"),
	)
	employeeGroup.GET("", d.GetCompanyHandler)

	// *MANAGER
	// can view company information
	// can update department information
	managerGroup := e.Group(
		"api/v1/manager/company",
		d.params.JWT.GetJWTMiddleware("jwt_auth"),
		middleware.MustBeManager,
	)
	managerGroup.PUT("/location/:locationId", d.ManagerUpdateLocationHandler)

	// *ADMIN
	// can view company information
	// can update company information, create & delete at system level
	// can create, update, delete departments
	// can create, update, delete locations
	// can create, update, delete positions
	adminGroup := e.Group(
		"api/v1/admin/company",
		d.params.JWT.GetJWTMiddleware("jwt_auth"),
		middleware.MustBeAdmin,
	)
	adminGroup.PUT("", d.UpdateCompanyHandler)

	adminGroup.POST("/department", d.CreateDepartmentHandler)
	adminGroup.PUT("/department/:department_id", d.UpdateDepartmentHandler)
	adminGroup.DELETE("/department/:departmentId", d.DeleteDepartmentHandler)

	adminGroup.POST("/location", d.CreateLocationHandler)
	adminGroup.PUT("/location/:locationId", d.UpdateLocationHandler)
	adminGroup.DELETE("/location/:locationId", d.DeleteLocationHandler)

	adminGroup.POST("/position", d.CreatePositionHandler)
	adminGroup.PUT("/position/:positionId", d.UpdatePositionHandler)
	adminGroup.DELETE("/position/:positionId", d.DeletePositionHandler)
}
