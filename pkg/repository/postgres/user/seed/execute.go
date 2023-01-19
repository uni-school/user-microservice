package seed_user_postgres_repository

import (
	gorm_seeder "github.com/kachit/gorm-seeder"
	"github.com/sirupsen/logrus"
	"github.com/uni-school/user-microservice/libs/config"
)

func ExcecuteUserSeeder() {
	usersSeeder := InitUsersSeeder(gorm_seeder.SeederConfiguration{})
	seedersStack := gorm_seeder.NewSeedersStack(config.DB)
	seedersStack.AddSeeder(&usersSeeder)

	err := seedersStack.Seed()
	if err != nil {
		logrus.Println(err)
	}
	logrus.Println("user seeder succesfully executed")
}
