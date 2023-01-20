package seed_user_postgres_repository

import (
	gorm_seeder "github.com/kachit/gorm-seeder"
	"github.com/sirupsen/logrus"
	"github.com/uni-school/user-microservice/libs/config"
	"github.com/uni-school/user-microservice/pkg/model"
	user_postgres_repository "github.com/uni-school/user-microservice/pkg/repository/postgres/user"
	"github.com/uni-school/user-microservice/shared/constant"
	"github.com/uni-school/user-microservice/shared/util"
	"gorm.io/gorm"
)

type UsersSeeder struct {
	gorm_seeder.SeederAbstract
	UserRepository user_postgres_repository.IUserRepository
}

func InitUsersSeeder(cfg gorm_seeder.SeederConfiguration) UsersSeeder {
	return UsersSeeder{gorm_seeder.NewSeederAbstract(cfg), user_postgres_repository.InitUserRepository(config.DB)}
}

func (s *UsersSeeder) Seed(db *gorm.DB) error {
	var users []model.User
	for _, userSeed := range UserSeedData {
		if validEmail := IsValidEmail(userSeed.Email); !validEmail {
			logrus.Println("this email", userSeed.Email, "not valid")
			continue
		}

		userModelGetResult, err := s.UserRepository.GetUser(nil, &model.User{
			Email: userSeed.Email,
		})
		if err != nil {
			if err.Error() != gorm.ErrRecordNotFound.Error() {
				logrus.Println("error get user by email", err)
				continue
			}
		}
		if userModelGetResult != nil {
			logrus.Println("user with email", userSeed.Email, "already registered")
			continue
		}

		hashedPassword, err := util.HashPassword(userSeed.Password)
		if err != nil {
			logrus.Println("error hashing password", err)
			continue
		}

		user := model.User{
			Name:     userSeed.Name,
			Email:    userSeed.Email,
			Password: hashedPassword,
			Role:     constant.ADMIN,
		}
		users = append(users, user)
	}
	return db.CreateInBatches(users, len(users)).Error
}

func (s *UsersSeeder) Clear(db *gorm.DB) error {
	return s.SeederAbstract.Delete(db, string(constant.USER_TABLE_NAME))
}
