package repository

import (
	"fmt"

	"server/infrastructure/models"

	"gorm.io/gorm"
)

type IssueRepository struct {
	db *gorm.DB
}

func NewIssueRepository(db *gorm.DB) *IssueRepository {
	return &IssueRepository{
		db: db,
	}
}

func (r *IssueRepository) CreateIssue(issue *models.Issue) error {
	if err := r.db.Create(&issue).Error; err != nil {
		return fmt.Errorf("create issue failed: %w", err)
	}

	return nil
}

func (r *IssueRepository) UpdateIssue(issue *models.Issue) error {
	result := r.db.
		Model(&models.Issue{}).
		Where("issue_id = ?", issue.IssueId).
		Updates(&issue)

	if result.Error != nil {
		return fmt.Errorf("update issue failed: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("issue not found")
	}

	return nil
}

func (r *IssueRepository) GetIssue(issueID string) (*models.Issue, error) {
	var issue models.Issue

	if err := r.db.
		Where("issue_id = ?", issueID).
		First(&issue).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("issue not found")
		}

		return nil, fmt.Errorf("get issue failed: %w", err)
	}

	return &issue, nil
}

func (r *IssueRepository) DeleteIssue(issueID string) error {
	result := r.db.
		Where("issue_id = ?", issueID).
		Delete(&models.Issue{})

	if result.Error != nil {
		return fmt.Errorf("delete issue failed: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("issue not found")
	}

	return nil
}

func (r *IssueRepository) GetAllIssues() ([]models.Issue, error) {
	var issues []models.Issue

	if err := r.db.
		Order("issue_id").
		Find(&issues).Error; err != nil {

		return nil, fmt.Errorf("get all issues failed: %w", err)
	}

	return issues, nil
}
