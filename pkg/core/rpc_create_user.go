package core

import (
	"context"

	"github.com/uni-school/user-microservice/libs/filter"
	pb "github.com/uni-school/user-microservice/proto"
)

func (c *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if err := c.UserService.CreateUser(req); err != nil {
		return nil, filter.CustomGRPCErrorHandler(err)
	}

	return &pb.CreateUserResponse{
		Message: "user created", 
		}, nil
}
