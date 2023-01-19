package core

import (
	"context"

	"github.com/uni-school/user-microservice/libs/filter"
	pb "github.com/uni-school/user-microservice/proto"
)

func (c *Server) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.GetUserByEmailResponse, error) {
	userProtoGetResponse, err := c.UserService.GetUserByEmail(req)
	if err != nil {
		return nil, filter.CustomGRPCErrorHandler(err)
	}

	return &pb.GetUserByEmailResponse{
		Message: "user found",
		Data:    userProtoGetResponse,
	}, nil
}
