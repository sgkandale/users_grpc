package server

import (
	"context"
	"log"

	"users_grpc/user"
	"users_grpc/usersPB"
)

func (*Server) GetUser(ctx context.Context, r *usersPB.GetUser_Request) (*usersPB.User, error) {
	log.Println("GetUser() called")

	userInstance, err := user.NewUser(r.Id)
	if err != nil {
		return nil, err
	}

	err = userInstance.LoadDetails()
	if err != nil {
		return nil, err
	}

	return &usersPB.User{
		Id:      userInstance.Id,
		Fname:   userInstance.Fname,
		City:    userInstance.City,
		Phone:   userInstance.Phone,
		Height:  userInstance.Height,
		Married: userInstance.Married,
	}, nil

}
