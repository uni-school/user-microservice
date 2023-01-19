package runner

import (
	seed_user_postgres_repository "github.com/uni-school/user-microservice/pkg/repository/postgres/user/seed"
)

func RunSeeder() {
	seed_user_postgres_repository.ExcecuteUserSeeder()
}