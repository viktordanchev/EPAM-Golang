package repositories

import (
	"fmt"
	"server/infrastructure/memory/models"

	"github.com/hashicorp/go-memdb"
)

type IssueRepository struct {
	store *memdb.MemDB
}

func NewIssueRepository(store *memdb.MemDB) *IssueRepository {
	return &IssueRepository{
		store: store,
	}
}

func (r *IssueRepository) CreateIssue(i models.Issue) error {
	txn := r.store.Txn(true)
	defer txn.Abort()

	if err := txn.Insert("issue", i); err != nil {
		return fmt.Errorf("insert issue failed: %w", err)
	}

	txn.Commit()

	return nil
}

func (r *IssueRepository) UpdateIssue(i models.Issue) error {
	txn := r.store.Txn(true)
	defer txn.Abort()

	existing, err := txn.First("issue", "id", i.IssueId)
	if err != nil {
		return err
	}

	if existing == nil {
		return fmt.Errorf("Issue not found")
	}

	if err := txn.Insert("issue", i); err != nil {
		return err
	}

	txn.Commit()
	return nil
}

func (r *IssueRepository) GetIssue(issueID string) (models.Issue, error) {
	txn := r.store.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("issue", "id", issueID)
	if err != nil {
		return models.Issue{}, fmt.Errorf("Receive issue failed: %w", err)
	}

	if raw == nil {
		return models.Issue{}, fmt.Errorf("Issue not found")
	}

	issue := raw.(models.Issue)

	return issue, nil
}

func (r *IssueRepository) DeleteIssue(issueID string) error {
	txn := r.store.Txn(false)
	defer txn.Abort()

	issue, err := txn.First("issue", "id", issueID)
	if err != nil {
		return err
	}

	if issue == nil {
		return fmt.Errorf("Issue not found")
	}

	txn.Delete("issue", issue)
	return nil
}

func (r *IssueRepository) GetAllIssues() ([]models.Issue, error) {
	txn := r.store.Txn(false)
	defer txn.Abort()

	it, err := txn.Get("issue", "id")
	if err != nil {
		return nil, err
	}

	var issues []models.Issue

	for obj := it.Next(); obj != nil; obj = it.Next() {
		issues = append(issues, obj.(models.Issue))
	}

	return issues, nil
}
