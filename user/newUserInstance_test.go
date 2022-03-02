package user_test

import (
	"testing"

	"users_grpc/user"
)

func TestNewUser(t *testing.T) {

	_, err := user.NewUser(0)
	if err == nil {
		t.Errorf("Error user.NewUser(0) : id cannot be zero")
	}

	user2, err := user.NewUser(1)
	if err != nil {
		t.Errorf("Error user.NewUser(1) : %s", err)
	}
	if user2.Id != 1 {
		t.Errorf("Error user.NewUser(1) : id is not 1")
	}

}
