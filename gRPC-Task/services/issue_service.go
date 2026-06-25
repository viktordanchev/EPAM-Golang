package services

import (
	"context"
	pb "server/gen/pb/issue"

	"google.golang.org/protobuf/types/known/emptypb"
)

type IssueService struct {
	pb.UnimplementedIssueServiceServer
}

func (s *IssueService) CreateIssue(ctx context.Context, req *pb.Issue) (*pb.Issue, error) {
	// тук нормално: save to DB, generate ID, etc.

	req.IssueId = "generated-id"
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
