package util

import (
	"strconv"

	"github.com/uni-school/user-microservice/shared/config"
	"github.com/uni-school/user-microservice/shared/constant"
)

func GetPortApp() string {
	return strconv.Itoa(config.AppConfig.Server.Port)
}

func GetEnvironmentType() constant.EnvironmentType {
	return config.Environment
}