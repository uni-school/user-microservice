package service_user

import (
	"errors"

	"github.com/uni-school/user-microservice/libs/util"
	"github.com/uni-school/user-microservice/pkg/model"
	"github.com/uni-school/user-microservice/shared/constant"
	"gorm.io/gorm"

	pb "github.com/uni-school/user-microservice/proto"
)

func (s *UserService) CreateUser(payload *pb.CreateUserRequest) error {
	userModelGetResult, err := s.UserResource.GetUser(nil, &model.User{
		Email: payload.GetEmail(),
	})
	if err != nil {
		if err.Error() != gorm.ErrRecordNotFound.Error() {
			return err
		}
	}
	if userModelGetResult != nil {
		return errors.New("user already registered")
	}

	userModelCreatePayload, err := util.ConvertToStruct[model.User](model.User{
		Name:     payload.GetName(),
		Email:    payload.GetEmail(),
		Password: payload.GetPassword(),
		Role:     constant.UserRole(payload.GetRole()),
	})
	if err != nil {
		return err
	}

	err = s.UserResource.CreateUser(&userModelCreatePayload)
	if err != nil {
		return err
	}

	return nil
}
