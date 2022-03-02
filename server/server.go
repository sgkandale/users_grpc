package server

import (
	"users_grpc/usersPB"
)

//Server exposed
type Server struct {
	usersPB.UnimplementedUsersServiceServer
}
