package services

import (
	"context"
	pb "server/gen/pb/issue"
	"server/infrastructure/memory/models"
	"server/infrastructure/memory/repositories"
	"server/utils"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IssueService struct {
	pb.UnimplementedIssueServiceServer
	repo *repositories.IssueRepository
}

func NewIssueService(repo *repositories.IssueRepository) *IssueService {
	return &IssueService{
		repo: repo,
	}
}

func (s *IssueService) CreateIssue(ctx context.Context, req *pb.Issue) (*pb.Issue, error) {
	if req.Summary == "" || req.Description == "" || req.ProjectId == "" || req.Type == pb.IssueType_TYPE_UNSPECIFIED ||
		req.Priority == pb.IssuePriority_PRIORITY_UNSPECIFIED {
		return nil, status.Error(codes.InvalidArgument, "Missing fields")
	}

	issueModel := models.Issue{
		IssueId:        utils.GenerateId(),
		CreateDate:     time.Now(),
		ModifyDate:     time.Now(),
		Summary:        req.Summary,
		Description:    req.Description,
		Status:         models.StatusNew,
		Resolution:     models.ResolutionUnspecified,
		Type:           mapTypeToModel(req.Type),
		Priority:       mapPriorityToModel(req.Priority),
		ProjectId:      req.ProjectId,
		AssigneeUserId: req.AssigneeUserId,
	}

	if req.AssigneeUserId != "" {
		issueModel.Status = models.StatusAssigned
	}

	err := s.repo.CreateIssue(issueModel)
	if err != nil {
		return nil, err
	}

	issue := &pb.Issue{
		IssueId:        issueModel.IssueId,
		CreateDate:     timestamppb.New(issueModel.CreateDate),
		ModifyDate:     timestamppb.New(issueModel.ModifyDate),
		Summary:        issueModel.Summary,
		Description:    issueModel.Description,
		Status:         pb.IssueStatus_NEW,
		Resolution:     pb.IssueResolution_RESOLUTION_UNSPECIFIED,
		Type:           req.Type,
		Priority:       req.Priority,
		ProjectId:      issueModel.ProjectId,
		AssigneeUserId: issueModel.AssigneeUserId,
	}

	if req.AssigneeUserId != "" {
		issue.Status = pb.IssueStatus_ASSIGNED
	}

	return issue, nil
}

func (s *IssueService) UpdateIssue(ctx context.Context, req *pb.Issue) (*pb.Issue, error) {
	if req.IssueId == "" ||
		req.Summary == "" ||
		req.Description == "" ||
		req.ProjectId == "" ||
		req.Type == pb.IssueType_TYPE_UNSPECIFIED ||
		req.Priority == pb.IssuePriority_PRIORITY_UNSPECIFIED ||
		req.Status == pb.IssueStatus_STATUS_UNSPECIFIED {

		return nil, status.Error(codes.InvalidArgument, "Missing fields")
	}

	if req.Status == pb.IssueStatus_ASSIGNED && req.AssigneeUserId == "" {
		return nil, status.Error(codes.InvalidArgument, "Assignee required for ASSIGNED status")
	}

	if (req.Status == pb.IssueStatus_RESOLVED || req.Status == pb.IssueStatus_CLOSED) && req.Resolution == pb.IssueResolution_RESOLUTION_UNSPECIFIED {
		return nil, status.Error(codes.InvalidArgument, "Resolution required for CLOSED/RESOLVED issues")
	} else {
		req.Resolution = pb.IssueResolution_RESOLUTION_UNSPECIFIED
	}

	switch req.Status {
	case pb.IssueStatus_IN_PROGRESS,
		pb.IssueStatus_RESOLVED,
		pb.IssueStatus_CLOSED:
	default:
		return nil, status.Error(codes.InvalidArgument, "Invalid status")
	}

	issueModel := models.Issue{
		IssueId:        req.IssueId,
		ModifyDate:     time.Now(),
		Summary:        req.Summary,
		Description:    req.Description,
		Status:         mapStatusToModel(req.Status),
		Resolution:     mapResolutionToModel(req.Resolution),
		Type:           mapTypeToModel(req.Type),
		Priority:       mapPriorityToModel(req.Priority),
		ProjectId:      req.ProjectId,
		AssigneeUserId: req.AssigneeUserId,
	}

	err := s.repo.UpdateIssue(issueModel)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (s *IssueService) GetIssue(ctx context.Context, req *pb.GetIssueRequest) (*pb.Issue, error) {
	i, err := s.repo.GetIssue(req.IssueId)
	if err != nil {
		return nil, err
	}

	issue := &pb.Issue{
		IssueId:        i.IssueId,
		CreateDate:     timestamppb.New(i.CreateDate),
		ModifyDate:     timestamppb.New(i.ModifyDate),
		Summary:        i.Summary,
		Description:    i.Description,
		Status:         mapStatusToPB(i.Status),
		Resolution:     mapResolutionToPB(i.Resolution),
		Type:           mapTypeToPB(i.Type),
		Priority:       mapPriorityToPB(i.Priority),
		ProjectId:      i.ProjectId,
		AssigneeUserId: i.AssigneeUserId,
	}

	return issue, nil
}

func (s *IssueService) DeleteIssue(ctx context.Context, req *pb.DeleteIssueRequest) (*emptypb.Empty, error) {
	err := s.repo.DeleteIssue(req.IssueId)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *IssueService) ListIssues(ctx context.Context, _ *emptypb.Empty) (*pb.IssueList, error) {
	issues, err := s.repo.GetAllIssues()
	if err != nil {
		return nil, err
	}

	issuesList := &pb.IssueList{}

	for _, i := range issues {
		issuesList.Issues = append(issuesList.Issues, &pb.Issue{
			IssueId:        i.IssueId,
			CreateDate:     timestamppb.New(i.CreateDate),
			ModifyDate:     timestamppb.New(i.ModifyDate),
			Summary:        i.Summary,
			Description:    i.Description,
			Status:         mapStatusToPB(i.Status),
			Resolution:     mapResolutionToPB(i.Resolution),
			Type:           mapTypeToPB(i.Type),
			Priority:       mapPriorityToPB(i.Priority),
			ProjectId:      i.ProjectId,
			AssigneeUserId: i.AssigneeUserId,
		})
	}

	return issuesList, nil
}

func mapTypeToModel(t pb.IssueType) models.Type {
	switch t {
	case pb.IssueType_BUG:
		return models.TypeBug
	case pb.IssueType_FEATURE:
		return models.TypeFeature
	case pb.IssueType_COSMETIC:
		return models.TypeCosmetic
	case pb.IssueType_PERFORMANCE:
		return models.TypePerformance
	default:
		return models.TypeUnspecified
	}
}

func mapPriorityToModel(p pb.IssuePriority) models.Priority {
	switch p {
	case pb.IssuePriority_CRITICAL:
		return models.PriorityCritical
	case pb.IssuePriority_MAJOR:
		return models.PriorityMajor
	case pb.IssuePriority_IMPORTANT:
		return models.PriorityImportant
	case pb.IssuePriority_MINOR:
		return models.PriorityMinor
	default:
		return models.PriorityUnspecified
	}
}

func mapStatusToModel(p pb.IssueStatus) models.Status {
	switch p {
	case pb.IssueStatus_ASSIGNED:
		return models.StatusAssigned
	case pb.IssueStatus_CLOSED:
		return models.StatusClosed
	case pb.IssueStatus_IN_PROGRESS:
		return models.StatusInProgress
	case pb.IssueStatus_NEW:
		return models.StatusNew
	case pb.IssueStatus_REOPENED:
		return models.StatusReopened
	case pb.IssueStatus_RESOLVED:
		return models.StatusResolved
	default:
		return models.StatusUnspecified
	}
}

func mapResolutionToModel(p pb.IssueResolution) models.Resolution {
	switch p {
	case pb.IssueResolution_FIXED:
		return models.ResolutionFixed
	case pb.IssueResolution_INVALID:
		return models.ResolutionInvalid
	case pb.IssueResolution_WONTFIX:
		return models.ResolutionWontFix
	case pb.IssueResolution_WORKSFORME:
		return models.ResolutionWorksForMe
	default:
		return models.ResolutionUnspecified
	}
}

func mapTypeToPB(t models.Type) pb.IssueType {
	switch t {
	case models.TypeBug:
		return pb.IssueType_BUG
	case models.TypeFeature:
		return pb.IssueType_FEATURE
	case models.TypeCosmetic:
		return pb.IssueType_COSMETIC
	case models.TypePerformance:
		return pb.IssueType_PERFORMANCE
	default:
		return pb.IssueType_TYPE_UNSPECIFIED
	}
}

func mapPriorityToPB(p models.Priority) pb.IssuePriority {
	switch p {
	case models.PriorityCritical:
		return pb.IssuePriority_CRITICAL
	case models.PriorityMajor:
		return pb.IssuePriority_MAJOR
	case models.PriorityImportant:
		return pb.IssuePriority_IMPORTANT
	case models.PriorityMinor:
		return pb.IssuePriority_MINOR
	default:
		return pb.IssuePriority_PRIORITY_UNSPECIFIED
	}
}

func mapStatusToPB(s models.Status) pb.IssueStatus {
	switch s {
	case models.StatusAssigned:
		return pb.IssueStatus_ASSIGNED
	case models.StatusClosed:
		return pb.IssueStatus_CLOSED
	case models.StatusInProgress:
		return pb.IssueStatus_IN_PROGRESS
	case models.StatusNew:
		return pb.IssueStatus_NEW
	case models.StatusReopened:
		return pb.IssueStatus_REOPENED
	case models.StatusResolved:
		return pb.IssueStatus_RESOLVED
	default:
		return pb.IssueStatus_STATUS_UNSPECIFIED
	}
}

func mapResolutionToPB(r models.Resolution) pb.IssueResolution {
	switch r {
	case models.ResolutionFixed:
		return pb.IssueResolution_FIXED
	case models.ResolutionInvalid:
		return pb.IssueResolution_INVALID
	case models.ResolutionWontFix:
		return pb.IssueResolution_WONTFIX
	case models.ResolutionWorksForMe:
		return pb.IssueResolution_WORKSFORME
	default:
		return pb.IssueResolution_RESOLUTION_UNSPECIFIED
	}
}
