package user_postgres_repository

import (
	"github.com/uni-school/user-microservice/pkg/model"
)

func (r *UserRepository) CreateUser(payload *model.User) error {
	user := new(model.User)

	sql := r.db.Model(user)

	if payload != nil {
		user = payload
	}

	if err := sql.Create(user).Error; err != nil {
		return err
	}

	return nil
}
