package services

import (
	"context"
	pb "server/gen/pb/issue"
	"server/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IssueService struct {
	pb.UnimplementedIssueServiceServer
}

func (s *IssueService) CreateIssue(ctx context.Context, req *pb.Issue) (*pb.Issue, error) {
	if req.Summary == "" || req.Description == "" || req.ProjectId == "" || req.Type == pb.IssueType_TYPE_UNSPECIFIED ||
		req.Priority == pb.IssuePriority_PRIORITY_UNSPECIFIED {
		return nil, status.Error(codes.InvalidArgument, "Missing fields")
	}

	issue := &pb.Issue{
		IssueId:        utils.GenerateId(),
		CreateDate:     timestamppb.Now(),
		ModifyDate:     timestamppb.Now(),
		Summary:        req.Summary,
		Description:    req.Description,
		Status:         pb.IssueStatus_NEW,
		Resolution:     pb.IssueResolution_RESOLUTION_UNSPECIFIED,
		Type:           req.Type,
		Priority:       req.Priority,
		ProjectId:      req.ProjectId,
		AssigneeUserId: req.AssigneeUserId,
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

	return req, nil
}

func (s *IssueService) GetIssue(ctx context.Context, req *pb.GetIssueRequest) (*pb.Issue, error) {
	return &pb.Issue{
		IssueId: req.IssueId,
		Summary: "example issue",
	}, nil
}

func (s *IssueService) DeleteIssue(ctx context.Context, req *pb.DeleteIssueRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *IssueService) ListIssues(ctx context.Context, _ *emptypb.Empty) (*pb.IssueList, error) {
	return &pb.IssueList{
		Issues: []*pb.Issue{},
	}, nil
}
