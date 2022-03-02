package server

import (
	"context"
	"log"

	"users_grpc/user"
	"users_grpc/usersPB"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (*Server) GetAllUsers(ctx context.Context, r *emptypb.Empty) (*usersPB.AllUsers, error) {
	log.Println("GetAllUsers() called")

	allUsers := user.GetAllUsers()

	toRespUsers := make([]*usersPB.User, len(allUsers))

	for i, userInstance := range allUsers {
		toRespUsers[i] = &usersPB.User{
			Id:      userInstance.Id,
			Fname:   userInstance.Fname,
			City:    userInstance.City,
			Phone:   userInstance.Phone,
			Height:  userInstance.Height,
			Married: userInstance.Married,
		}
	}

	return &usersPB.AllUsers{
		Users: toRespUsers,
	}, nil

}
