package bootstrap

import (
	"gorm.io/gorm"
)

type Application struct {
	Env    *Env
	DB     *gorm.DB // Changed to *gorm.DB
	Logger *Logger  // Add Logger field
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Logger = NewLogger() // Initialize Logger
	app.DB = NewMySQLDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMySQLConnection(app.DB)
}
