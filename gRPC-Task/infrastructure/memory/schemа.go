package memory

import (
	"time"

	"github.com/hashicorp/go-memdb"
)

type User struct {
	UserId       string
	FirstName    string
	LastName     string
	EmailAddress string
}

type Project struct {
	ProjectId   string
	Name        string
	Description string
}

type Issue struct {
	IssueId    string
	CreateDate time.Time
	ModifyDate time.Time

	Summary     string
	Description string

	IssueStatus     string
	IssueResolution string
	IssueType       string
	IssuePriority   string
	ProjectId       string
	AssigneeUserId  string
}

var Schema = &memdb.DBSchema{
	Tables: map[string]*memdb.TableSchema{

		"user": {
			Name: "user",
			Indexes: map[string]*memdb.IndexSchema{
				"id": {
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.StringFieldIndex{Field: "UserId"},
				},
				"firstName": {
					Name:    "firstName",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "FirstName"},
				},
				"lastName": {
					Name:    "lastName",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "LastName"},
				},
				"emailAddress": {
					Name:    "emailAddress",
					Unique:  true,
					Indexer: &memdb.StringFieldIndex{Field: "EmailAddress"},
				},
			},
		},

		"project": {
			Name: "project",
			Indexes: map[string]*memdb.IndexSchema{
				"id": {
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.StringFieldIndex{Field: "ProjectId"},
				},
				"name": {
					Name:    "name",
					Unique:  true,
					Indexer: &memdb.StringFieldIndex{Field: "Name"},
				},
				"description": {
					Name:    "description",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Description"},
				},
			},
		},

		"issue": {
			Name: "issue",
			Indexes: map[string]*memdb.IndexSchema{
				"id": {
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.StringFieldIndex{Field: "IssueId"},
				},
				"summary": {
					Name:    "summary",
					Indexer: &memdb.StringFieldIndex{Field: "Summary"},
				},
				"description": {
					Name:    "description",
					Indexer: &memdb.StringFieldIndex{Field: "Description"},
				},
				"issueStatus": {
					Name:    "issueStatus",
					Indexer: &memdb.StringFieldIndex{Field: "IssueStatus"},
				},
				"issuePriority": {
					Name:    "issuePriority",
					Indexer: &memdb.StringFieldIndex{Field: "IssuePriority"},
				},
				"issueType": {
					Name:    "issueType",
					Indexer: &memdb.StringFieldIndex{Field: "IssueType"},
				},
				"issueResolution": {
					Name:    "issueResolution",
					Indexer: &memdb.StringFieldIndex{Field: "IssueResolution"},
				},
				"projectId": {
					Name:    "projectId",
					Indexer: &memdb.StringFieldIndex{Field: "ProjectId"},
				},
				"assigneeUserId": {
					Name:    "assigneeUserId",
					Indexer: &memdb.StringFieldIndex{Field: "AssigneeUserId"},
				},
				"createDate": {
					Name:    "createDate",
					Indexer: &memdb.IntFieldIndex{Field: "CreateDate"},
				},
				"modifyDate": {
					Name:    "modifyDate",
					Indexer: &memdb.IntFieldIndex{Field: "ModifyDate"},
				},
			},
		},
	},
}
