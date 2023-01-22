package seed_user_postgres_repository

import (
	"github.com/uni-school/user-microservice/pkg/model"
	"github.com/uni-school/user-microservice/shared/constant"
)

var UserSeedData = []model.User{
	{
		ID:       "8ea778bc-3958-4e9f-8fa2-a8a9ad8f2ab1",
		Name:     "uni-admin",
		Email:    "uni-admin@gmail.com",
		Password: "p4ssw0rd",
		Role:     constant.ADMIN,
	},
}
