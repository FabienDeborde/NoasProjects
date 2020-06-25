package usecase

import (
	entity "github.com/FabienDeborde/noas_projects/domain/entity"
	database "github.com/FabienDeborde/noas_projects/infrastructure"
)

// GetProjects is GetProjects
func GetProjects() []entity.Project {
	// return repository.GetProjects
	db := database.DBConn
	var projects []entity.Project
	db.Find((&projects))

	return projects

}
