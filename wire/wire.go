//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"

	resource_user "github.com/uni-school/user-microservice/pkg/core/user/resource"
	service_user "github.com/uni-school/user-microservice/pkg/core/user/service"
)

func UserService() service_user.IUserService {
	wire.Build(
		service_user.InitUserService,
		resource_user.InitUserResource,
	)
	return &service_user.UserService{}
}
