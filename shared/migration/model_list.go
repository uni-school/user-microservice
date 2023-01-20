package migration

import (
	"github.com/uni-school/user-microservice/pkg/model"
)

var (
	AutoMigrateModelList = []interface{}{
		&model.User{},
	}
)