package user

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (user *User) LoadDetails() error {

	for _, eachUser := range AllUsers {
		if eachUser.Id == user.Id {
			user.Fname = eachUser.Fname
			user.City = eachUser.City
			user.Phone = eachUser.Phone
			user.Height = eachUser.Height
			user.Married = eachUser.Married
			return nil
		}
	}

	if user.Fname == "" {
		return status.Error(codes.NotFound, "user not found")
	}

	return nil
}
