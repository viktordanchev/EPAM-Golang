package main

import (
	"log"
	"net"
	pbIssue "server/gen/pb/issue"
	pbProject "server/gen/pb/project"
	pbUser "server/gen/pb/user"
	memorydb "server/infrastructure/memory"
	"server/infrastructure/memory/repository"
	"server/services"

	"github.com/hashicorp/go-memdb"
	"google.golang.org/grpc"
)

func main() {
	db := createMemeryDb()
	userRepo := repository.NewUserRepository(db)
	projectRepo := repository.NewProjectRepository(db)
	issueRepo := repository.NewIssueRepository(db)

	userService := services.NewUserService(userRepo)
	projectService := services.NewProjectService(projectRepo)
	issueService := services.NewIssueService(issueRepo)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pbUser.RegisterUserServiceServer(s, userService)
	pbProject.RegisterProjectServiceServer(s, projectService)
	pbIssue.RegisterIssueServiceServer(s, issueService)

	log.Println("gRPC server running on :50051")

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func createMemeryDb() *memdb.MemDB {
	db, err := memorydb.CreateMemoryStore()
	if err != nil {
		panic(err)
	}

	memdb := db.GetStore()

	return memdb
}
