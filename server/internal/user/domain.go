package user

import (
	"context"
	// "net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/alsey89/gogetter/pkg/jwt_manager"
	"github.com/alsey89/gogetter/pkg/pg_connector"
	"github.com/alsey89/gogetter/pkg/server"
	"github.com/alsey89/people-matter/internal/middleware"
)

// ----------------------------------

type Domain struct {
	params Params
	scope  string
	logger *zap.Logger
}

type Params struct {
	fx.In

	Lifecycle fx.Lifecycle
	Logger    *zap.Logger
	Server    *server.Module
	Database  *pg_connector.Module
	JWT       *jwt_manager.Module
}

// ----------------------------------

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

	// // *System
	// // for creating and deleting companies
	// systemGroup := e.Group(
	// 	"api/v1/system/user",
	// 	d.params.JWT.GetJWTMiddleware("jwt_auth"),
	// )
	// systemGroup.POST("", d.CreateRootUserHandler)
	// systemGroup.DELETE("", d.DeleteRootUserHandler, middleware.MustBeAdmin)

	// // *EMPLOYEE
	// // can view their own user information
	// employeeGroup := e.Group(
	// 	"api/v1/user",
	// 	d.params.JWT.GetJWTMiddleware("jwt_auth"),
	// )
	// employeeGroup.GET("", d.GetSelfDataHandler)
	// employeeGroup.PUT("", d.UpdateSelfDataHandler)

	// *MANAGER
	// can view all users from their branch
	// can update, delete user from their branch
	// managerGroup := e.Group(
	// 	"api/v1/manager/user",
	// 	d.params.JWT.GetJWTMiddleware("jwt_auth"),
	// 	middleware.MustBeManager,
	// )
	// managerGroup.GET("", d.GetAllLocationUsersHandler)
	// managerGroup.POST("", d.CreateLocationUserHandler)
	// managerGroup.PUT("/user/:userId", d.UpdateLocationUserHandler)
	// managerGroup.DELETE("/user/:userId", d.DeleteLocationUserHandler)
	// *ADMIN
	// can view all users information
	// can update, delete user from any branch
	adminGroup := e.Group(
		"api/v1/admin/user",
		d.params.JWT.GetJWTMiddleware("jwt_auth"),
		middleware.MustBeAdmin,
	)
	adminGroup.GET("", d.GetAllUsersHandler)
	adminGroup.POST("", d.CreateUserHandler)
	adminGroup.PUT("/:userId", d.UpdateUserHandler)
	adminGroup.DELETE("/:userId", d.DeleteUserHandler)
}

// ----------------------------------
