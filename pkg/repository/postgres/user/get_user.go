package user_postgres_repository

import (
	"github.com/uni-school/user-microservice/libs/constant"
	"github.com/uni-school/user-microservice/libs/types"
	"github.com/uni-school/user-microservice/libs/util"
	"github.com/uni-school/user-microservice/pkg/model"
)

func (r *UserRepository) GetUser(query *types.Query, payload *model.User) (*model.User, error) {
	model := new(model.User)

	sql := r.db.Model(model)

	if query != nil {
		if len(query.SelectColumns) > 0 {
			sql = sql.Select(util.MergeSlices(query.SelectColumns, constant.DEFAULT_SELECT_COLUMNS))
		}
	}

	if payload != nil {
		if payload.ID != "" {
			sql = sql.Where("id = ?", payload.ID)
		}
		if payload.Email != "" {
			sql = sql.Where("email = ?", payload.Email)
		}
	}

	if err := sql.First(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}
