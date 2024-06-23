package main

import (
	"go.uber.org/fx"

	"github.com/alsey89/gogetter/pkg/config_manager"
	"github.com/alsey89/gogetter/pkg/jwt_manager"
	"github.com/alsey89/gogetter/pkg/logger"
	"github.com/alsey89/gogetter/pkg/mailer"
	"github.com/alsey89/gogetter/pkg/pg_connector"
	"github.com/alsey89/gogetter/pkg/server"

	"github.com/alsey89/people-matter/internal/auth"
	"github.com/alsey89/people-matter/internal/company"
	"github.com/alsey89/people-matter/internal/position"
	"github.com/alsey89/people-matter/internal/user"
	"github.com/alsey89/people-matter/schema"
)

var config *config_manager.Module

func init() {
	config_manager.SetSystemLogLevel("debug")
	config = config_manager.SetUpConfig("SERVER", "yaml")
	//! CONFIG PRECEDENCE: ENV > CONFIG FILE > FALLBACK
	config.SetFallbackConfigs(map[string]interface{}{
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
		"database.auto_migrate": true,

		"mailer.host":         "smtp.mailersend.net",
		"mailer.port":         587,
		"mailer.username":     "MS_Rf5egx@peoplematter.app",
		"mailer.app_password": "ZishGZriU3z2KHPf",
		"mailer.tls":          true,

		"jwt_auth.signing_key":    "authsecret",
		"jwt_auth.token_lookup":   "cookie:jwt",
		"jwt_auth.signing_method": "HS256",
		"jwt_auth.exp_in_hours":   72,

		"jwt_email.signing_key":    "emailsecret",
		"jwt_email.token_lookup":   "query:token",
		"jwt_email.signing_method": "HS256",
		"jwt_email.exp_in_hours":   1,
	})
}
func main() {
	app := fx.New(
		fx.Supply(config),
		logger.InitiateModule(),
		server.InitiateModule("server"),
		pg_connector.InitiateModuleAndSchema(
			"database",
			// ...schema,
			schema.Company{},
			schema.Department{},
			schema.Location{},
			schema.User{},
			schema.Role{},
			schema.UserRole{},
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
		jwt_manager.InitiateModule("jwt", "jwt_auth", "jwt_email"),
		mailer.InitiateModule("mailer"),

		//-- Internal Domains Start --
		auth.InitiateDomain("auth"),
		company.InitiateDomain("company"),
		user.InitiateDomain("user"),
		position.InitiateDomain("position"),
		//-- Internal Domains End --
		// fx.NopLogger,
	)
	app.Run()
}
