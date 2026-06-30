package main

import (
	"log"
	"net"
	pbUser "server/gen/pb/user"
	memorydb "server/infrastructure/memory"
	"server/infrastructure/memory/repositories"
	"server/services"

	"github.com/hashicorp/go-memdb"
	"google.golang.org/grpc"
)

func main() {
	db := createMemeryDb()
	userRepo := repositories.NewUserRepository(db)

	userService := services.NewUserService(userRepo)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pbUser.RegisterUserServiceServer(s, userService)

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
