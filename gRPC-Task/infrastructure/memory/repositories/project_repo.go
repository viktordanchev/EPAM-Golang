package repositories

import (
	"fmt"
	"server/infrastructure/memory/models"

	"github.com/hashicorp/go-memdb"
)

type ProjectRepository struct {
	store *memdb.MemDB
}

func NewProjectRepository(store *memdb.MemDB) *ProjectRepository {
	return &ProjectRepository{
		store: store,
	}
}

func (r *ProjectRepository) CreateProject(p models.Project) error {
	txn := r.store.Txn(true)
	defer txn.Abort()

	if err := txn.Insert("project", p); err != nil {
		return fmt.Errorf("insert project failed: %w", err)
	}

	txn.Commit()

	return nil
}

func (r *ProjectRepository) UpdateProject(p models.Project) error {
	txn := r.store.Txn(true)
	defer txn.Abort()

	existing, err := txn.First("user", "id", p.ProjectId)
	if err != nil {
		panic(err)
	}

	if existing == nil {
		return fmt.Errorf("Project not found")
	}

	if err := txn.Insert("project", p); err != nil {
		panic(err)
	}

	txn.Commit()
	return nil
}

func (r *ProjectRepository) GetProject(projectID string) (models.Project, error) {
	txn := r.store.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("project", "id", projectID)
	if err != nil {
		return models.Project{}, fmt.Errorf("Receive project failed: %w", err)
	}

	if raw == nil {
		return models.Project{}, fmt.Errorf("Project not found")
	}

	project := raw.(models.Project)

	return project, nil
}

func (r *ProjectRepository) DeleteProject(projectID string) error {
	txn := r.store.Txn(false)
	defer txn.Abort()

	project, err := txn.First("project", "id", projectID)
	if err != nil {
		panic(err)
	}

	if project == nil {
		return fmt.Errorf("Project not found")
	}

	txn.Delete("project", project)
	return nil
}

func (r *ProjectRepository) GetAllProjects() ([]models.Project, error) {
	txn := r.store.Txn(false)
	defer txn.Abort()

	it, err := txn.Get("project", "id")
	if err != nil {
		panic(err)
	}

	var projects []models.Project

	for obj := it.Next(); obj != nil; obj = it.Next() {
		projects = append(projects, obj.(models.Project))
	}

	return projects, nil
}
