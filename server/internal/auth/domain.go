package auth

import (
	"context"
	"fmt"

	// "net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"

	postgres "github.com/alsey89/gogetter/database/postgres"
	jwt "github.com/alsey89/gogetter/jwt"
	server "github.com/alsey89/gogetter/server/echo"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

const (
	defaultSigningKey    = "othersecret"
	defaultSigningMethod = "HS256"
	defaultExpInHours    = 1
)

type Domain struct {
	config *Config
	params Params
	scope  string
	logger *zap.Logger
}

type Config struct {
	SigningKey    string
	TokenLookup   string
	SigningMethod string
	ExpInHours    int
}

type Params struct {
	fx.In

	Lifecycle fx.Lifecycle
	Logger    *zap.Logger
	Server    *server.Module
	Database  *postgres.Module
	JWT       *jwt.Module
}

func InitiateDomain(scope string) fx.Option {

	var d *Domain

	return fx.Options(
		fx.Provide(func(p Params) *Domain {

			c := loadConfig(scope)

			d := &Domain{
				config: c,
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

func loadConfig(scope string) *Config {
	getConfigPath := func(key string) string {
		return fmt.Sprintf("%s.%s", scope, key)
	}

	//set defaults
	viper.SetDefault(getConfigPath("signing_key"), defaultSigningKey)
	viper.SetDefault(getConfigPath("signing_method"), defaultSigningMethod)
	viper.SetDefault(getConfigPath("exp_in_hours"), defaultExpInHours)

	return &Config{
		SigningKey:    viper.GetString(getConfigPath("signing_key")),
		SigningMethod: viper.GetString(getConfigPath("signing_method")),
		ExpInHours:    viper.GetInt(getConfigPath("exp_in_hours")),
	}
}

func (d *Domain) onStart(ctx context.Context) error {

	d.logger.Info("Starting APIs")

	// d.AddDefaultData(ctx)

	// Router
	server := d.params.Server.GetServer()
	authGroup := server.Group("api/v1/auth")

	// authGroup.POST("/signup", d.SignupHandler)
	authGroup.POST("/signin", d.SigninHandler)
	authGroup.POST("/signout", d.SignoutHandler)

	authGroup.GET("/check", d.CheckAuth, d.GetAuthJWTMiddleware())
	// authGroup.GET("/check", d.CheckAuth, d.params.JWT.Middleware())
	authGroup.GET("/csrf", d.GetCSRFToken)

	authGroup.GET("/confirmation", d.ConfirmationHandler)

	d.PrintDebugLogs()
	return nil
}

func (d *Domain) onStop(ctx context.Context) error {
	d.logger.Info("Stopped APIs")

	return nil
}

func (m *Domain) PrintDebugLogs() {
	m.logger.Debug("----- Auth Domain Configuration -----")
	m.logger.Debug("SigningKey", zap.Any("SigningKey", m.config.SigningKey))
	m.logger.Debug("SigningMethod", zap.String("SigningMethod", m.config.SigningMethod))
	m.logger.Debug("ExpInHours", zap.Int("ExpInHours", m.config.ExpInHours))
}

// ----------------------------------

func (d *Domain) GetAuthJWTMiddleware() echo.MiddlewareFunc {
	authConfig, err := d.params.JWT.GetConfig("jwt_auth")
	if err != nil {
		d.logger.Error("error getting jwt auth config", zap.Error(err))
	}

	middleware := d.params.Server.GetEchoJWTMiddleware(authConfig.SigningKey, authConfig.SigningMethod, authConfig.TokenLookup)
	if middleware == nil {
		d.logger.Error("error getting jwt auth middleware")
	}

	return *middleware
}
