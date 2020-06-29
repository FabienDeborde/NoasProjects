package main

import (
	"os"

	"github.com/FabienDeborde/noas_projects/app"

	"github.com/FabienDeborde/noas_projects/app/utils/logger"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
)

func main() {
	_, slogger := logger.Init()

	// TODO: only use .env in development
	err := godotenv.Load()
	if err != nil {
		slogger.Error("Error loading .env file")
	}

	app, err := app.Init()
	if err != nil {
		slogger.Error("Error")
	}

	// Get the PORT from heroku env
	port := os.Getenv("PORT")

	// Verify if heroku provided the port or not
	if port == "" {
		slogger.Warn("Couldn't get the PORT from environment. Switching to default PORT 3000.")
		port = "3000"
	}

	// Start server on http://${heroku-url}:${port}
	app.Listen(port)
}
