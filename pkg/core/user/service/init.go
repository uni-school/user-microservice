package service_user

import (
	resource_user "github.com/uni-school/user-microservice/pkg/core/user/resource"
	pb "github.com/uni-school/user-microservice/proto"
)

type IUserService interface {
	CreateUser(payload *pb.CreateUserRequest) error
	GetUserByEmail(payload *pb.GetUserByEmailRequest) (*pb.GetUserByEmailResponseData, error)
}

type UserService struct {
	UserResource resource_user.IUserResource
}

func InitUserService(userResource resource_user.IUserResource) IUserService {
	return &UserService{
		UserResource: userResource,
	}
}
