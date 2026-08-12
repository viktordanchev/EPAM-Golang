package main

import (
	"log"
	"net"

	pbIssue "server/gen/pb/issue"
	pbProject "server/gen/pb/project"
	pbUser "server/gen/pb/user"
	"server/repository"

	memorydb "server/infrastructure/memory"
	memoryRepository "server/infrastructure/memory/repository"
	"server/service"

	"github.com/hashicorp/go-memdb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func main() {
	server := createMemoryServer()
	// server := createPostgresServer()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("gRPC server running on :50051")

	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

// ---------------------------------------------------------
// Memory
// ---------------------------------------------------------

func createMemoryServer() *grpc.Server {
	db := createMemoryDb()

	userRepo := memoryRepository.NewUserRepository(db)
	projectRepo := memoryRepository.NewProjectRepository(db)
	issueRepo := memoryRepository.NewIssueRepository(db)

	return createServer(
		userRepo,
		projectRepo,
		issueRepo,
	)
}

func createMemoryDb() *memdb.MemDB {
	db, err := memorydb.CreateMemoryStore()
	if err != nil {
		log.Fatal(err)
	}

	return db.GetStore()
}

// ---------------------------------------------------------
// PostgreSQL
// ---------------------------------------------------------

/*func createPostgresServer() *grpc.Server {
	db := createPostgresDb()

	userRepo := postgresRepository.NewUserRepository(db)
	projectRepo := postgresRepository.NewProjectRepository(db)
	issueRepo := postgresRepository.NewIssueRepository(db)

	return createServer(
		userRepo,
		projectRepo,
		issueRepo,
	)
}*/

func createPostgresDb() *gorm.DB {
	// Тук ще си сложиш твоята PostgreSQL/GORM конфигурация.
	//
	// Пример:
	//
	// db, err := postgres.NewPostgres(...)
	// if err != nil {
	//     log.Fatal(err)
	// }
	//
	// return db

	panic("PostgreSQL database is not configured yet")
}

// ---------------------------------------------------------
// Dependency Injection
// ---------------------------------------------------------

func createServer(
	userRepo repository.UserRepository,
	projectRepo repository.ProjectRepository,
	issueRepo repository.IssueRepository,
) *grpc.Server {
	userService := service.NewUserService(userRepo)
	projectService := service.NewProjectService(projectRepo)
	issueService := service.NewIssueService(issueRepo)

	s := grpc.NewServer()

	pbUser.RegisterUserServiceServer(s, userService)
	pbProject.RegisterProjectServiceServer(s, projectService)
	pbIssue.RegisterIssueServiceServer(s, issueService)

	return s
}
