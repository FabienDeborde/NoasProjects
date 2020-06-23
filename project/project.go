package project

import (
	"github.com/fabiendeborde/noas_projects/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	Title       string `json:"title" binding:"required" gorm:"type:varchar(100);"`
	Description string `json:"description" gorm:"type:varchar(1000);"`
	Link        string `json:"link" binding:"required" gorm:"type:varchar(100);"`
}

func GetProjects(c *fiber.Ctx) {
	db := database.DBConn
	var projects []Project
	db.Find((&projects))
	c.JSON(projects)
}
func GetProject(c *fiber.Ctx) {
	id := c.Params(("id"))
	db := database.DBConn

	var project Project
	db.Find(&project, id)
	if project.Title == "" {
		c.Status(500).Send("No project found with given ID")
		return
	}
	c.JSON(project)
}
func NewProject(c *fiber.Ctx) {
	db := database.DBConn

	project := new(Project)
	if err := c.BodyParser(project); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&project)
	c.JSON(project)
}
func DeleteProject(c *fiber.Ctx) {
	id := c.Params(("id"))
	db := database.DBConn

	var project Project
	db.First(&project, id)
	if project.Title == "" {
		c.Status(500).Send("No project found with given ID")
		return
	}
	db.Delete(&project)
	c.Send("Project successfully deleted.")
}
func AddLikeProject(c *fiber.Ctx) {
	id := c.Params(("id"))
	db := database.DBConn

	var project Project
	db.Find(&project, id)

	// TODO: update project: increment like

	c.JSON(project)
}
func RemoveLikeProject(c *fiber.Ctx) {
	id := c.Params(("id"))
	db := database.DBConn

	var project Project
	db.Find(&project, id)

	// TODO: update project: decrement like

	c.JSON(project)
}
