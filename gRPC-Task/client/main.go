package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pbUser "server/gen/pb/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	userRes, err := userClient.CreateUser(ctx, &pbUser.User{
		FirstName:    "John",
		LastName:     "Doe",
		EmailAddress: "john@example2.com",
	})
	if err != nil {
		log.Fatal(err)
	}

	userRes2, err := userClient.CreateUser(ctx, &pbUser.User{
		FirstName:    "Viktor",
		LastName:     "Danchev",
		EmailAddress: "viktor@example.com",
	})
	if err != nil {
		log.Fatal(err)
	}

	user, err := userClient.GetUser(ctx, &pbUser.GetUserRequest{
		UserId: userRes.UserId,
	})
	if err != nil {
		log.Fatal(err)
	}

	user2, err := userClient.GetUser(ctx, &pbUser.GetUserRequest{
		UserId: userRes2.UserId,
	})
	if err != nil {
		log.Fatal(err)
	}

	users, err := userClient.ListUsers(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("CREATED:", user)
	fmt.Println("CREATED:", user2)
	fmt.Println("LIST:", users)

	deleted, err := userClient.DeleteUser(ctx, &pbUser.DeleteUserRequest{
		UserId: userRes.UserId,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DELETED:", deleted)

	users2, err := userClient.ListUsers(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("LIST:", users2)
}
