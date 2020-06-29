package repository

import (
	"github.com/FabienDeborde/noas_projects/app/domain/entity"
	"github.com/FabienDeborde/noas_projects/app/infrastructure/database"
)

func GetProjects() *[]entity.Project {
	db := database.DBConn
	var projects []entity.Project
	db.Find((&projects))

	return &projects
}
