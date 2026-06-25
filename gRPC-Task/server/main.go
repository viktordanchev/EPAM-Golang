package main

import (
	"log"
	"net"
	pbIssue "server/gen/pb/issue"
	pbProject "server/gen/pb/project"
	pbUser "server/gen/pb/user"
	"server/services"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pbIssue.RegisterIssueServiceServer(s, &services.IssueService{})
	pbProject.RegisterProjectServiceServer(s, &services.ProjectService{})
	pbUser.RegisterUserServiceServer(s, &services.UserService{})

	log.Println("gRPC server running on :50051")

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
