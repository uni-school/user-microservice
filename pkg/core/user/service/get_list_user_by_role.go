package service_user

import (
	"github.com/uni-school/user-microservice/libs/types"
	"github.com/uni-school/user-microservice/libs/util"
	"github.com/uni-school/user-microservice/pkg/model"
	pb "github.com/uni-school/user-microservice/proto"
	"github.com/uni-school/user-microservice/shared/constant"
	"github.com/uni-school/user-microservice/shared/custom"
	"google.golang.org/grpc/codes"
)

func (s *UserService) GetListUserByRole(payload *pb.GetListUserByRoleRequest) ([]*pb.GetListUserByRoleResponseData, error) {
	userModelGetListResult, err := s.UserResource.GetListUser(&types.Query{
		SelectColumns: []string{
			"id",
			"name",
			"email",
			"role",
		},
	}, &model.User{
		Role: constant.UserRole(payload.GetRole()),
	})
	if err != nil {
		return nil, err
	}

	if len(userModelGetListResult) < 1 {
		return nil, custom.NewCustomGRPCError(codes.NotFound, "users not found")
	}

	userProtoGetListResponseData := make([]*pb.GetListUserByRoleResponseData, 0)
	for _, value := range userModelGetListResult {
		data, err := util.ConvertToStruct[*pb.GetListUserByRoleResponseData](value)
		if err != nil {
			return nil, err
		}
		userProtoGetListResponseData = append(userProtoGetListResponseData, data)
	}


	return userProtoGetListResponseData, nil
}
