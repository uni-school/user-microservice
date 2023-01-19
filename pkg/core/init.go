package core

import (
	service_user "github.com/uni-school/user-microservice/pkg/core/user/service"
	pb "github.com/uni-school/user-microservice/proto"
	"github.com/uni-school/user-microservice/wire"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	UserService service_user.IUserService
}

func InitServer() *Server {
	return &Server{
		UnimplementedUserServiceServer: pb.UnimplementedUserServiceServer{},
		UserService:             wire.UserService(),
	}
}
