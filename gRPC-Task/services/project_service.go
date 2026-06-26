package services

import (
	"context"
	"errors"
	pb "server/gen/pb/project"
	"server/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProjectService struct {
	pb.UnimplementedProjectServiceServer
}

func (s *ProjectService) CreateProject(ctx context.Context, req *pb.Project) (*pb.Project, error) {
	if req.Name == "" || req.Description == "" {
		return nil, errors.New("Missing fields")
	}

	project := &pb.Project{
		ProjectId:   utils.GenerateId(),
		Name:        req.Name,
		Description: req.Description,
	}

	return project, nil
}

func (s *ProjectService) UpdateProject(ctx context.Context, req *pb.Project) (*pb.Project, error) {
	if req.Name == "" || req.Description == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing fields")
	}

	return req, nil
}

func (s *ProjectService) GetProject(ctx context.Context, req *pb.GetProjectRequest) (*pb.Project, error) {
	return &pb.Project{
		ProjectId: req.ProjectId,
	}, nil
}

func (s *ProjectService) DeleteProject(ctx context.Context, req *pb.DeleteProjectRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *ProjectService) ListProjects(ctx context.Context, _ *emptypb.Empty) (*pb.ProjectList, error) {
	return &pb.ProjectList{
		Projects: []*pb.Project{},
	}, nil
}
