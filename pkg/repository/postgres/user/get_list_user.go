package user_postgres_repository

import (
	"github.com/uni-school/user-microservice/libs/constant"
	"github.com/uni-school/user-microservice/libs/types"
	"github.com/uni-school/user-microservice/libs/util"
	"github.com/uni-school/user-microservice/pkg/model"
)

func (r *UserRepository) GetListUser(query *types.Query, payload *model.User) ([]model.User, error) {
	user := new(model.User)
	users := make([]model.User, 0)

	sql := r.db.Model(user)

	if query != nil {
		if len(query.SelectColumns) > 0 {
			sql = sql.Select(util.MergeSlices(query.SelectColumns, constant.DEFAULT_SELECT_COLUMNS))
		}
	}

	if payload != nil {
		if payload.Role != "" {
			sql = sql.Where("role = ?", payload.Role)
		}
	}

	if err := sql.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}