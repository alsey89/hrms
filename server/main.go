package main

import (
	"go.uber.org/fx"

	config "github.com/alsey89/gogetter/config/viper"
	postgres "github.com/alsey89/gogetter/database/postgres"
	jwt "github.com/alsey89/gogetter/jwt/echo"
	logger "github.com/alsey89/gogetter/logging/zap"
	mailer "github.com/alsey89/gogetter/mail/gomail"
	server "github.com/alsey89/gogetter/server/echo"
)

var configuration *config.Config

func init() {
	config.SetSystemLogLevel("debug")
	configuration = config.SetUpConfig("SERVER", "yaml")
	configuration.SetFallbackConfigs(map[string]interface{}{
		"server.host": "0.0.0.0",
		"server.port": 3001,

		"database.host":         "postgres",
		"database.port":         5432,
		"database.dbname":       "postgres",
		"database.user":         "postgres",
		"database.password":     "password",
		"database.sslmode":      "prefer",
		"databse.loglevel":      "error",
		"database.auto_migrate": true,

		"mailer.host":         "smtp.gmail.com",
		"mailer.port":         587,
		"mailer.username":     "phyokyawsoe89@gmail.com",
		"mailer.app_password": "lyzo mila fxha dupi",
		"mailer.tls":          true,

		"auth_jwt.signing_key":  "thisisasecret",
		"auth_jwt.token_lookup": "cookie:jwt",
	})
}
func main() {
	app := fx.New(
		fx.Supply(configuration),
		logger.InitiateModule(),

		jwt.InitiateModule("auth_jwt"),
		mailer.InitiateModule("mailer"),
		postgres.InitiateModuleAndSchema(
			"database",
			// ...schema,
			// example: &User{},
			// example: &Post{},
			// example: &Comment{},
		),
		server.InitiateModule("server"),

		fx.NopLogger,
	)
	app.Run()
}
