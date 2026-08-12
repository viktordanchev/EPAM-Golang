package repository

import (
	"fmt"

	"server/infrastructure/models"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		db: db,
	}
}

func (r *ProjectRepository) CreateProject(p *models.Project) error {
	if err := r.db.Create(p).Error; err != nil {
		return fmt.Errorf("create project failed: %w", err)
	}

	return nil
}

func (r *ProjectRepository) UpdateProject(p *models.Project) error {
	result := r.db.
		Model(&models.Project{}).
		Where("project_id = ?", p.ProjectId).
		Updates(p)

	if result.Error != nil {
		return fmt.Errorf("update project failed: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("project not found")
	}

	return nil
}

func (r *ProjectRepository) GetProject(projectID string) (*models.Project, error) {
	var project models.Project

	if err := r.db.
		Where("project_id = ?", projectID).
		First(&project).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("project not found")
		}

		return nil, fmt.Errorf("get project failed: %w", err)
	}

	return &project, nil
}

func (r *ProjectRepository) DeleteProject(projectID string) error {
	result := r.db.
		Where("project_id = ?", projectID).
		Delete(&models.Project{})

	if result.Error != nil {
		return fmt.Errorf("delete project failed: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("project not found")
	}

	return nil
}

func (r *ProjectRepository) GetAllProjects() ([]models.Project, error) {
	var projects []models.Project

	if err := r.db.
		Order("project_id").
		Find(&projects).Error; err != nil {

		return nil, fmt.Errorf("get all projects failed: %w", err)
	}

	return projects, nil
}
