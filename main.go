package main

import (
	"fmt"

	"github.com/FabienDeborde/noas_projects/database"
	"github.com/FabienDeborde/noas_projects/project"
	"github.com/FabienDeborde/noas_projects/utils/logger"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("api/v1/projects", project.GetProjects)
	app.Get("api/v1/projects/:id", project.GetProject)
	app.Post("api/v1/projects", project.NewProject)
	app.Delete("api/v1/projects/:id", project.DeleteProject)
	app.Post("api/v1/projects/:id/like", project.AddLikeProject)
	app.Post("api/v1/projects/:id/unlike", project.RemoveLikeProject)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "projects.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened!")

	database.DBConn.AutoMigrate(&project.Project{})
	fmt.Println("Database migrated!")

}

func main() {
	_, slogger := logger.Init()

	err := godotenv.Load()
	if err != nil {
		slogger.Fatal("Error loading .env file")
	}

	// Pass Settings creating a new instance
	app := fiber.New(&fiber.Settings{
		Prefork:      false,
		ServerHeader: "Fiber",
		BodyLimit:    4 * 1024 * 1024,
	})

	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404) // => 404 "Not Found"
	})

	app.Listen(3000)
}
