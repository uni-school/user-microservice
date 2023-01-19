package seed_user_postgres_repository

import (
	"github.com/uni-school/user-microservice/pkg/model"
	"github.com/uni-school/user-microservice/shared/constant"
)

var UserSeedData = []model.User{
	{
		Name:     "uni-admin",
		Email:    "uni-admin@gmail.com",
		Password: "p4ssw0rd",
		Role:     constant.ADMIN,
	},
}
