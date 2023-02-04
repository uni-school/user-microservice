package runner

import (
	seed_user_postgres_repository "github.com/uni-school/user-microservice/pkg/repository/postgres/user/seed"
	"github.com/uni-school/user-microservice/shared/config"
)

func RunSeeder() {
	seed_user_postgres_repository.ExecuteUserSeeder(config.AppConfig.DatabaseSeeder.User.IsRebuildData)
}
