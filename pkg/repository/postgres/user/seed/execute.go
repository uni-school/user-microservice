package seed_user_postgres_repository

import (
	gorm_seeder "github.com/kachit/gorm-seeder"
	"github.com/sirupsen/logrus"
	"github.com/uni-school/user-microservice/shared/config"
)

func ExecuteUserSeeder(isWithClear bool) {
	usersSeeder := InitUsersSeeder(gorm_seeder.SeederConfiguration{})
	seedersStack := gorm_seeder.NewSeedersStack(config.DB)
	seedersStack.AddSeeder(&usersSeeder)

	if isWithClear {
		err := seedersStack.Clear()
		if err != nil {
			logrus.Println(err)
		}
		logrus.Println("user seeder succesfully cleared")
	}

	err := seedersStack.Seed()
	if err != nil {
		logrus.Println(err)
	}
	logrus.Println("user seeder succesfully executed")
}
