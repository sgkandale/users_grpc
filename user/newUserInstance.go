package user

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewUser(id int32) (*User, error) {

	if id == 0 {
		return nil, status.Error(codes.InvalidArgument, "id cannot be zero")
	}

	return &User{
		Id: id,
	}, nil
}
