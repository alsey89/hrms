package company

import (
	"context"
	// "net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/alsey89/gogetter/pkg/pg_connector"

	"github.com/alsey89/gogetter/pkg/mailer"
	"github.com/alsey89/gogetter/pkg/server"

	auth "github.com/alsey89/people-matter/internal/auth"
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

	// Router
	server := d.params.Server.GetServer()
	companyGroup := server.Group("api/v1/company")

	adminGroup := companyGroup.Group("/admin")

	// Routes
	companyGroup.GET("", d.GetCompanyHandler)
	companyGroup.POST("", d.CreateCompanyHandler)
	adminGroup.PUT("", d.UpdateCompanyHandler)
	adminGroup.DELETE("", d.DeleteCompanyHandler)

	adminGroup.POST("/department", d.CreateDepartmentHandler)
	adminGroup.PUT("/department/:department_id", d.UpdateDepartmentHandler)
	adminGroup.DELETE("/department/:department_id", d.DeleteDepartmentHandler)

	adminGroup.POST("/location", d.CreateLocationHandler)
	adminGroup.PUT("/location/:location_id", d.UpdateLocationHandler)
	adminGroup.DELETE("/location/:location_id", d.DeleteLocationHandler)

	adminGroup.POST("/position", d.CreatePositionHandler)
	adminGroup.PUT("/position/:position_id", d.UpdatePositionHandler)
	adminGroup.DELETE("/position/:position_id", d.DeletePositionHandler)

	return nil
}

func (d *Domain) onStop(ctx context.Context) error {
	d.logger.Info("Stopped APIs")

	return nil
}
