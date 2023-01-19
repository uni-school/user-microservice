package service_user

import (
	"github.com/uni-school/user-microservice/libs/util"
	"github.com/uni-school/user-microservice/pkg/model"
	pb "github.com/uni-school/user-microservice/proto"
)

func (s *UserService) GetUserByEmail(payload *pb.GetUserByEmailRequest) (*pb.GetUserByEmailResponseData, error) {
	userModelGetPayload, err := util.ConvertToStruct[model.User](model.User{
		Email:    payload.GetEmail(),
	})
	if err != nil {
		return nil, err
	}

	userModelGetResult, err := s.UserResource.GetUser(nil, &userModelGetPayload)
	if err != nil {
		return nil, err
	}

	userProtoGetResponse, err := util.ConvertToStruct[pb.GetUserByEmailResponseData](*userModelGetResult)
	if err != nil {
		return nil, err
	}

	return &userProtoGetResponse, nil
}
