package user_postgres_repository

import (
	"github.com/uni-school/user-microservice/pkg/model"
)

func (r *UserRepository) CreateUser(payload *model.User) error {
	model := new(model.User)

	sql := r.db.Model(model)

	if payload != nil {
		model = payload
	}

	if err := sql.Create(model).Error; err != nil {
		return err
	}

	return nil
}
