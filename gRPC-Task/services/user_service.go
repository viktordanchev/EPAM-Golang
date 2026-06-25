package services

import (
	"context"
	pb "server/gen/pb/user"

	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	user := &pb.User{
		UserId:       "generated-id",
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		EmailAddress: req.EmailAddress,
	}

	return user, nil
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
