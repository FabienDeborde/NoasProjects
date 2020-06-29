package router

import (
	handlers "github.com/FabienDeborde/noas_projects/app/handlers"
	"github.com/gofiber/fiber"
)

func projectRoutes(group *fiber.Group) {
	projects := group.Group("/projects") // /api/v1/projects
	projects.Get("/", handlers.GetProjects)
	// projects.Post("/", project.NewProject)
	// projects.Get("/:id", project.GetProject)
	// projects.Delete("/:id", project.DeleteProject)
	// projects.Post("/:id/like", project.AddLikeProject)
	// projects.Post("/:id/unlike", project.RemoveLikeProject)
}
