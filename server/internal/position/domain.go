package position

import (
	"context"
	// "net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/alsey89/gogetter/pkg/jwt_manager"
	"github.com/alsey89/gogetter/pkg/pg_connector"
	"github.com/alsey89/gogetter/pkg/server"
	"github.com/alsey89/people-matter/internal/common/middleware"
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

	// *EMPLOYEE
	// can view their own user position information
	// employeeGroup := e.Group(
	// 	"api/v1/position",
	// 	d.params.JWT.GetJWTMiddleware("jwt_auth"),
	// )
	// employeeGroup.GET("", d.GetSelfPositionHandler)

	// *MANAGER
	// can assign, reassign or dismiss users in their branch from their positions
	// managerGroup := e.Group(
	// 	"api/v1/manager/user",
	// 	d.params.JWT.GetJWTMiddleware("jwt_auth"),
	// 	middleware.MustBeManager,
	// )
	// managerGroup.POST("/:userId/position", d.CreateUserPositionHandler)
	// managerGroup.PUT("/:userId/position", d.UpdateUserPositionHandler)
	// managerGroup.DELETE("/:userId/position", d.DeleteUserPositionHandler)

	// *ADMIN
	// can assign, reassign or dismiss users from their positions
	adminGroup := e.Group(
		"api/v1/admin",
		d.params.JWT.GetJWTMiddleware("jwt_auth"),
		middleware.MustBeAdmin,
	)
	// for selector options
	adminGroup.GET("/position", d.GetAllPositionsHandler)
	// for assignment
	adminGroup.POST("/user/:userId/position", d.AssignPositionHandler)
	adminGroup.DELETE("/user/:userId/position/:positionId", d.UnassignPositionHandler)
}

// ----------------------------------
