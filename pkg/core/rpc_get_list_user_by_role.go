package core

import (
	"context"

	"github.com/uni-school/user-microservice/shared/custom"
	pb "github.com/uni-school/user-microservice/proto"
)

func (c *Server) GetListUserByRole(ctx context.Context, req *pb.GetListUserByRoleRequest) (*pb.GetListUserByRoleResponse, error) {
	userProtoGetListResponse, err := c.UserService.GetListUserByRole(req)
	if err != nil {
		return nil, custom.CustomGRPCErrorHandler(err)
	}

	return &pb.GetListUserByRoleResponse{
		Message: "users found",
		Data:    userProtoGetListResponse,
	}, nil
}
