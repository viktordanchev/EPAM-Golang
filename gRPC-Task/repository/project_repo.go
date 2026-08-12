package repository

import "server/infrastructure/models"

type ProjectRepository interface {
	CreateProject(p *models.Project) error
	UpdateProject(p *models.Project) error
	GetProject(projectID string) (*models.Project, error)
	DeleteProject(projectID string) error
	GetAllProjects() ([]models.Project, error)
}
