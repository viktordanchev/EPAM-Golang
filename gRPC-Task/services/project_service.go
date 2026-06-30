package services

import (
	"context"
	"errors"
	pb "server/gen/pb/project"
	"server/infrastructure/memory/models"
	"server/infrastructure/memory/repositories"
	"server/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProjectService struct {
	pb.UnimplementedProjectServiceServer
	repo *repositories.ProjectRepository
}

func NewProjectService(repo *repositories.ProjectRepository) *ProjectService {
	return &ProjectService{
		repo: repo,
	}
}

func (s *ProjectService) CreateProject(ctx context.Context, req *pb.Project) (*pb.Project, error) {
	if req.Name == "" || req.Description == "" {
		return nil, errors.New("Missing fields")
	}

	porjectModel := models.Project{
		ProjectId:   utils.GenerateId(),
		Name:        req.Name,
		Description: req.Description,
	}

	err := s.repo.CreateProject(porjectModel)
	if err != nil {
		return nil, err
	}

	project := &pb.Project{
		ProjectId:   porjectModel.ProjectId,
		Name:        porjectModel.Name,
		Description: porjectModel.Description,
	}

	return project, nil
}

func (s *ProjectService) UpdateProject(ctx context.Context, req *pb.Project) (*pb.Project, error) {
	if req.Name == "" || req.Description == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing fields")
	}

	porjectModel := models.Project{
		ProjectId:   req.ProjectId,
		Name:        req.Name,
		Description: req.Description,
	}

	err := s.repo.UpdateProject(porjectModel)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (s *ProjectService) GetProject(ctx context.Context, req *pb.GetProjectRequest) (*pb.Project, error) {
	p, err := s.repo.GetProject(req.ProjectId)
	if err != nil {
		return nil, err
	}

	project := &pb.Project{
		ProjectId:   p.ProjectId,
		Name:        p.Name,
		Description: p.Description,
	}

	return project, nil
}

func (s *ProjectService) DeleteProject(ctx context.Context, req *pb.DeleteProjectRequest) (*emptypb.Empty, error) {
	err := s.repo.DeleteProject(req.ProjectId)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ProjectService) ListProjects(ctx context.Context, _ *emptypb.Empty) (*pb.ProjectList, error) {
	projects, err := s.repo.GetAllProjects()
	if err != nil {
		return nil, err
	}

	projectList := &pb.ProjectList{}

	for _, p := range projects {
		projectList.Projects = append(projectList.Projects, &pb.Project{
			ProjectId:   p.ProjectId,
			Name:        p.Name,
			Description: p.Description,
		})
	}

	return projectList, nil
}
