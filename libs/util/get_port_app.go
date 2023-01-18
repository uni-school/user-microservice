package util

import (
	"strconv"

	"github.com/uni-school/user-microservice/libs/config"
)

func GetPortApp() string {
	return strconv.Itoa(config.AppConfig.Server.Port)
}
