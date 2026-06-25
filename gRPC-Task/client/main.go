package main

import (
	"context"
	"log"
	"time"

	pbProject "server/gen/pb/project"
	pbUser "server/gen/pb/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	userClient := pbUser.NewUserServiceClient(conn)
	projectClient := pbProject.NewProjectServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	userRes, err := userClient.CreateUser(ctx, &pbUser.User{
		FirstName:    "John",
		LastName:     "Doe",
		EmailAddress: "john@example.com",
	})
	if err != nil {
		log.Fatal(err)
	}

	projectRes, err := projectClient.CreateProject(ctx, &pbProject.Project{
		Name:        "TestProject",
		Description: "Run",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(userRes)
	log.Println(projectRes)
}
