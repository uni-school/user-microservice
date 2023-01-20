package resource_user

import (
	"github.com/uni-school/user-microservice/libs/types"
	"github.com/uni-school/user-microservice/shared/config"
	"github.com/uni-school/user-microservice/pkg/model"
	postgres_repository "github.com/uni-school/user-microservice/pkg/repository/postgres"
)

type IUserResource interface {
	CreateUser(payload *model.User) error
	GetUser(query *types.Query, payload *model.User) (*model.User, error)
}

type UserResource struct {
	Postgres *postgres_repository.PostgresRepository
}

func InitUserResource() IUserResource {
	var (
		postgres = postgres_repository.InitPostgresRepository(config.DB)
	)

	return &UserResource{
		Postgres: postgres,
	}
}
