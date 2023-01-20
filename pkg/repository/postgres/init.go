package postgres_repository

import (
	user_postgres_repository "github.com/uni-school/user-microservice/pkg/repository/postgres/user"
	"github.com/uni-school/user-microservice/shared/config"
)

type PostgresRepository struct {
	UserRepository user_postgres_repository.IUserRepository
}

func InitPostgresRepository() *PostgresRepository {
	var (
		userRepository = user_postgres_repository.InitUserRepository(config.DB)
	)

	return &PostgresRepository{
		UserRepository: userRepository,
	}
}
