package usecase

import (
	entity "github.com/FabienDeborde/noas_projects/app/domain/entity"
	"github.com/FabienDeborde/noas_projects/app/domain/repository"
)

// GetProjects is GetProjects
func GetProjects() *[]entity.Project {
	return repository.GetProjects
}
