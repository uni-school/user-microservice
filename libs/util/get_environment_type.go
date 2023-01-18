package util

import (
	"github.com/uni-school/user-microservice/libs/config"
	"github.com/uni-school/user-microservice/libs/constant"
)

func GetEnvironmentType() constant.EnvironmentType {
	return config.Environment
}
