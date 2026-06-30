package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"server/client/calls"
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

	calls.MakeUserCalls(ctx, userClient)
	fmt.Println()
	calls.MakeProjectCalls(ctx, projectClient)
}
