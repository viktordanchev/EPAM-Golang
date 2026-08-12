package database

import (
	"server/infrastructure/models"
	"time"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	users := getUsers()

	for _, user := range users {
		db.FirstOrCreate(&user, models.User{UserId: user.UserId})
	}

	projects := getProjects()

	for _, project := range projects {
		db.FirstOrCreate(&project, models.Project{ProjectId: project.ProjectId})
	}

	issues := getIssues()

	for _, issue := range issues {
		db.FirstOrCreate(&issue, models.Issue{IssueId: issue.IssueId})
	}

	return nil
}

func getUsers() []models.User {
	data := []models.User{
		{
			UserId:       "1",
			FirstName:    "Ivan",
			LastName:     "Petrov",
			EmailAddress: "ivan@test.com",
		},
		{
			UserId:       "2",
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
			ProjectId:   "1",
			Name:        "Project A",
			Description: "Description Project A",
		},
		{
			ProjectId:   "2",
			Name:        "Project B",
			Description: "Description Project B",
		},
	}

	return data
}

func getIssues() []models.Issue {
	data := []models.Issue{
		{
			IssueId:        "1",
			Summary:        "Summary TEST",
			Description:    "Description Issue 1",
			Status:         "NEW",
			Resolution:     "WORKSFORME",
			Type:           "BUG",
			Priority:       "MAJOR",
			ProjectId:      "1",
			AssigneeUserId: "1",
			CreateDate:     time.Now(),
		},
	}

	return data
}
