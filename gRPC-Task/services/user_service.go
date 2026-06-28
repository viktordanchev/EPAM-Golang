package services

import (
	"context"
	pb "server/gen/pb/user"
	"server/infrastructure/memory/models"
	"server/infrastructure/memory/repositories"
	"server/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	repo *repositories.UserRepository
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	if req.FirstName == "" || req.LastName == "" || req.EmailAddress == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing fields")
	}

	userModel := models.User{
		UserId:       utils.GenerateId(),
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		EmailAddress: req.EmailAddress,
	}

	s.repo.CreateUser(userModel)

	user := &pb.User{
		UserId:       userModel.UserId,
		FirstName:    userModel.FirstName,
		LastName:     userModel.LastName,
		EmailAddress: userModel.EmailAddress,
	}

	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	if req.FirstName == "" || req.LastName == "" || req.EmailAddress == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing fields")
	}

	userModel := models.User{
		UserId:       req.UserId,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		EmailAddress: req.EmailAddress,
	}

	err := s.repo.UpdateUser(userModel)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	return &pb.User{
		UserId: req.UserId,
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *UserService) ListUsers(ctx context.Context, _ *emptypb.Empty) (*pb.UserList, error) {
	return &pb.UserList{
		Users: []*pb.User{},
	}, nil
}
