package server

import (
	"context"
	"log"

	"users_grpc/user"
	"users_grpc/usersPB"
)

func (*Server) GetUser(ctx context.Context, r *usersPB.GetUser_Request) (*usersPB.User, error) {
	log.Println("GetUser() called")

	userInstace, err := user.NewUser(r.Id)
	if err != nil {
		return nil, err
	}

	err = userInstace.LoadDetails()
	if err != nil {
		return nil, err
	}

	return &usersPB.User{
		Id:      userInstace.Id,
		Fname:   userInstace.Fname,
		City:    userInstace.City,
		Phone:   userInstace.Phone,
		Height:  userInstace.Height,
		Married: userInstace.Married,
	}, nil

}
