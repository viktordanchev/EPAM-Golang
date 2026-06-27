package models

import "time"

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
