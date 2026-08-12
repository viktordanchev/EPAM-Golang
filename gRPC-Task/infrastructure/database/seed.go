package database

import (
	"server/infrastructure/database/models"
	"time"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	users := getUsers()

	for _, user := range users {
		db.FirstOrCreate(&user, models.User{ID: user.ID})
	}

	projects := getProjects()

	for _, project := range projects {
		db.FirstOrCreate(&project, models.Project{ID: project.ID})
	}

	issues := getIssues()

	for _, issue := range issues {
		db.FirstOrCreate(&issue, models.Issue{ID: issue.ID})
	}

	return nil
}

func getUsers() []models.User {
	data := []models.User{
		{
			ID:           "1",
			FirstName:    "Ivan",
			LastName:     "Petrov",
			EmailAddress: "ivan@test.com",
		},
		{
			ID:           "2",
			FirstName:    "Maria",
			LastName:     "Ivanova",
			EmailAddress: "maria@test.com",
		},
	}

	return data
}

func getProjects() []models.Project {
	data := []models.Project{
		{
			ID:          "1",
			Name:        "Project A",
			Description: "Description Project A",
		},
		{
			ID:          "2",
			Name:        "Project B",
			Description: "Description Project B",
		},
	}

	return data
}

func getIssues() []models.Issue {
	data := []models.Issue{
		{
			ID:             "1",
			Summary:        "Summary TEST",
			Description:    "Description Issue 1",
			Status:         "NEW",
			Resolution:     "WORKSFORME",
			Type:           "BUG",
			Priority:       "MAJOR",
			ProjectID:      "1",
			AssigneeUserID: "1",
			CreatedAt:      time.Now(),
		},
	}

	return data
}
