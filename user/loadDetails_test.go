package user_test

import (
	"testing"

	"users_grpc/user"
)

func TestLoadDetails(t *testing.T) {

	givenUser := &user.User{Id: 1}

	err := givenUser.LoadDetails()
	if err != nil {
		t.Errorf("Error user.LoadDetails() : %s", err)
	}

	if givenUser.Fname == "" {
		t.Errorf("Error user.LoadDetails() : Fname is empty")
	}
	if givenUser.City == "" {
		t.Errorf("Error user.LoadDetails() : City is empty")
	}
	if givenUser.Phone == 0 {
		t.Errorf("Error user.LoadDetails() : Phone is empty")
	}
	if givenUser.Height == 0 {
		t.Errorf("Error user.LoadDetails() : Height is empty")
	}
}
