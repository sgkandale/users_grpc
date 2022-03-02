package user

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (user *User) LoadDetails() error {

	for _, user := range AllUsers {
		if user.Id == user.Id {
			user.Fname = user.Fname
			user.City = user.City
			user.Phone = user.Phone
			user.Height = user.Height
			user.Married = user.Married
		}
	}

	if user.Fname == "" {
		return status.Error(codes.NotFound, "user not found")
	}

	return nil
}
