package calls

import (
	"context"
	"fmt"
	"log"
	pbUser "server/gen/pb/user"

	"google.golang.org/protobuf/types/known/emptypb"
)

func MakeUserCalls(ctx context.Context, client pbUser.UserServiceClient) {
	user := createUser(ctx, client, "Viktor", "Danchev", "vitkor@abv.bg")
	user2 := createUser(ctx, client, "Mitko", "Ivanov", "mitko@abv.bg")

	getUser(ctx, client, user.UserId)
	getUser(ctx, client, user2.UserId)

	listUsers(ctx, client)

	deleteUser(ctx, client, user.UserId)

	listUsers(ctx, client)
	updateUser(ctx, client, user2)
}

func createUser(ctx context.Context, client pbUser.UserServiceClient, first, last, email string) *pbUser.User {
	user, err := client.CreateUser(ctx, &pbUser.User{
		FirstName:    first,
		LastName:     last,
		EmailAddress: email,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("CREATED USER:", user)
	return user
}

func updateUser(ctx context.Context, client pbUser.UserServiceClient, user *pbUser.User) *pbUser.User {
	user, err := client.UpdateUser(ctx, &pbUser.User{
		UserId:       user.UserId,
		FirstName:    "UPDATED",
		LastName:     user.LastName,
		EmailAddress: user.EmailAddress,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("UPDATED USER:", user)
	return user
}

func getUser(ctx context.Context, client pbUser.UserServiceClient, userID string) {
	user, err := client.GetUser(ctx, &pbUser.GetUserRequest{
		UserId: userID,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GET USER:", user)
}

func deleteUser(ctx context.Context, client pbUser.UserServiceClient, userID string) {
	res, err := client.DeleteUser(ctx, &pbUser.DeleteUserRequest{
		UserId: userID,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DELETED USER:", res, userID)
}

func listUsers(ctx context.Context, client pbUser.UserServiceClient) {
	users, err := client.ListUsers(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("LIST USERS:", users)
}
