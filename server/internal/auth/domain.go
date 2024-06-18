package auth

import (
	"context"
	"fmt"

	// "net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/alsey89/gogetter/pkg/jwt_manager"
	"github.com/alsey89/gogetter/pkg/pg_connector"

	"github.com/alsey89/gogetter/pkg/mailer"
	"github.com/alsey89/gogetter/pkg/server"

	jwtV5 "github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// ----------------------------------

const (
	defaultJWTAuthScope  = "jwt_auth"
	defaultJWTEmailScope = "jwt_email"
	defaultJWTResetScope = "jwt_reset"
)

type Domain struct {
	config *Config
	params Params
	scope  string
	logger *zap.Logger
}

type Config struct {
	JWTAuthScope  string
	JWTEmailScope string
	JWTResetScope string
}

type Params struct {
	fx.In

	Lifecycle fx.Lifecycle
	Logger    *zap.Logger
	Server    *server.Module
	Database  *pg_connector.Module
	JWT       *jwt_manager.Module
	Mailer    *mailer.Module
}

// ----------------------------------

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

func (d *Domain) onStart(ctx context.Context) error {

	d.logger.Info("Starting APIs")

	err := d.registerRoutes()
	if err != nil {
		d.logger.Error("error loading routes", zap.Error(err))
	}

	d.printDebugLogs()
	return nil
}

func (d *Domain) onStop(ctx context.Context) error {
	d.logger.Info("Stopped APIs")

	return nil
}

func loadConfig(scope string) *Config {
	getConfigPath := func(key string) string {
		return fmt.Sprintf("%s.%s", scope, key)
	}

	//set defaults
	viper.SetDefault(getConfigPath("jwt_auth_scope"), defaultJWTAuthScope)
	viper.SetDefault(getConfigPath("jwt_email_scope"), defaultJWTEmailScope)
	viper.SetDefault(getConfigPath("jwt_reset_scope"), defaultJWTResetScope)

	return &Config{
		JWTAuthScope:  viper.GetString(getConfigPath("jwt_auth_scope")),
		JWTEmailScope: viper.GetString(getConfigPath("jwt_email_scope")),
		JWTResetScope: viper.GetString(getConfigPath("jwt_reset_scope")),
	}
}

func (d *Domain) registerRoutes() error {
	d.logger.Info("Loading Routes")

	e := d.params.Server.GetServer()
	authGroup := e.Group("api/v1/auth")

	authGroup.POST("/signin", d.SigninHandler)
	// authGroup.POST("/signup", d.SignupHandler)
	authGroup.POST("/signout", d.SignoutHandler)

	authGroup.GET("/confirmation",
		d.ConfirmationHandler,
		d.params.JWT.GetJWTMiddleware("jwt_email"),
	)

	authGroup.GET("/check", d.CheckAuth, d.mustBeLoggedIn())
	authGroup.GET("/csrf", d.GetCSRFToken)

	return nil
}

func (m *Domain) printDebugLogs() {
	m.logger.Debug("----- Auth Domain Configuration -----")
	m.logger.Debug("JWT_Auth_Scope", zap.Any("JWT_Auth_Scope", m.config.JWTAuthScope))
	m.logger.Debug("JWT_Email_Scope", zap.Any("JWT_Email_Scope", m.config.JWTEmailScope))
	m.logger.Debug("JWT_Reset_Scope", zap.Any("JWT_Reset_Scope", m.config.JWTResetScope))
}

// ----------------------------------

func (d *Domain) mustBeLoggedIn() echo.MiddlewareFunc {
	authConfig, err := d.params.JWT.GetConfig("jwt_auth")
	if err != nil {
		d.logger.Error("error getting jwt auth config", zap.Error(err))
		return nil
	}

	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(authConfig.SigningKey),
		SigningMethod: authConfig.SigningMethod,
		TokenLookup:   authConfig.TokenLookup,
	})

}

// ----------------------------------

func (d *Domain) SendConfirmationEmail(emailAddress string, userID uint, CompanyID uint) error {
	//generate jwt token
	additionalClaims := jwtV5.MapClaims{
		"id":        userID,
		"companyId": CompanyID,
	}

	token, err := d.params.JWT.GenerateToken("jwt_email", additionalClaims)
	if err != nil {
		d.logger.Error("[CreateNewCompanyAndAdminUser]", zap.Error(err))
	}

	//send confirmation email
	err = d.params.Mailer.SendTransactionalMail(
		"hello@peoplematter.app",
		emailAddress,
		"Welcome to People Matter",
		"<p>Welcome to People Matter</p><p>Your account has been created. Please click the link below to confirm your email address.</p><a href=\"http://localhost:3000/onboarding/confirmation?token="+*token+"\">Confirm Email</a>",
	)
	if err != nil {
		return fmt.Errorf("[CreateNewCompanyAndAdminUser] Error sending confirmation email %w", err)
	}

	return nil
}
