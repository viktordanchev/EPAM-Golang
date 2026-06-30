package models

import "time"

type Status string
type Resolution string
type Type string
type Priority string

const (
	StatusUnspecified Status = "STATUS_UNSPECIFIED"
	StatusNew         Status = "NEW"
	StatusAssigned    Status = "ASSIGNED"
	StatusInProgress  Status = "IN_PROGRESS"
	StatusResolved    Status = "RESOLVED"
	StatusClosed      Status = "CLOSED"
	StatusReopened    Status = "REOPENED"
)

const (
	ResolutionUnspecified Resolution = "RESOLUTION_UNSPECIFIED"
	ResolutionFixed       Resolution = "FIXED"
	ResolutionInvalid     Resolution = "INVALID"
	ResolutionWontFix     Resolution = "WONTFIX"
	ResolutionWorksForMe  Resolution = "WORKSFORME"
)

const (
	TypeUnspecified Type = "TYPE_UNSPECIFIED"
	TypeCosmetic    Type = "COSMETIC"
	TypeBug         Type = "BUG"
	TypeFeature     Type = "FEATURE"
	TypePerformance Type = "PERFORMANCE"
)

const (
	PriorityUnspecified Priority = "PRIORITY_UNSPECIFIED"
	PriorityCritical    Priority = "CRITICAL"
	PriorityMajor       Priority = "MAJOR"
	PriorityImportant   Priority = "IMPORTANT"
	PriorityMinor       Priority = "MINOR"
)

type Issue struct {
	IssueId    string
	CreateDate time.Time
	ModifyDate time.Time

	Summary     string
	Description string

	Status         Status
	Resolution     Resolution
	Type           Type
	Priority       Priority
	ProjectId      string
	AssigneeUserId string
}
