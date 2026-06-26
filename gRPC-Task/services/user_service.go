package services

import (
	"context"
	pb "server/gen/pb/user"
	"server/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	if req.FirstName == "" || req.LastName == "" || req.EmailAddress == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing fields")
	}

	user := &pb.User{
		UserId:       utils.GenerateId(),
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		EmailAddress: req.EmailAddress,
	}

	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	if req.FirstName == "" || req.LastName == "" || req.EmailAddress == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing fields")
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
