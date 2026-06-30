package main

import (
	"log"
	"net"
	pbUser "server/gen/pb/user"
	memdb "server/infrastructure/memory"
	"server/infrastructure/memory/repositories"
	"server/services"

	"google.golang.org/grpc"
)

func main() {
	db, err := memdb.CreateMemoryStore()
	if err != nil {
		panic(err)
	}

	memdb := db.GetStore()
	userRepo := repositories.NewUserRepository(memdb)

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
