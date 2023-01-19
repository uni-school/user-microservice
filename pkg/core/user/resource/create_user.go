package resource_user

import (
	"github.com/uni-school/user-microservice/pkg/model"
)

func (r *UserResource) CreateUser(payload *model.User) error {
	return r.Postgres.UserRepository.CreateUser(payload)
}
