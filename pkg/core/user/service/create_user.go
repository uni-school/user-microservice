package service_user

import (
	libs_util "github.com/uni-school/user-microservice/libs/util"
	"github.com/uni-school/user-microservice/pkg/model"
	"github.com/uni-school/user-microservice/shared/constant"
	"github.com/uni-school/user-microservice/shared/custom"
	shared_util "github.com/uni-school/user-microservice/shared/util"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"

	pb "github.com/uni-school/user-microservice/proto"
)

type (
	CreateUserRequestPbToUser struct {
		Name     string            `json:"name"`
		Email    string            `json:"email"`
		Password string            `json:"password"`
		Role     constant.UserRole `json:"role"`
	}
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
		return custom.NewCustomGRPCError(codes.AlreadyExists, "user already registered")
	}

	var (
		hashedPassword string
	)
	if payload.GetPassword() != "" {
		hashedPassword, err = shared_util.HashPassword(payload.GetPassword())
		if err != nil {
			return err
		}
	}

	userModelCreatePayload, err := libs_util.ConvertToStruct[*model.User](
		&CreateUserRequestPbToUser{
			Name:     payload.GetName(),
			Email:    payload.GetEmail(),
			Password: hashedPassword,
			Role:     constant.UserRole(payload.GetRole()),
		},
	)
	if err != nil {
		return err
	}

	err = s.UserResource.CreateUser(userModelCreatePayload)
	if err != nil {
		return err
	}

	return nil
}
