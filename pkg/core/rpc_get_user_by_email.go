package core

import (
	"context"

	"github.com/uni-school/user-microservice/shared/custom"
	pb "github.com/uni-school/user-microservice/proto"
)

func (c *Server) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.GetUserByEmailResponse, error) {
	userProtoGetResponse, err := c.UserService.GetUserByEmail(req)
	if err != nil {
		return nil, custom.CustomGRPCErrorHandler(err)
	}

	return &pb.GetUserByEmailResponse{
		Message: "user found",
		Data:    userProtoGetResponse,
	}, nil
}
