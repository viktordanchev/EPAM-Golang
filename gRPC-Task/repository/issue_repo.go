package repository

import "server/infrastructure/models"

type IssueRepository interface {
	CreateIssue(i *models.Issue) error
	UpdateIssue(i *models.Issue) error
	GetIssue(issueID string) (*models.Issue, error)
	DeleteIssue(issueID string) error
	GetAllIssues() ([]models.Issue, error)
}
