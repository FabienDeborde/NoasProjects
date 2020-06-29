package app

import (
	"os"
	"strconv"

	"github.com/FabienDeborde/noas_projects/app/utils/logger"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/helmet"
)

// Init is app init
func Init() (*fiber.App, error) {
	_, slogger := logger.Init()
	// Get the PREFORK option from heroku env
	prefork := os.Getenv("PREFORK")
	preforkB, err := strconv.ParseBool(prefork)
	// Verify if heroku provided the port or not
	if err != nil {
		slogger.Warn("Couldn't get PREFORK from environment. Switching to default PREFORK = false.")
		preforkB = false
	}

	// Pass Settings creating a new instance
	app := fiber.New(&fiber.Settings{
		Prefork:      preforkB,
		ServerHeader: "NoasProjects",
		BodyLimit:    4 * 1024 * 1024,
	})
	app.Use(middleware.Recover()) // TODO: check if it is working?
	app.Use(middleware.Compress())
	app.Use(cors.New())
	app.Use(helmet.New())

	// TODO: move init router
	// database.Init()
	// database.DBConn.AutoMigrate(&project.Project{})
	// defer database.DBConn.Close()

	// router.Init(&app)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404) // => 404 "Not Found"
	})

	// Log all registered routes
	for _, r := range app.Routes() {
		if r.Method != "USE" {
			slogger.Infow("Routes",
				"Method", r.Method,
				"Path", r.Path,
			)
		}
	}

	return app, err
}
