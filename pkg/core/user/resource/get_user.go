package resource_user

import (
	"github.com/uni-school/user-microservice/libs/types"
	"github.com/uni-school/user-microservice/pkg/model"
)

func (r *UserResource) GetUser(query *types.Query, payload *model.User) (*model.User, error) {
	return r.Postgres.UserRepository.GetUser(query, payload)
}
