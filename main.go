package main

import (
	"os"

	"github.com/FabienDeborde/noas_projects/database"
	"github.com/FabienDeborde/noas_projects/project"
	"github.com/FabienDeborde/noas_projects/utils/logger"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/helmet"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
)

func projectRoutes(group *fiber.Group) {
	projects := group.Group("/projects") // /api/v1/projects
	projects.Get("/", project.GetProjects)
	projects.Post("/", project.NewProject)
	projects.Get("/:id", project.GetProject)
	projects.Delete("/:id", project.DeleteProject)
	projects.Post("/:id/like", project.AddLikeProject)
	projects.Post("/:id/unlike", project.RemoveLikeProject)
}

func main() {
	_, slogger := logger.Init()

	err := godotenv.Load()
	if err != nil {
		slogger.Error("Error loading .env file")
	}

	// Pass Settings creating a new instance
	app := fiber.New(&fiber.Settings{
		Prefork:      false,
		ServerHeader: "NoasProjects",
		BodyLimit:    4 * 1024 * 1024,
	})
	app.Use(middleware.Recover())
	app.Use(middleware.Compress())
	app.Use(cors.New())
	app.Use(helmet.New())

	database.Init()
	database.DBConn.AutoMigrate(&project.Project{})
	defer database.DBConn.Close()

	// setupRoutes(app)
	v1 := app.Group("/api/v1") // /api/v1
	projectRoutes(v1)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404) // => 404 "Not Found"
	})

	// Log all registered routes
	for _, r := range app.Routes() {
		slogger.Infow("Routes",
			"Method", r.Method,
			"Path", r.Path,
		)
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
