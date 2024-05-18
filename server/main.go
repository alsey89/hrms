package main

import (
	"go.uber.org/fx"

	config "github.com/alsey89/gogetter/config/viper"
	postgres "github.com/alsey89/gogetter/database/postgres"
	jwt "github.com/alsey89/gogetter/jwt/echo"
	logger "github.com/alsey89/gogetter/logging/zap"
	mailer "github.com/alsey89/gogetter/mail/gomail"
	server "github.com/alsey89/gogetter/server/echo"

	"github.com/alsey89/people-matter/internal/auth"
	"github.com/alsey89/people-matter/internal/company"
	"github.com/alsey89/people-matter/schema"
)

var configuration *config.Module

func init() {
	config.SetSystemLogLevel("debug")
	configuration = config.SetUpConfig("SERVER", "yaml")
	//! CONFIG PRECEDENCE: ENV > CONFIG FILE > FALLBACK
	configuration.SetFallbackConfigs(map[string]interface{}{
		"server.host":      "0.0.0.0",
		"server.port":      3001,
		"server.log_level": "DEV",

		// "server.allow_headers":   "*",
		// "server.allow_methods":   "*",
		"server.allow_origins":   "http://localhost:3000, http://localhost:3001",
		"server.csrf_protection": true,
		"server.csrf_secure":     false,
		"server.csrf_domain":     "localhost",

		"database.host":         "postgres",
		"database.port":         5432,
		"database.dbname":       "postgres",
		"database.user":         "postgres",
		"database.password":     "password",
		"database.sslmode":      "prefer",
		"databse.loglevel":      "error",
		"database.auto_migrate": false,

		"mailer.host":         "smtp.mailersend.net",
		"mailer.port":         587,
		"mailer.username":     "MS_ZsiAoC@peoplematter.app",
		"mailer.app_password": "jmU7b5NxKr3Da75n",
		"mailer.tls":          true,

		"echo_jwt.signing_key":    "authsecret",
		"echo_jwt.token_lookup":   "cookie:jwt",
		"echo_jwt.signing_method": "HS256",
		"echo_jwt.exp_in_hours":   72,

		"auth.signing_key":    "confirmationsecret",
		"auth.signing_method": "HS256",
		"auth.exp_in_hours":   1,
	})
}
func main() {
	app := fx.New(
		fx.Supply(configuration),
		logger.InitiateModule(),
		server.InitiateModule("server"),
		postgres.InitiateModuleAndSchema(
			"database",
			// ...schema,
			schema.Company{},
			schema.Department{},
			schema.Location{},
			schema.User{},
			schema.ContactInfo{},
			schema.EmergencyContact{},
			schema.Position{},
			schema.UserPosition{},
			schema.Leave{},
			schema.Attendance{},
			schema.Salary{},
			schema.Payment{},
			schema.Adjustments{},
			schema.Document{},
		),
		jwt.InitiateModule("echo_jwt"),
		mailer.InitiateModule("mailer"),

		//-- Internal Domains Start --
		auth.InitiateDomain("auth"),
		company.InitiateDomain("company"),
		//-- Internal Domains End --
		fx.NopLogger,
	)
	app.Run()
}
