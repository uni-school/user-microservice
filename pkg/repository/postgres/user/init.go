package user_postgres_repository

import (
	"github.com/uni-school/user-microservice/libs/types"
	"github.com/uni-school/user-microservice/pkg/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(payload *model.User) error
	GetUser(query *types.Query, payload *model.User) (*model.User, error)
	GetListUser(query *types.Query, payload *model.User) ([]*model.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func InitUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}
